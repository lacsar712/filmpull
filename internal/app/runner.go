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
	state := a.lineFSM.State()
	if state == model.LineStretch {
		if err := a.RegulateTensionZones(ctx); err != nil {
			return err
		}
	}
	if state == model.LineThreading || state == model.LineStretch {
		a.AdvanceNipRamps()
	}
	if state == model.LineStretch {
		lag := a.speedCoord.ZoneLag(a.zones.Zones(), model.ZoneID("zone-1"))
		if lag > 5 {
			a.stats.IncEvent("speed_lag")
		}
	}
	snap, ok := a.mem.Snapshot(a.lineID)
	if ok {
		a.registry.Set("snapshot_version", float64(a.mem.Version()))
		_ = snap
	}
	return nil
}