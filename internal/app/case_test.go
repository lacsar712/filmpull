package app

import (
	"context"
	"testing"

	"github.com/lacsar712/filmpull/internal/config"
)

func TestCase(t *testing.T) {
	a, err := New(config.Default())
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	if err := a.SpeedRamp(ctx, float64(a.grade.MaxLineSpeedMPM)*0.5); err == nil {
		t.Fatal("expected cancel during speed ramp")
	}
}
