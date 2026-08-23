package nip

import (
	"context"

	"github.com/lacsar712/filmpull/internal/fsm"
	"github.com/lacsar712/filmpull/internal/model"
)

type Coordinator struct {
	rollers map[model.NipID]*RollerPair
	fsms    map[model.NipID]*fsm.NipFSM
}

func NewCoordinator(ids []model.NipID, maxKPa float64) *Coordinator {
	rollers := make(map[model.NipID]*RollerPair, len(ids))
	fsms := make(map[model.NipID]*fsm.NipFSM, len(ids))
	for _, id := range ids {
		rollers[id] = NewRollerPair(id, maxKPa)
		fsms[id] = fsm.NewNipFSM(id, nil)
	}
	return &Coordinator{rollers: rollers, fsms: fsms}
}

func (c *Coordinator) Close(ctx context.Context, id model.NipID, targetKPa float64) error {
	f, ok := c.fsms[id]
	if !ok {
		return model.Wrap("nip", "close", model.ErrNotFound)
	}
	r, ok := c.rollers[id]
	if !ok {
		return model.Wrap("nip", "close", model.ErrNotFound)
	}
	if err := f.Apply(ctx, "close"); err != nil {
		return err
	}
	r.SetState(f.State())
	cur := r.Pressure()
	for cur < targetKPa {
		select {
		case <-ctx.Done():
			return context.Cause(ctx)
		default:
		}
		cur = RampPressure(cur, targetKPa, 5)
		if err := r.ApplyPressure(cur); err != nil {
			return err
		}
	}
	if err := f.Apply(ctx, "seated"); err != nil {
		return err
	}
	r.SetState(f.State())
	return nil
}

func (c *Coordinator) Open(ctx context.Context, id model.NipID) error {
	f, ok := c.fsms[id]
	if !ok {
		return model.Wrap("nip", "open", model.ErrNotFound)
	}
	r := c.rollers[id]
	if err := f.Apply(ctx, "open"); err != nil {
		return err
	}
	r.SetState(f.State())
	return r.ApplyPressure(0)
}

func (c *Coordinator) Snapshots() []model.NipRoll {
	out := make([]model.NipRoll, 0, len(c.rollers))
	for _, r := range c.rollers {
		out = append(out, r.Snapshot())
	}
	return out
}

func (c *Coordinator) Roller(id model.NipID) (*RollerPair, bool) {
	r, ok := c.rollers[id]
	return r, ok
}