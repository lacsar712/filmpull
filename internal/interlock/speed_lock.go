package interlock

import (
	"github.com/lacsar712/filmpull/internal/model"
)

type SpeedLock struct {
	zeroRequired bool
	maxMPM       float64
}

func NewSpeedLock(maxMPM float64) *SpeedLock {
	return &SpeedLock{zeroRequired: true, maxMPM: maxMPM}
}

func (l *SpeedLock) PermitSpeedChange(current, target float64, nips []model.NipRoll) error {
	if l.zeroRequired && target > 0 {
		for _, n := range nips {
			if n.State != model.NipClosed && n.SpeedMPM > 0 {
				return model.Wrap("interlock", "nip_not_closed", model.ErrInterlock)
			}
		}
	}
	if target > l.maxMPM {
		return model.Wrap("interlock", "speed_cap", model.ErrSpeed)
	}
	if current < 0 || target < 0 {
		return model.Wrap("interlock", "negative_speed", model.ErrInvalid)
	}
	return nil
}

func (l *SpeedLock) PermitTensionRamp(tension float64, sp model.TensionSetpoint) error {
	if !sp.Within(tension) && tension > sp.Newtons*1.5 {
		return model.Wrap("interlock", "tension_surge", model.ErrTension)
	}
	return nil
}

func (l *SpeedLock) SetZeroRequired(on bool) { l.zeroRequired = on }