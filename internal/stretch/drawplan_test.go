package stretch

import (
	"math"
	"testing"

	"github.com/lacsar712/filmpull/internal/model"
)

func TestDrawPlannerPlan(t *testing.T) {
	p, err := NewDrawPlanner(4, 5.5)
	if err != nil {
		t.Fatal(err)
	}
	grade := model.DefaultGrade()
	plan, err := p.Plan(grade)
	if err != nil {
		t.Fatal(err)
	}
	if len(plan.Steps) != 4 {
		t.Fatalf("steps %d", len(plan.Steps))
	}
	acc := 1.0
	for _, s := range plan.Steps {
		acc *= s.DrawRatio
	}
	if math.Abs(acc-plan.TotalDraw) > 0.01 {
		t.Fatalf("draw %v want %v", acc, plan.TotalDraw)
	}
}

func TestDrawPlannerIncremental(t *testing.T) {
	p, _ := NewDrawPlanner(3, 5.5)
	grade := model.DefaultGrade()
	plan, err := p.PlanIncremental(grade, 2)
	if err != nil {
		t.Fatal(err)
	}
	if len(plan.Steps) != 2 {
		t.Fatalf("steps %d", len(plan.Steps))
	}
}

func TestSpeedCoordinatorBalance(t *testing.T) {
	c := NewSpeedCoordinator(50, 2)
	grade := model.DefaultGrade()
	table, _ := NewZoneTable(3, grade)
	zones := table.Zones()
	setpoints := c.Balance(zones, 120)
	if len(setpoints) != 3 {
		t.Fatalf("setpoints %d", len(setpoints))
	}
}

func TestSpeedCoordinatorRampStep(t *testing.T) {
	c := NewSpeedCoordinator(10, 2)
	got := c.RampStep(10, 20)
	if got != 12 {
		t.Fatalf("ramp %v", got)
	}
}
