package app

import (
	"context"

	"github.com/lacsar712/filmpull/internal/model"
)

func (a *App) BeginStretchScope(ctx context.Context, line model.LineID) (context.Context, context.CancelFunc) {
	if line == "" {
		line = a.lineID
	}
	a.stretchMu.Lock()
	if cancel, ok := a.stretchCancels[line]; ok {
		cancel()
	}
	child, cancel := context.WithCancel(ctx)
	a.stretchCancels[line] = cancel
	a.stretchMu.Unlock()
	release := func() {
		a.stretchMu.Lock()
		delete(a.stretchCancels, line)
		a.stretchMu.Unlock()
		cancel()
	}
	return child, release
}

func (a *App) RunStretch(ctx context.Context, line model.LineID, fn func(context.Context) error) error {
	stretchCtx, release := a.BeginStretchScope(ctx, line)
	defer release()
	return fn(stretchCtx)
}
