package store

import (
	"testing"
	"time"

	"github.com/lacsar712/filmpull/internal/clock"
	"github.com/lacsar712/filmpull/internal/model"
)

func TestSnapshotBuilder(t *testing.T) {
	id, _ := model.ParseLineID("line-a")
	b := NewSnapshotBuilder(id).State(model.LineStretch).Speed(200)
	snap := b.Build(time.Now())
	if snap.State != model.LineStretch || snap.SpeedMPM != 200 {
		t.Fatalf("snap %+v", snap)
	}
}

func TestMemorySnapshot(t *testing.T) {
	mem := NewMemory()
	id, _ := model.ParseLineID("line-a")
	snap := NewSnapshotBuilder(id).Build(time.Now())
	mem.PutSnapshot(snap)
	got, ok := mem.Snapshot(id)
	if !ok || got.ID != id {
		t.Fatal("missing snapshot")
	}
}

func TestScheduleStoreActive(t *testing.T) {
	mem := NewMemory()
	start := time.Date(2024, 6, 1, 0, 0, 0, 0, time.UTC)
	clk := clock.NewProcessClock(start, time.Second)
	store := NewScheduleStore(mem, clk)
	sch := model.StretchSchedule{ID: "sch-1", Entries: []model.StretchScheduleEntry{{
		ID: "e1", Zone: "zone-a", Start: start, End: start.Add(10 * time.Second), Draw: 3.5,
	}}}
	store.Upsert(sch)
	clk.Advance(5 * time.Second)
	entry, ok := store.ActiveEntry(sch)
	if !ok || entry.ID != "e1" {
		t.Fatalf("entry %+v ok=%v", entry, ok)
	}
}