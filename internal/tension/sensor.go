package tension

import (
	"time"

	"github.com/lacsar712/filmpull/internal/model"
)

type SensorBank struct {
	readings map[model.SensorID]model.TensionReading
}

func NewSensorBank() *SensorBank {
	return &SensorBank{readings: make(map[model.SensorID]model.TensionReading)}
}

func (b *SensorBank) Ingest(r model.TensionReading) {
	b.readings[r.Sensor] = r
}

func (b *SensorBank) Latest(id model.SensorID) (model.TensionReading, bool) {
	r, ok := b.readings[id]
	return r, ok
}

func (b *SensorBank) Average(zone model.ZoneID, sensors []model.SensorID) (float64, bool) {
	if len(sensors) == 0 {
		return 0, false
	}
	var sum float64
	var count int
	for _, sid := range sensors {
		r, ok := b.readings[sid]
		if !ok {
			continue
		}
		sum += r.Newtons
		count++
	}
	if count == 0 {
		return 0, false
	}
	_ = zone
	return sum / float64(count), true
}

func (b *SensorBank) Stale(before time.Time) []model.SensorID {
	var out []model.SensorID
	for id, r := range b.readings {
		if r.At.Before(before) {
			out = append(out, id)
		}
	}
	return out
}