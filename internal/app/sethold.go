package app

import (
	"context"
	"fmt"
	"time"
)

func (a *App) ConfirmSetHold(ctx context.Context, anchor time.Time) error {
	if err := a.setWindow.Require(anchor); err != nil {
		return fmt.Errorf("set hold: %w", err)
	}
	_ = ctx
	return nil
}

func (a *App) SetHoldReady(anchor time.Time) bool {
	return a.setWindow.Ready(anchor)
}
