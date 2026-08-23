package app

import (
	"context"

	"github.com/lacsar712/filmpull/internal/model"
	"github.com/lacsar712/filmpull/internal/nip"
)

func (a *App) ApplyGradePlan(gradeID model.FilmGradeID) error {
	grade, ok := a.grades.Lookup(gradeID)
	if !ok {
		return model.Wrap("app", "grade", model.ErrNotFound)
	}
	a.grade = grade
	plan, err := a.drawPlan.Plan(grade)
	if err != nil {
		return err
	}
	if err := a.drawPlan.Apply(a.zones, plan); err != nil {
		return err
	}
	a.speedCoord.SetInlet(a.speedMPM)
	a.zoneReg.BindAll(a.zones.Zones())
	a.registry.Set("planned_draw", plan.TotalDraw)
	return nil
}

func (a *App) PlanDrawRamp(ctx context.Context, enabledZones int) error {
	plan, err := a.drawPlan.PlanIncremental(a.grade, enabledZones)
	if err != nil {
		return err
	}
	if err := a.drawPlan.Apply(a.zones, plan); err != nil {
		return err
	}
	setpoints := a.speedCoord.Balance(a.zones.Zones(), a.speedMPM)
	for i, sp := range setpoints {
		zones := a.zones.Zones()
		if i >= len(zones) {
			break
		}
		target := sp.MPM
		cur := zones[i].SpeedMPM
		for cur < target {
			select {
			case <-ctx.Done():
				return context.Cause(ctx)
			default:
			}
			cur = a.speedCoord.RampStep(cur, target)
			if err := a.zones.UpdateSpeed(zones[i].ID, cur); err != nil {
				return err
			}
			a.clk.Step()
		}
	}
	a.registry.Set("draw_ramp_zones", float64(enabledZones))
	a.persistSnapshot()
	return nil
}

func (a *App) StartNipRamp(id model.NipID, workingKPa float64) {
	a.nipRamps.Start(id, workingKPa, a.clk.Now())
	a.stats.IncEvent("nip_ramp_start")
}

func (a *App) AdvanceNipRamps() {
	dt := a.cfg.ProcessTick()
	for id := range a.nipRamps.Snapshot() {
		kpa, ok := a.nipRamps.Advance(id, dt)
		if !ok {
			continue
		}
		if r, ok := a.nips.Roller(id); ok {
			_ = r.ApplyPressure(kpa)
		}
	}
}

func (a *App) ReleaseNipRamp(id model.NipID) {
	a.nipRamps.RequestRelease(id)
}

func (a *App) RegulateTensionZones(ctx context.Context) error {
	zones := a.zones.Zones()
	for _, z := range zones {
		if !z.Enabled {
			continue
		}
		sid := model.FormatSensor(z.ID, 1)
		if err := a.zoneReg.Regulate(ctx, z.ID, []model.SensorID{sid}); err != nil {
			if model.Is(err, model.ErrTension) {
				a.stats.IncEvent("tension_trim_miss")
				continue
			}
			return err
		}
	}
	a.zoneReg.CascadeTrim(zones)
	if a.zoneReg.AllInBand() {
		a.registry.Set("tension_in_band", 1)
	} else {
		a.registry.Set("tension_in_band", 0)
	}
	return nil
}

func (a *App) CoordinateLineSpeed(targetMPM float64) error {
	nipSnaps := a.nips.Snapshots()
	if err := a.speed.PermitSpeedChange(a.speedMPM, targetMPM, nipSnaps); err != nil {
		return err
	}
	setpoints := a.speedCoord.Balance(a.zones.Zones(), targetMPM)
	if len(setpoints) > 0 {
		a.speedMPM = setpoints[len(setpoints)-1].MPM
	} else {
		a.speedMPM = targetMPM
	}
	a.speedCoord.SetInlet(a.speedMPM)
	a.registry.Set("line_speed_mpm", a.speedMPM)
	a.registry.Set("outlet_speed_mpm", a.speedCoord.OutletMPM(a.zones.Zones()))
	return nil
}

func (a *App) NipRampPhases() map[model.NipID]nip.RampPhase {
	return a.nipRamps.Snapshot()
}

func (a *App) GradeTable() *model.GradeTable { return a.grades }
