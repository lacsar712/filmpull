package stats

import (
	"sync"
	"time"

	"github.com/lacsar712/filmpull/internal/model"
)

type Sample struct {
	Zone    model.ZoneID
	Newtons float64
	At      time.Time
}

type Collector struct {
	mu      sync.Mutex
	samples []Sample
	events  map[string]int64
}

func NewCollector() *Collector {
	return &Collector{events: make(map[string]int64)}
}

func (c *Collector) RecordTension(zone model.ZoneID, n float64, at time.Time) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.samples = append(c.samples, Sample{Zone: zone, Newtons: n, At: at})
}

func (c *Collector) IncEvent(name string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.events[name]++
}

func (c *Collector) Count(name string) int64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.events[name]
}

func (c *Collector) Samples() []Sample {
	c.mu.Lock()
	defer c.mu.Unlock()
	out := make([]Sample, len(c.samples))
	copy(out, c.samples)
	return out
}

func (c *Collector) MeanTension(zone model.ZoneID) (float64, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	var sum float64
	var count int
	for _, s := range c.samples {
		if s.Zone != zone {
			continue
		}
		sum += s.Newtons
		count++
	}
	if count == 0 {
		return 0, false
	}
	return sum / float64(count), true
}