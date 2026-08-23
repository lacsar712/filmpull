package app

import (
	"context"
	"time"

	"github.com/lacsar712/filmpull/internal/model"
)

func (a *App) Run(ctx context.Context) error {
	ticker := time.NewTicker(a.cfg.ProcessTick())
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return context.Cause(ctx)
		case <-ticker.C:
			a.clk.Step()
			if err := a.tick(ctx); err != nil {
				return err
			}
		}
	}
}

func (a *App) tick(ctx context.Context) error {
	if a.lineFSM.State() == model.LineStretch {
		zones := a.zones.Zones()
		for _, z := range zones {
			if !z.Enabled {
				continue
			}
			sid := model.FormatSensor(z.ID, 1)
			avg, err := a.tension.Evaluate(ctx, z.ID, []model.SensorID{sid})
			if err != nil && model.Is(err, model.ErrTension) {
				a.stats.IncEvent("tension_hold_miss")
				continue
			}
			_ = avg
		}
	}
	snap, ok := a.mem.Snapshot(a.lineID)
	if ok {
		a.registry.Set("snapshot_version", float64(a.mem.Version()))
		_ = snap
	}
	return nil
}