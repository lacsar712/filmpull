package fsm

import (
	"context"

	"github.com/lacsar712/filmpull/internal/model"
)

type NipSideEffect func(ctx context.Context, nip model.NipID, from, to model.NipState) error

type NipFSM struct {
	id       model.NipID
	state    model.NipState
	onChange NipSideEffect
}

func NewNipFSM(id model.NipID, effect NipSideEffect) *NipFSM {
	return &NipFSM{id: id, state: model.NipOpen, onChange: effect}
}

func (f *NipFSM) State() model.NipState { return f.state }

func (f *NipFSM) Apply(ctx context.Context, event string) error {
	next, err := MustNip(f.state, event)
	if err != nil {
		return err
	}
	prev := f.state
	if f.onChange != nil {
		if err := f.onChange(ctx, f.id, prev, next); err != nil {
			return model.Wrap("nip_fsm", "side_effect", err)
		}
	}
	f.state = next
	return nil
}