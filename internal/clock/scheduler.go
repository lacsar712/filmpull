package clock

import (
	"context"
	"sync"
	"time"

	"github.com/lacsar712/filmpull/internal/model"
)

type LineScheduler struct {
	clk           ProcessClock
	mu            sync.Mutex
	preheatSteps  int
}

func NewLineScheduler(clk ProcessClock) *LineScheduler {
	return &LineScheduler{clk: clk}
}

func (s *LineScheduler) PreheatStepsDone() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.preheatSteps
}

func (s *LineScheduler) InstallPreheatPlan(entries []model.StretchScheduleEntry, steps int, planID string) error {
	return s.InstallPreheatPlanCtx(context.Background(), entries, steps, planID)
}

func (s *LineScheduler) InstallPreheatPlanCtx(ctx context.Context, entries []model.StretchScheduleEntry, steps int, planID string) error {
	if steps <= 0 {
		steps = len(entries)
	}
	if steps <= 0 {
		steps = 60
	}
	for i := 0; i < steps; i++ {
		if err := ctx.Err(); err != nil {
			return err
		}
		s.mu.Lock()
		s.preheatSteps = i + 1
		s.mu.Unlock()
		s.clk.Step()
		time.Sleep(2 * time.Millisecond)
	}
	_ = planID
	_ = entries
	return nil
}
