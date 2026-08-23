package stats

import "sync"

type Registry struct {
	mu    sync.RWMutex
	gauge map[string]float64
}

func NewRegistry() *Registry {
	return &Registry{gauge: make(map[string]float64)}
}

func (r *Registry) Set(name string, v float64) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.gauge[name] = v
}

func (r *Registry) Get(name string) (float64, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	v, ok := r.gauge[name]
	return v, ok
}

func (r *Registry) Snapshot() map[string]float64 {
	r.mu.RLock()
	defer r.mu.RUnlock()
	out := make(map[string]float64, len(r.gauge))
	for k, v := range r.gauge {
		out[k] = v
	}
	return out
}