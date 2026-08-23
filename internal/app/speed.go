package app

import (
	"context"
	"time"

	"github.com/lacsar712/filmpull/internal/clock"
)

type SpeedRamp struct {
	clk   *clock.ProcessClock
	steps int
	delay time.Duration
}

func NewSpeedRamp(clk *clock.ProcessClock, steps int, delay time.Duration) *SpeedRamp {
	if steps <= 0 {
		steps = 40
	}
	if delay <= 0 {
		delay = time.Millisecond
	}
	return &SpeedRamp{clk: clk, steps: steps, delay: delay}
}

func (r *SpeedRamp) Ramp(ctx context.Context, start, target float64, apply func(float64)) error {
	if r.steps <= 0 {
		apply(target)
		return nil
	}
	step := (target - start) / float64(r.steps)
	cur := start
	for i := 0; i < r.steps; i++ {
		select {
		case <-ctx.Done():
			return context.Cause(ctx)
		default:
		}
		cur += step
		if i == r.steps-1 {
			cur = target
		}
		apply(cur)
		r.clk.Step()
		time.Sleep(r.delay)
	}
	return nil
}

func (a *App) SpeedRamp(ctx context.Context, target float64) error {
	return a.speedRamp.Ramp(ctx, a.speedMPM, target, func(v float64) {
		a.speedMPM = v
		a.registry.Set("line_speed_mpm", v)
	})
}
