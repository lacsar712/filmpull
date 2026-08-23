package tension

import (
	"github.com/lacsar712/filmpull/internal/model"
)

// Profile maps film grade nominal tension to zone-specific setpoints during stretch.
type Profile struct {
	grade model.FilmGrade
	zones int
}

func NewProfile(grade model.FilmGrade, zoneCount int) Profile {
	return Profile{grade: grade, zones: zoneCount}
}

func (p Profile) ZoneSetpoint(index int) model.TensionSetpoint {
	if index < 0 {
		index = 0
	}
	factor := 1.0 + float64(index)*0.05
	return model.TensionSetpoint{
		Newtons:      p.grade.NominalTensionN * factor,
		TolerancePct: 8,
	}
}

func (p Profile) All() []model.TensionSetpoint {
	out := make([]model.TensionSetpoint, p.zones)
	for i := 0; i < p.zones; i++ {
		out[i] = p.ZoneSetpoint(i)
	}
	return out
}

func (p Profile) PreheatHold() model.TensionSetpoint {
	return model.TensionSetpoint{Newtons: p.grade.NominalTensionN * 0.6, TolerancePct: 12}
}

func (p Profile) ThreadingHold() model.TensionSetpoint {
	return model.TensionSetpoint{Newtons: p.grade.NominalTensionN * 0.3, TolerancePct: 15}
}
