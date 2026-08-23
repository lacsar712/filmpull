package stretch

import (
	"fmt"

	"github.com/lacsar712/filmpull/internal/model"
)

type ZoneTable struct {
	zones []model.StretchZone
}

func NewZoneTable(count int, grade model.FilmGrade) (*ZoneTable, error) {
	if count < 1 {
		return nil, model.Wrap("stretch", "zone_count", model.ErrInvalid)
	}
	zones := make([]model.StretchZone, count)
	for i := 0; i < count; i++ {
		id := model.ZoneID(fmt.Sprintf("zone-%d", i+1))
		ratio := grade.DrawRatio * float64(i+1) / float64(count)
		zones[i] = model.StretchZone{
			ID: id, DrawRatio: ratio, SpeedMPM: 0,
			Tension: model.TensionSetpoint{Newtons: grade.NominalTensionN, TolerancePct: 8},
			Enabled: i == 0,
		}
	}
	return &ZoneTable{zones: zones}, nil
}

func (z *ZoneTable) Zones() []model.StretchZone {
	out := make([]model.StretchZone, len(z.zones))
	copy(out, z.zones)
	return out
}

func (z *ZoneTable) UpdateSpeed(id model.ZoneID, mpm float64) error {
	for i := range z.zones {
		if z.zones[i].ID == id {
			z.zones[i].SpeedMPM = mpm
			return nil
		}
	}
	return model.Wrap("stretch", "update_speed", model.ErrNotFound)
}

func (z *ZoneTable) SetDraw(id model.ZoneID, ratio float64) error {
	for i := range z.zones {
		if z.zones[i].ID == id {
			z.zones[i].DrawRatio = ratio
			return nil
		}
	}
	return model.Wrap("stretch", "set_draw", model.ErrNotFound)
}

func (z *ZoneTable) Enable(id model.ZoneID, on bool) error {
	for i := range z.zones {
		if z.zones[i].ID == id {
			z.zones[i].Enabled = on
			return nil
		}
	}
	return model.Wrap("stretch", "enable", model.ErrNotFound)
}

func (z *ZoneTable) RecordTension(id model.ZoneID, n float64) error {
	for i := range z.zones {
		if z.zones[i].ID == id {
			z.zones[i].LastTension = n
			return nil
		}
	}
	return model.Wrap("stretch", "record_tension", model.ErrNotFound)
}