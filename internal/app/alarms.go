package app

import (
	"context"
	"fmt"

	"github.com/lacsar712/filmpull/internal/model"
)

func (a *App) ReportTensionFault(ctx context.Context, newtons float64) error {
	_ = ctx
	limit := a.grade.NominalTensionN * (1 + a.cfg.TensionTolerance/100)
	if newtons <= limit {
		return nil
	}
	return fmt.Errorf("tension fault: %w", model.ErrTensionExceeded)
}

func (a *App) HandleGripSlip(ctx context.Context, slipPct float64) error {
	if err := a.gripGuard.Permit(slipPct); err != nil {
		a.stats.IncEvent("grip_slip_alarm")
		return fmt.Errorf("line: %w", err)
	}
	_ = ctx
	return nil
}
