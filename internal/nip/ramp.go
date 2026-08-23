package nip

import (
	"time"

	"github.com/lacsar712/filmpull/internal/model"
)

type RampPhase string

const (
	RampSeat    RampPhase = "seat"
	RampWorking RampPhase = "working"
	RampRelease RampPhase = "release"
)

// PressureRampProfile steps nip pressure through seat and working setpoints.
type PressureRampProfile struct {
	nipID       model.NipID
	phase       RampPhase
	currentKPa  float64
	targetKPa   float64
	seatKPa     float64
	workingKPa  float64
	rateKPaPerS float64
	started     time.Time
}

func NewPressureRampProfile(id model.NipID, working, seatFactor, rate float64) PressureRampProfile {
	if seatFactor <= 0 {
		seatFactor = 0.85
	}
	if rate <= 0 {
		rate = 5
	}
	seat := SeatPressure(working, seatFactor)
	return PressureRampProfile{
		nipID: id, phase: RampSeat, seatKPa: seat, workingKPa: working,
		targetKPa: seat, rateKPaPerS: rate,
	}
}

func (p *PressureRampProfile) Begin(at time.Time) {
	p.started = at
	p.currentKPa = 0
	p.phase = RampSeat
	p.targetKPa = p.seatKPa
}

func (p *PressureRampProfile) Tick(dt time.Duration) float64 {
	step := p.rateKPaPerS * dt.Seconds()
	p.currentKPa = RampPressure(p.currentKPa, p.targetKPa, step)
	if p.phase == RampSeat && p.currentKPa >= p.seatKPa {
		p.phase = RampWorking
		p.targetKPa = p.workingKPa
	}
	return p.currentKPa
}

func (p *PressureRampProfile) Release(dt time.Duration) float64 {
	p.phase = RampRelease
	step := p.rateKPaPerS * dt.Seconds()
	p.currentKPa = ReleasePressure(p.currentKPa, step)
	return p.currentKPa
}

func (p *PressureRampProfile) Done() bool {
	if p.phase == RampRelease {
		return p.currentKPa <= 0
	}
	return p.phase == RampWorking && p.currentKPa >= p.workingKPa
}

func (p *PressureRampProfile) Phase() RampPhase { return p.phase }
func (p *PressureRampProfile) Current() float64 { return p.currentKPa }

// RampBank tracks active pressure ramps for all nip rolls on the line.
type RampBank struct {
	profiles map[model.NipID]*PressureRampProfile
}

func NewRampBank() *RampBank {
	return &RampBank{profiles: make(map[model.NipID]*PressureRampProfile)}
}

func (b *RampBank) Start(id model.NipID, working float64, at time.Time) {
	prof := NewPressureRampProfile(id, working, 0.85, 5)
	prof.Begin(at)
	b.profiles[id] = &prof
}

func (b *RampBank) Advance(id model.NipID, dt time.Duration) (float64, bool) {
	prof, ok := b.profiles[id]
	if !ok {
		return 0, false
	}
	if prof.phase == RampRelease {
		kpa := prof.Release(dt)
		if prof.Done() {
			delete(b.profiles, id)
		}
		return kpa, true
	}
	return prof.Tick(dt), true
}

func (b *RampBank) RequestRelease(id model.NipID) {
	prof, ok := b.profiles[id]
	if !ok {
		return
	}
	prof.phase = RampRelease
}

func (b *RampBank) Active(id model.NipID) bool {
	_, ok := b.profiles[id]
	return ok
}

func (b *RampBank) Snapshot() map[model.NipID]RampPhase {
	out := make(map[model.NipID]RampPhase, len(b.profiles))
	for id, p := range b.profiles {
		out[id] = p.phase
	}
	return out
}
