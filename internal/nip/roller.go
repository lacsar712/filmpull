package nip

import (
	"github.com/lacsar712/filmpull/internal/model"
)

type RollerPair struct {
	id      model.NipID
	state   model.NipState
	pressure float64
	maxKPa   float64
	speed    float64
}

func NewRollerPair(id model.NipID, maxKPa float64) *RollerPair {
	return &RollerPair{id: id, state: model.NipOpen, maxKPa: maxKPa}
}

func (r *RollerPair) ID() model.NipID { return r.id }
func (r *RollerPair) State() model.NipState { return r.state }
func (r *RollerPair) Pressure() float64 { return r.pressure }
func (r *RollerPair) Speed() float64 { return r.speed }

func (r *RollerPair) SetState(s model.NipState) { r.state = s }

func (r *RollerPair) ApplyPressure(kpa float64) error {
	if kpa < 0 || kpa > r.maxKPa {
		return model.Wrap("nip", "pressure", model.ErrInvalid)
	}
	r.pressure = kpa
	return nil
}

func (r *RollerPair) SetSpeed(mpm float64) { r.speed = mpm }

func (r *RollerPair) Snapshot() model.NipRoll {
	return model.NipRoll{ID: r.id, State: r.state, PressureKPa: r.pressure, MaxPressure: r.maxKPa, SpeedMPM: r.speed}
}