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

// Note: window expiry is measured against clk (the process clock that follows
// the stretch rhythm), not the wall clock. Using time.Since(anchor) here would
// tie the hold window to the shop wall clock, so it would keep elapsing during
// a stop-line drill (where the process beat is frozen) and the set-hold
// judgement would report "complete" early instead of tracking the beat.

func NewSetWindow(clk Clock, duration time.Duration) *SetWindow {
	if duration <= 0 {
		duration = 5 * time.Minute
	}
	return &SetWindow{clk: clk, duration: duration}
}

func (w *SetWindow) Ready(anchor time.Time) bool {
	return w.clk.Now().Sub(anchor) < w.duration
}

func (w *SetWindow) Require(anchor time.Time) error {
	if w.Ready(anchor) {
		return nil
	}
	return fmt.Errorf("set window: %w", model.ErrSetHold)
}
