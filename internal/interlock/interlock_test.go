package interlock

import (
	"testing"

	"github.com/lacsar712/filmpull/internal/model"
)

func TestGuardPermit(t *testing.T) {
	g := NewGuard(map[model.ZoneID]model.NipID{"zone-a": "nip-1"})
	if err := g.Permit("zone-a", "nip-1"); err != nil {
		t.Fatal(err)
	}
	if err := g.Permit("zone-a", "nip-2"); err == nil {
		t.Fatal("expected mismatch")
	}
}

func TestSpeedLockNipClosed(t *testing.T) {
	lock := NewSpeedLock(300)
	nips := []model.NipRoll{{State: model.NipOpen, SpeedMPM: 5}}
	if err := lock.PermitSpeedChange(0, 50, nips); err == nil {
		t.Fatal("expected interlock")
	}
	nips[0].State = model.NipClosed
	if err := lock.PermitSpeedChange(0, 50, nips); err != nil {
		t.Fatal(err)
	}
}