package tension

import (
	"context"

	"github.com/lacsar712/filmpull/internal/clock"
	"github.com/lacsar712/filmpull/internal/model"
)

// ZoneState tracks regulation output for one stretch zone.
type ZoneState struct {
	Zone       model.ZoneID
	Setpoint   model.TensionSetpoint
	Output     float64
	LastActual float64
	InBand     bool
}

// ZoneRegulator applies cascade tension trim across enabled stretch zones.
type ZoneRegulator struct {
	states  map[model.ZoneID]*ZoneState
	profile Profile
	reg     *Regulator
	clk     clock.Clock
}

func NewZoneRegulator(profile Profile, reg *Regulator, clk clock.Clock) *ZoneRegulator {
	return &ZoneRegulator{
		states: make(map[model.ZoneID]*ZoneState),
		profile: profile, reg: reg, clk: clk,
	}
}

func (z *ZoneRegulator) BindZone(zone model.ZoneID, index int) {
	sp := z.profile.ZoneSetpoint(index)
	z.states[zone] = &ZoneState{Zone: zone, Setpoint: sp}
}

func (z *ZoneRegulator) BindAll(zones []model.StretchZone) {
	for i, zn := range zones {
		if zn.Enabled {
			z.BindZone(zn.ID, i)
		}
	}
}

func (z *ZoneRegulator) Regulate(ctx context.Context, zone model.ZoneID, sensors []model.SensorID) error {
	st, ok := z.states[zone]
	if !ok {
		return model.Wrap("tension", "zone_unbound", model.ErrNotFound)
	}
	actual, err := z.reg.Evaluate(ctx, zone, sensors)
	if err != nil {
		return err
	}
	st.LastActual = actual
	st.InBand = z.reg.InBand(st.Setpoint, actual)
	st.Output = z.reg.TrimOutput(st.Setpoint, actual, st.Output)
	return nil
}

func (z *ZoneRegulator) CascadeTrim(zones []model.StretchZone) {
	var upstream float64
	for i, zn := range zones {
		if !zn.Enabled {
			continue
		}
		st, ok := z.states[zn.ID]
		if !ok {
			z.BindZone(zn.ID, i)
			st = z.states[zn.ID]
		}
		if upstream > 0 && st.LastActual > 0 {
			overshoot := st.LastActual - st.Setpoint.Newtons
			if overshoot > 0 {
				st.Output -= overshoot * 0.05
			}
			if upstream > st.Setpoint.Newtons*1.1 {
				st.Output += (upstream - st.Setpoint.Newtons) * 0.02
			}
		}
		upstream = st.LastActual
	}
}

func (z *ZoneRegulator) ApplyPhase(phase string) {
	var sp model.TensionSetpoint
	switch phase {
	case "thread":
		sp = z.profile.ThreadingHold()
	case "preheat":
		sp = z.profile.PreheatHold()
	default:
		sp = z.profile.ZoneSetpoint(0)
	}
	for _, st := range z.states {
		st.Setpoint = sp
		st.Output = 0
	}
}

func (z *ZoneRegulator) State(zone model.ZoneID) (ZoneState, bool) {
	st, ok := z.states[zone]
	if !ok {
		return ZoneState{}, false
	}
	return *st, true
}

func (z *ZoneRegulator) AllInBand() bool {
	if len(z.states) == 0 {
		return false
	}
	for _, st := range z.states {
		if !st.InBand {
			return false
		}
	}
	return true
}

func (z *ZoneRegulator) MeanOutput() float64 {
	if len(z.states) == 0 {
		return 0
	}
	var sum float64
	for _, st := range z.states {
		sum += st.Output
	}
	return sum / float64(len(z.states))
}
