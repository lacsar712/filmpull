package stretch

import (
	"math"

	"github.com/lacsar712/filmpull/internal/model"
)

// DrawPlan captures per-zone draw targets for a stretch run.
type DrawPlan struct {
	TotalDraw float64
	Steps     []DrawStep
}

type DrawStep struct {
	Zone      model.ZoneID
	DrawRatio float64
	Enable    bool
	Order     int
}

// DrawPlanner distributes a grade draw ratio across stretch zones.
type DrawPlanner struct {
	maxDraw   float64
	zoneCount int
}

func NewDrawPlanner(zoneCount int, maxDraw float64) (*DrawPlanner, error) {
	if zoneCount < 1 {
		return nil, model.Wrap("stretch", "draw_plan", model.ErrInvalid)
	}
	if maxDraw <= 1 {
		maxDraw = 5.5
	}
	return &DrawPlanner{maxDraw: maxDraw, zoneCount: zoneCount}, nil
}

func (p *DrawPlanner) Plan(grade model.FilmGrade) (DrawPlan, error) {
	target := grade.DrawRatio
	if target <= 1 {
		return DrawPlan{}, model.Wrap("stretch", "draw_target", model.ErrInvalid)
	}
	if target > p.maxDraw {
		return DrawPlan{}, model.Wrap("stretch", "draw_cap", model.ErrInvalid)
	}
	perZone := math.Pow(target, 1.0/float64(p.zoneCount))
	steps := make([]DrawStep, p.zoneCount)
	for i := 0; i < p.zoneCount; i++ {
		id := model.ZoneID("zone-" + itoa(i+1))
		steps[i] = DrawStep{
			Zone:      id,
			DrawRatio: roundDraw(perZone),
			Enable:    i == 0,
			Order:     i,
		}
	}
	return DrawPlan{TotalDraw: target, Steps: steps}, nil
}

func (p *DrawPlanner) PlanIncremental(grade model.FilmGrade, enabledZones int) (DrawPlan, error) {
	full, err := p.Plan(grade)
	if err != nil {
		return DrawPlan{}, err
	}
	if enabledZones < 1 {
		enabledZones = 1
	}
	if enabledZones > len(full.Steps) {
		enabledZones = len(full.Steps)
	}
	partial := make([]DrawStep, enabledZones)
	acc := 1.0
	for i := 0; i < enabledZones; i++ {
		step := full.Steps[i]
		step.Enable = true
		partial[i] = step
		acc *= step.DrawRatio
	}
	return DrawPlan{TotalDraw: acc, Steps: partial}, nil
}

func (p *DrawPlanner) Validate(zones []model.StretchZone) error {
	draw := EffectiveDraw(zones)
	if draw > p.maxDraw {
		return model.Wrap("stretch", "validate_draw", model.ErrInvalid)
	}
	for _, z := range zones {
		if z.DrawRatio > p.maxDraw {
			return model.Wrap("stretch", "zone_draw", model.ErrInvalid)
		}
	}
	return nil
}

func (p *DrawPlanner) Apply(table *ZoneTable, plan DrawPlan) error {
	for _, step := range plan.Steps {
		if err := table.SetDraw(step.Zone, step.DrawRatio); err != nil {
			return err
		}
		if err := table.Enable(step.Zone, step.Enable); err != nil {
			return err
		}
	}
	return nil
}

func (p *DrawPlanner) NextEnableStep(zones []model.StretchZone) (model.ZoneID, bool) {
	for i, z := range zones {
		if !z.Enabled {
			return zones[i].ID, true
		}
	}
	return "", false
}

func roundDraw(v float64) float64 {
	return math.Round(v*1000) / 1000
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	digits := make([]byte, 0, 4)
	for n > 0 {
		digits = append([]byte{byte('0' + n%10)}, digits...)
		n /= 10
	}
	return string(digits)
}
