package stretch

import (
	"github.com/lacsar712/filmpull/internal/model"
)

// SpeedCoordinator keeps inlet and zone outlet speeds aligned during ramps.
type SpeedCoordinator struct {
	inletMPM   float64
	maxRamp    float64
	planner    *LineSpeedPlanner
	lastOutlet float64
}

func NewSpeedCoordinator(inletMPM, maxRamp float64) *SpeedCoordinator {
	if maxRamp <= 0 {
		maxRamp = 2.0
	}
	return &SpeedCoordinator{
		inletMPM: inletMPM,
		maxRamp:  maxRamp,
		planner:  NewLineSpeedPlanner(inletMPM),
	}
}

func (c *SpeedCoordinator) SetInlet(mpm float64) {
	c.inletMPM = mpm
	c.planner.SetInlet(mpm)
}

func (c *SpeedCoordinator) Plan(zones []model.StretchZone) []model.SpeedSetpoint {
	return c.planner.Plan(zones)
}

func (c *SpeedCoordinator) OutletMPM(zones []model.StretchZone) float64 {
	out := c.planner.OutletMPM(zones)
	c.lastOutlet = out
	return out
}

func (c *SpeedCoordinator) RampStep(current, target float64) float64 {
	delta := target - current
	if delta > c.maxRamp {
		return current + c.maxRamp
	}
	if delta < -c.maxRamp {
		return current - c.maxRamp
	}
	return target
}

func (c *SpeedCoordinator) Balance(zones []model.StretchZone, lineMPM float64) []model.SpeedSetpoint {
	planned := c.Plan(zones)
	if len(planned) == 0 {
		return planned
	}
	draw := EffectiveDraw(zones)
	if draw <= 0 {
		draw = 1
	}
	targetOutlet := lineMPM
	if targetOutlet <= 0 {
		targetOutlet = c.inletMPM * draw
	}
	scale := 1.0
	if planned[len(planned)-1].MPM > 0 {
		scale = targetOutlet / planned[len(planned)-1].MPM
	}
	for i := range planned {
		planned[i].MPM *= scale
		if planned[i].RampMPMPerS > c.maxRamp {
			planned[i].RampMPMPerS = c.maxRamp
		}
	}
	return planned
}

func (c *SpeedCoordinator) ZoneLag(zones []model.StretchZone, zone model.ZoneID) float64 {
	speeds := ZoneSpeeds(c.inletMPM, zones)
	var idx int
	for i, z := range zones {
		if z.ID == zone {
			idx = i
			break
		}
	}
	if idx >= len(speeds) {
		return 0
	}
	return speeds[idx] - zones[idx].SpeedMPM
}

func (c *SpeedCoordinator) ReadyForStretch(zones []model.StretchZone, minMPM float64) bool {
	for _, z := range zones {
		if !z.Enabled {
			continue
		}
		if z.SpeedMPM < minMPM {
			return false
		}
	}
	return c.inletMPM >= minMPM
}

func (c *SpeedCoordinator) LastOutlet() float64 { return c.lastOutlet }
