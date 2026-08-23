package fsm

import (
	"context"

	"github.com/lacsar712/filmpull/internal/model"
)

type LineTransitionHook func(ctx context.Context, line model.LineID, from, to model.LineState, event string) error

type LineHookChain struct {
	after []LineTransitionHook
}

func NewLineHookChain() *LineHookChain { return &LineHookChain{} }

func (h *LineHookChain) OnAfter(fn LineTransitionHook) { h.after = append(h.after, fn) }

func (h *LineHookChain) RunAfter(ctx context.Context, line model.LineID, from, to model.LineState, event string) error {
	for _, fn := range h.after {
		if err := fn(ctx, line, from, to, event); err != nil {
			return err
		}
	}
	return nil
}
