package stretch

import "github.com/lacsar712/filmpull/internal/model"

// LineSpeedPlanner coordinates inlet and outlet speeds across enabled stretch zones.
type LineSpeedPlanner struct {
	inletMPM float64
}

func NewLineSpeedPlanner(inlet float64) *LineSpeedPlanner {
	return &LineSpeedPlanner{inletMPM: inlet}
}

func (p *LineSpeedPlanner) Plan(zones []model.StretchZone) []model.SpeedSetpoint {
	out := make([]model.SpeedSetpoint, len(zones))
	speeds := ZoneSpeeds(p.inletMPM, zones)
	for i, z := range zones {
		ramp := 2.0
		if z.DrawRatio > 3 {
			ramp = 1.5
		}
		out[i] = model.SpeedSetpoint{MPM: speeds[i], RampMPMPerS: ramp}
	}
	return out
}

func (p *LineSpeedPlanner) OutletMPM(zones []model.StretchZone) float64 {
	speeds := ZoneSpeeds(p.inletMPM, zones)
	if len(speeds) == 0 {
		return p.inletMPM
	}
	return speeds[len(speeds)-1]
}

func (p *LineSpeedPlanner) SetInlet(mpm float64) { p.inletMPM = mpm }
