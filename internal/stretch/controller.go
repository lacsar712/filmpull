package stretch

import (
	"context"

	"github.com/lacsar712/filmpull/internal/clock"
	"github.com/lacsar712/filmpull/internal/model"
)

type Controller struct {
	table *ZoneTable
	clk   clock.Clock
	maxDraw float64
}

func NewController(table *ZoneTable, clk clock.Clock, maxDraw float64) *Controller {
	return &Controller{table: table, clk: clk, maxDraw: maxDraw}
}

func (c *Controller) RampZone(ctx context.Context, id model.ZoneID, target float64, ramp float64) error {
	if target < 0 {
		return model.Wrap("stretch", "ramp", model.ErrInvalid)
	}
	zones := c.table.Zones()
	var cur float64
	for _, z := range zones {
		if z.ID == id {
			cur = z.SpeedMPM
			break
		}
	}
	for cur < target {
		select {
		case <-ctx.Done():
			return context.Cause(ctx)
		default:
		}
		step := ramp * 0.1
		if cur+step > target {
			cur = target
		} else {
			cur += step
		}
		if err := c.table.UpdateSpeed(id, cur); err != nil {
			return err
		}
	}
	return c.table.UpdateSpeed(id, target)
}

func (c *Controller) ApplyDraw(id model.ZoneID, ratio float64) error {
	if ratio > c.maxDraw {
		return model.Wrap("stretch", "draw_limit", model.ErrInvalid)
	}
	return c.table.SetDraw(id, ratio)
}

func (c *Controller) Snapshot() []model.StretchZone {
	return c.table.Zones()
}