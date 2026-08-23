package stretch

import (
	"github.com/lacsar712/filmpull/internal/model"
)

func EffectiveDraw(zones []model.StretchZone) float64 {
	if len(zones) == 0 {
		return 1
	}
	ratio := 1.0
	for _, z := range zones {
		if !z.Enabled || z.DrawRatio <= 0 {
			continue
		}
		ratio *= z.DrawRatio
	}
	return ratio
}

func SpeedAtDraw(inletMPM float64, draw float64) float64 {
	if draw <= 0 {
		return inletMPM
	}
	return inletMPM * draw
}

func ZoneSpeeds(inletMPM float64, zones []model.StretchZone) []float64 {
	out := make([]float64, len(zones))
	acc := inletMPM
	for i, z := range zones {
		if z.Enabled && z.DrawRatio > 0 {
			acc *= z.DrawRatio
		}
		out[i] = acc
	}
	return out
}