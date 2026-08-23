package store

import (
	"time"

	"github.com/lacsar712/filmpull/internal/clock"
	"github.com/lacsar712/filmpull/internal/model"
)

type ScheduleStore struct {
	mem  *Memory
	clk  clock.Clock
}

func NewScheduleStore(mem *Memory, clk clock.Clock) *ScheduleStore {
	return &ScheduleStore{mem: mem, clk: clk}
}

func (s *ScheduleStore) ActiveEntry(sch model.StretchSchedule) (model.StretchScheduleEntry, bool) {
	now := s.clk.Now()
	for _, e := range sch.Entries {
		if !now.Before(e.Start) && now.Before(e.End) {
			return e, true
		}
	}
	return model.StretchScheduleEntry{}, false
}

func (s *ScheduleStore) Upsert(sch model.StretchSchedule) {
	sch.Version = s.mem.Version() + 1
	s.mem.PutSchedule(sch)
}

func (s *ScheduleStore) Merge(id model.ScheduleID, entry model.StretchScheduleEntry) error {
	prev, ok := s.mem.Schedule(id)
	if !ok {
		prev = model.StretchSchedule{ID: id}
	}
	prev.Entries = append(prev.Entries, entry)
	s.Upsert(prev)
	return nil
}

func (s *ScheduleStore) EntriesEndingBefore(t time.Time) []model.StretchScheduleEntry {
	var out []model.StretchScheduleEntry
	for _, sch := range s.allSchedules() {
		for _, e := range sch.Entries {
			if !e.End.After(t) {
				out = append(out, e)
			}
		}
	}
	return out
}

func (s *ScheduleStore) allSchedules() []model.StretchSchedule {
	s.mem.mu.RLock()
	defer s.mem.mu.RUnlock()
	out := make([]model.StretchSchedule, 0, len(s.mem.schedules))
	for _, sch := range s.mem.schedules {
		out = append(out, sch)
	}
	return out
}