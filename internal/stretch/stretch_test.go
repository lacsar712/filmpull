package stretch

import (
	"math"
	"testing"
	"time"

	"github.com/lacsar712/filmpull/internal/clock"
	"github.com/lacsar712/filmpull/internal/model"
)

func TestEffectiveDraw(t *testing.T) {
	zones := []model.StretchZone{
		{Enabled: true, DrawRatio: 1.2},
		{Enabled: true, DrawRatio: 1.5},
		{Enabled: false, DrawRatio: 2.0},
	}
	got := EffectiveDraw(zones)
	want := 1.2 * 1.5
	if math.Abs(got-want) > 1e-9 {
		t.Fatalf("draw %v want %v", got, want)
	}
}

func TestZoneTableUpdate(t *testing.T) {
	grade := model.DefaultGrade()
	table, err := NewZoneTable(3, grade)
	if err != nil {
		t.Fatal(err)
	}
	zones := table.Zones()
	if err := table.UpdateSpeed(zones[0].ID, 100); err != nil {
		t.Fatal(err)
	}
}

func TestControllerApplyDraw(t *testing.T) {
	grade := model.DefaultGrade()
	table, _ := NewZoneTable(2, grade)
	clk := clock.NewProcessClock(time.Now(), 10*time.Millisecond)
	ctrl := NewController(table, clk, 5.0)
	zones := table.Zones()
	if err := ctrl.ApplyDraw(zones[0].ID, 3.5); err != nil {
		t.Fatal(err)
	}
	if err := ctrl.ApplyDraw(zones[0].ID, 9.0); err == nil {
		t.Fatal("expected draw limit error")
	}
}