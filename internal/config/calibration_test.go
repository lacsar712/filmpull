package config

import (
	"testing"

	"github.com/lacsar712/filmpull/internal/model"
)

func TestDefaultCalibration(t *testing.T) {
	c := DefaultCalibration(4, 3)
	if len(c.ZoneFactors) != 4 || len(c.NipMaxKPa) != 3 {
		t.Fatalf("cal %+v", c)
	}
	scaled := c.ScaleZone(model.ZoneID("zone-1"), 100)
	if scaled < 100 {
		t.Fatalf("scaled %v", scaled)
	}
	if c.MaxNipPressure("nip-1") < 400 {
		t.Fatal()
	}
}
