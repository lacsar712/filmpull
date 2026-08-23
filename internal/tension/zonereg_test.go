package tension

import (
	"context"
	"testing"
	"time"

	"github.com/lacsar712/filmpull/internal/clock"
	"github.com/lacsar712/filmpull/internal/model"
)

func TestZoneRegulatorBind(t *testing.T) {
	grade := model.DefaultGrade()
	profile := NewProfile(grade, 3)
	clk := clock.NewProcessClock(time.Now(), 100*time.Millisecond)
	bank := NewSensorBank()
	reg := NewRegulator(bank, clk)
	zr := NewZoneRegulator(profile, reg, clk)
	zones := []model.StretchZone{
		{ID: "zone-1", Enabled: true},
		{ID: "zone-2", Enabled: true},
	}
	zr.BindAll(zones)
	if _, ok := zr.State("zone-1"); !ok {
		t.Fatal("zone not bound")
	}
}

func TestZoneRegulatorPhase(t *testing.T) {
	grade := model.DefaultGrade()
	profile := NewProfile(grade, 2)
	clk := clock.NewProcessClock(time.Now(), 100*time.Millisecond)
	zr := NewZoneRegulator(profile, NewRegulator(NewSensorBank(), clk), clk)
	zr.BindZone("zone-1", 0)
	zr.ApplyPhase("thread")
	st, _ := zr.State("zone-1")
	if st.Setpoint.Newtons <= 0 {
		t.Fatalf("setpoint %+v", st.Setpoint)
	}
}

func TestZoneRegulatorCascade(t *testing.T) {
	grade := model.DefaultGrade()
	profile := NewProfile(grade, 2)
	clk := clock.NewProcessClock(time.Now(), 100*time.Millisecond)
	bank := NewSensorBank()
	reg := NewRegulator(bank, clk)
	zr := NewZoneRegulator(profile, reg, clk)
	zones := []model.StretchZone{
		{ID: "zone-1", Enabled: true},
		{ID: "zone-2", Enabled: true},
	}
	zr.BindAll(zones)
	bank.Ingest(model.TensionReading{Sensor: "s1", Newtons: 150, At: clk.Now()})
	_ = zr.Regulate(context.Background(), "zone-1", []model.SensorID{"s1"})
	zr.CascadeTrim(zones)
}
