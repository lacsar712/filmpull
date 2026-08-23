package app

import (
	"context"

	"github.com/lacsar712/filmpull/internal/model"
)

func (a *App) ExecutePlan(ctx context.Context, entries []model.StretchScheduleEntry) error {
	return a.lineSched.InstallPreheatPlanCtx(ctx, entries, a.cfg.PreheatPlanSteps, "preheat-plan")
}

func (a *App) PreheatStepsDone() int {
	return a.lineSched.PreheatStepsDone()
}
