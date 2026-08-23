package store

import (
	"sync"

	"github.com/lacsar712/filmpull/internal/model"
)

type Memory struct {
	mu        sync.RWMutex
	snapshots map[model.LineID]model.LineSnapshot
	schedules map[model.ScheduleID]model.StretchSchedule
	version   int64
}

func NewMemory() *Memory {
	return &Memory{
		snapshots: make(map[model.LineID]model.LineSnapshot),
		schedules: make(map[model.ScheduleID]model.StretchSchedule),
	}
}

func (m *Memory) PutSnapshot(s model.LineSnapshot) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.snapshots[s.ID] = s
	m.version++
}

func (m *Memory) Snapshot(id model.LineID) (model.LineSnapshot, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	s, ok := m.snapshots[id]
	return s, ok
}

func (m *Memory) PutSchedule(s model.StretchSchedule) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.schedules[s.ID] = s
	m.version++
}

func (m *Memory) Schedule(id model.ScheduleID) (model.StretchSchedule, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	s, ok := m.schedules[id]
	return s, ok
}

func (m *Memory) Version() int64 {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return m.version
}