package clock

import (
	"fmt"
	"time"

	"github.com/lacsar712/filmpull/internal/model"
)

type SetWindow struct {
	clk      Clock
	duration time.Duration
}

func NewSetWindow(clk Clock, duration time.Duration) *SetWindow {
	if duration <= 0 {
		duration = 5 * time.Minute
	}
	return &SetWindow{clk: clk, duration: duration}
}

func (w *SetWindow) Ready(anchor time.Time) bool {
	return time.Since(anchor) < w.duration
}

func (w *SetWindow) Require(anchor time.Time) error {
	if w.Ready(anchor) {
		return nil
	}
	return fmt.Errorf("set window: %w", model.ErrSetHold)
}
