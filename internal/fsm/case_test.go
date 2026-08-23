package fsm

import (
	"context"
	"testing"

	"github.com/lacsar712/filmpull/internal/model"
)

func TestCase(t *testing.T) {
	GripDrivePulse = nil
	var pulses int
	GripDrivePulse = func() { pulses++ }
	line := NewLineFSM(model.LineID("line-a"), nil)
	if err := line.Apply(context.Background(), "stretch"); err == nil {
		t.Fatal("expected illegal line transition error")
	}
	if pulses != 0 {
		t.Fatalf("illegal line transition should not pulse grip drive, got %d", pulses)
	}
}
