package interlock

import "github.com/lacsar712/filmpull/internal/model"

type GripGuard struct {
	threshold float64
}

func NewGripGuard(threshold float64) *GripGuard {
	return &GripGuard{threshold: threshold}
}

func (g *GripGuard) Permit(slipPct float64) error {
	if slipPct > g.threshold {
		return model.Wrap("grip_guard", "slip", model.ErrGripSlip)
	}
	return nil
}
