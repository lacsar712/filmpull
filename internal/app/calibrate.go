package app

import (
	"context"
	"fmt"

	"github.com/lacsar712/filmpull/internal/model"
)

// CalibrateProbe allows acceptance tests to inject tension calibration faults.
var CalibrateProbe func(ctx context.Context) error

func (a *App) CalibrateGrip(ctx context.Context, nip model.NipID, holder string) error {
	lease, err := a.gripLeases.Acquire(string(nip), holder, a.clk.Now())
	if err != nil {
		return err
	}
	defer lease.Release()
	if CalibrateProbe != nil {
		if err := CalibrateProbe(ctx); err != nil {
			return fmt.Errorf("calibrate: %w", err)
		}
	}
	return nil
}

func (a *App) GripHeld(nip model.NipID) bool {
	return a.gripLeases.IsHeld(string(nip))
}
