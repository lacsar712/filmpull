package fsm

import (
	"context"

	"github.com/lacsar712/filmpull/internal/model"
)

type LineSideEffect func(ctx context.Context, line model.LineID, from, to model.LineState) error

type LineFSM struct {
	id       model.LineID
	state    model.LineState
	onChange LineSideEffect
}

func NewLineFSM(id model.LineID, effect LineSideEffect) *LineFSM {
	return &LineFSM{id: id, state: model.LineIdle, onChange: effect}
}

func (f *LineFSM) State() model.LineState { return f.state }

func (f *LineFSM) Apply(ctx context.Context, event string) error {
	next, err := MustLine(f.state, event)
	if err != nil {
		return err
	}
	prev := f.state
	if f.onChange != nil {
		if err := f.onChange(ctx, f.id, prev, next); err != nil {
			return model.Wrap("line_fsm", "side_effect", err)
		}
	}
	f.state = next
	return nil
}

func (f *LineFSM) Force(state model.LineState) { f.state = state }