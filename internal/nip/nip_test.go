package nip

import (
	"context"
	"testing"

	"github.com/lacsar712/filmpull/internal/model"
)

func TestRampPressure(t *testing.T) {
	got := RampPressure(10, 20, 3)
	if got != 13 {
		t.Fatalf("got %v", got)
	}
}

func TestCoordinatorClose(t *testing.T) {
	c := NewCoordinator([]model.NipID{"nip-1"}, 500)
	if err := c.Close(context.Background(), "nip-1", 100); err != nil {
		t.Fatal(err)
	}
	snaps := c.Snapshots()
	if len(snaps) != 1 || snaps[0].State != model.NipClosed {
		t.Fatalf("snap %+v", snaps)
	}
}