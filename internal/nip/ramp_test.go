package nip

import (
	"testing"
	"time"

	"github.com/lacsar712/filmpull/internal/model"
)

func TestPressureRampProfile(t *testing.T) {
	prof := NewPressureRampProfile("nip-1", 100, 0.85, 10)
	prof.Begin(time.Now())
	for i := 0; i < 50; i++ {
		prof.Tick(100 * time.Millisecond)
	}
	if prof.Current() <= 0 {
		t.Fatal("expected pressure")
	}
}

func TestRampBank(t *testing.T) {
	bank := NewRampBank()
	id := model.NipID("nip-1")
	bank.Start(id, 80, time.Now())
	if !bank.Active(id) {
		t.Fatal("expected active ramp")
	}
	for i := 0; i < 30; i++ {
		_, ok := bank.Advance(id, 100*time.Millisecond)
		if !ok {
			break
		}
	}
	phases := bank.Snapshot()
	if len(phases) > 0 && phases[id] == RampRelease {
		t.Fatalf("unexpected release %v", phases)
	}
}
