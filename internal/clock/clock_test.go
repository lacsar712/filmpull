package clock

import (
	"context"
	"testing"
	"time"
)

func TestProcessClockStep(t *testing.T) {
	start := time.Date(2024, 3, 1, 0, 0, 0, 0, time.UTC)
	clk := NewProcessClock(start, 50*time.Millisecond)
	if clk.Now() != start {
		t.Fatal("start mismatch")
	}
	next := clk.Step()
	if !next.Equal(start.Add(50 * time.Millisecond)) {
		t.Fatalf("step: %v", next)
	}
}

func TestWindowElapsed(t *testing.T) {
	start := time.Date(2024, 3, 1, 0, 0, 0, 0, time.UTC)
	clk := NewProcessClock(start, time.Second)
	winStart := start.Add(2 * time.Second)
	clk.Advance(3 * time.Second)
	if !WindowElapsed(clk, winStart, 2*time.Second) {
		t.Fatal("expected window active")
	}
	clk.Advance(2 * time.Second)
	if WindowElapsed(clk, winStart, 2*time.Second) {
		t.Fatal("expected window closed")
	}
}

func TestWaitUntilContext(t *testing.T) {
	start := time.Date(2024, 3, 1, 0, 0, 0, 0, time.UTC)
	clk := NewProcessClock(start, 10*time.Millisecond)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(5 * time.Millisecond)
			clk.Step()
		}
	}()
	if err := WaitUntilContext(ctx, clk, start.Add(20*time.Millisecond)); err != nil {
		t.Fatal(err)
	}
}