package tension

import (
	"context"
	"testing"
	"time"

	"github.com/lacsar712/filmpull/internal/clock"
	"github.com/lacsar712/filmpull/internal/model"
)

func TestSensorBankAverage(t *testing.T) {
	bank := NewSensorBank()
	s1 := model.FormatSensor("zone-a", 1)
	s2 := model.FormatSensor("zone-a", 2)
	now := time.Now()
	bank.Ingest(model.TensionReading{Sensor: s1, Newtons: 100, At: now})
	bank.Ingest(model.TensionReading{Sensor: s2, Newtons: 120, At: now})
	avg, ok := bank.Average("zone-a", []model.SensorID{s1, s2})
	if !ok || avg != 110 {
		t.Fatalf("avg %v", avg)
	}
}

func TestHoldWindow(t *testing.T) {
	start := time.Date(2024, 5, 1, 0, 0, 0, 0, time.UTC)
	clk := clock.NewProcessClock(start, time.Second)
	w := NewHoldWindow(clk, 5*time.Second, model.TensionSetpoint{Newtons: 100, TolerancePct: 10})
	clk.Advance(2 * time.Second)
	if !w.Active(clk) {
		t.Fatal("window should be active")
	}
	if !w.Satisfied(clk, 105) {
		t.Fatal("should satisfy setpoint")
	}
}

func TestRegulatorEvaluate(t *testing.T) {
	bank := NewSensorBank()
	clk := clock.NewProcessClock(time.Now(), 10*time.Millisecond)
	reg := NewRegulator(bank, clk)
	sid := model.FormatSensor("zone-a", 1)
	bank.Ingest(model.TensionReading{Sensor: sid, Newtons: 90, At: clk.Now()})
	_, err := reg.Evaluate(context.Background(), "zone-a", []model.SensorID{sid})
	if err != nil {
		t.Fatal(err)
	}
}