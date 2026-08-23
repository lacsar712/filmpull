package clock

import (
	"testing"
	"time"
)

func TestCase(t *testing.T) {
	start := time.Unix(0, 0)
	clk := NewProcessClock(start, time.Millisecond)
	w := NewSetWindow(clk, 2*time.Second)
	anchor := clk.Now()
	if !w.Ready(anchor) {
		t.Fatal("set window should start ready")
	}
	time.Sleep(3 * time.Second)
	if !w.Ready(anchor) {
		t.Fatal("window closed on wall clock while process clock frozen")
	}
	clk.Advance(3 * time.Second)
	if w.Ready(anchor) {
		t.Fatal("window should end after process clock advance")
	}
}
