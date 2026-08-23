package store

import (
	"time"

	"github.com/lacsar712/filmpull/internal/model"
)

type SnapshotBuilder struct {
	id        model.LineID
	state     model.LineState
	zones     []model.StretchZone
	nips      []model.NipRoll
	speed     float64
	draw      float64
	grade     model.FilmGradeID
}

func NewSnapshotBuilder(id model.LineID) *SnapshotBuilder {
	return &SnapshotBuilder{id: id, state: model.LineIdle, grade: model.DefaultGrade().ID}
}

func (b *SnapshotBuilder) State(s model.LineState) *SnapshotBuilder {
	b.state = s
	return b
}

func (b *SnapshotBuilder) Zone(z model.StretchZone) *SnapshotBuilder {
	b.zones = append(b.zones, z)
	return b
}

func (b *SnapshotBuilder) Nip(n model.NipRoll) *SnapshotBuilder {
	b.nips = append(b.nips, n)
	return b
}

func (b *SnapshotBuilder) Speed(mpm float64) *SnapshotBuilder {
	b.speed = mpm
	return b
}

func (b *SnapshotBuilder) Draw(r float64) *SnapshotBuilder {
	b.draw = r
	return b
}

func (b *SnapshotBuilder) Grade(g model.FilmGradeID) *SnapshotBuilder {
	b.grade = g
	return b
}

func (b *SnapshotBuilder) Build(at time.Time) model.LineSnapshot {
	zones := make([]model.StretchZone, len(b.zones))
	copy(zones, b.zones)
	nips := make([]model.NipRoll, len(b.nips))
	copy(nips, b.nips)
	return model.LineSnapshot{
		ID: b.id, State: b.state, Zones: zones, Nips: nips,
		SpeedMPM: b.speed, DrawRatio: b.draw, Grade: b.grade, UpdatedAt: at,
	}
}

func DiffZones(before, after model.LineSnapshot) []model.ZoneID {
	index := make(map[model.ZoneID]model.StretchZone)
	for _, z := range before.Zones {
		index[z.ID] = z
	}
	var changed []model.ZoneID
	for _, z := range after.Zones {
		prev, ok := index[z.ID]
		if !ok || prev.SpeedMPM != z.SpeedMPM || prev.Enabled != z.Enabled || prev.DrawRatio != z.DrawRatio {
			changed = append(changed, z.ID)
		}
	}
	return changed
}