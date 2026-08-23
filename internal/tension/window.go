package tension

import (
	"time"

	"github.com/lacsar712/filmpull/internal/clock"
	"github.com/lacsar712/filmpull/internal/model"
)

type HoldWindow struct {
	start    time.Time
	duration time.Duration
	setpoint model.TensionSetpoint
}

func NewHoldWindow(clk clock.Clock, duration time.Duration, sp model.TensionSetpoint) HoldWindow {
	return HoldWindow{start: clk.Now(), duration: duration, setpoint: sp}
}

func (w HoldWindow) Active(clk clock.Clock) bool {
	return clock.WindowElapsed(clk, w.start, w.duration)
}

func (w HoldWindow) Satisfied(clk clock.Clock, actual float64) bool {
	if !w.Active(clk) {
		return false
	}
	return w.setpoint.Within(actual)
}

func (w HoldWindow) Remaining(clk clock.Clock) time.Duration {
	end := w.start.Add(w.duration)
	now := clk.Now()
	if !now.Before(end) {
		return 0
	}
	return end.Sub(now)
}