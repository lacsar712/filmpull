package interlock

import (
	"github.com/lacsar712/filmpull/internal/model"
)

type Guard struct {
	allowed map[model.ZoneID]model.NipID
}

func NewGuard(pairs map[model.ZoneID]model.NipID) *Guard {
	cp := make(map[model.ZoneID]model.NipID, len(pairs))
	for k, v := range pairs {
		cp[k] = v
	}
	return &Guard{allowed: cp}
}

func (g *Guard) Permit(zone model.ZoneID, nip model.NipID) error {
	want, ok := g.allowed[zone]
	if !ok {
		return model.Wrap("interlock", "unknown_zone", model.ErrNotFound)
	}
	if want != nip {
		return model.Wrap("interlock", "nip_mismatch", model.ErrInterlock)
	}
	return nil
}

func (g *Guard) ZonesFor(nip model.NipID) []model.ZoneID {
	var out []model.ZoneID
	for zone, n := range g.allowed {
		if n == nip {
			out = append(out, zone)
		}
	}
	return out
}