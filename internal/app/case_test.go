package app

import (
	"context"
	"fmt"
	"testing"

	"github.com/lacsar712/filmpull/internal/config"
	"github.com/lacsar712/filmpull/internal/model"
)

func TestCase(t *testing.T) {
	a, err := New(config.Default())
	if err != nil {
		t.Fatal(err)
	}
	nip := model.NipID("nip-1")
	CalibrateProbe = func(ctx context.Context) error {
		return fmt.Errorf("tension probe fault")
	}
	if err := a.CalibrateGrip(context.Background(), nip, "operator-a"); err == nil {
		t.Fatal("expected calibration probe failure")
	}
	CalibrateProbe = nil
	if err := a.CalibrateGrip(context.Background(), nip, "operator-b"); err != nil {
		t.Fatalf("second calibration blocked by leaked grip lease: %v", err)
	}
}
