package tension

import (
	"context"
	"time"

	"github.com/lacsar712/filmpull/internal/clock"
	"github.com/lacsar712/filmpull/internal/model"
)

type Regulator struct {
	bank   *SensorBank
	clk    clock.Clock
	window HoldWindow
}

func NewRegulator(bank *SensorBank, clk clock.Clock) *Regulator {
	return &Regulator{bank: bank, clk: clk}
}

func (r *Regulator) BeginHold(duration time.Duration, sp model.TensionSetpoint) {
	r.window = NewHoldWindow(r.clk, duration, sp)
}

func (r *Regulator) Evaluate(ctx context.Context, zone model.ZoneID, sensors []model.SensorID) (float64, error) {
	select {
	case <-ctx.Done():
		return 0, context.Cause(ctx)
	default:
	}
	avg, ok := r.bank.Average(zone, sensors)
	if !ok {
		return 0, model.Wrap("tension", "no_reading", model.ErrNotFound)
	}
	if r.window.setpoint.Newtons > 0 && !r.window.Satisfied(r.clk, avg) && r.window.Active(r.clk) {
		return avg, model.Wrap("tension", "hold_window", model.ErrTension)
	}
	return avg, nil
}

func (r *Regulator) InBand(sp model.TensionSetpoint, actual float64) bool {
	return sp.Within(actual)
}

func (r *Regulator) TrimOutput(sp model.TensionSetpoint, actual, output float64) float64 {
	delta := sp.Newtons - actual
	if delta > 0 {
		return output + delta*0.1
	}
	return output + delta*0.1
}