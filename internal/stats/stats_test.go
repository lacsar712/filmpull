package stats

import (
	"testing"
	"time"
)

func TestCollectorMean(t *testing.T) {
	c := NewCollector()
	now := time.Now()
	c.RecordTension("zone-a", 100, now)
	c.RecordTension("zone-a", 120, now)
	mean, ok := c.MeanTension("zone-a")
	if !ok || mean != 110 {
		t.Fatalf("mean %v", mean)
	}
}

func TestRegistrySnapshot(t *testing.T) {
	r := NewRegistry()
	r.Set("line_speed_mpm", 220)
	snap := r.Snapshot()
	if snap["line_speed_mpm"] != 220 {
		t.Fatal(snap)
	}
}

func TestCollectorEvents(t *testing.T) {
	c := NewCollector()
	c.IncEvent("thread_start")
	c.IncEvent("thread_start")
	if c.Count("thread_start") != 2 {
		t.Fatal()
	}
}