package stretch

import (
	"testing"

	"github.com/lacsar712/filmpull/internal/model"
)

func TestLineSpeedPlanner(t *testing.T) {
	planner := NewLineSpeedPlanner(100)
	zones := []model.StretchZone{
		{Enabled: true, DrawRatio: 1.2},
		{Enabled: true, DrawRatio: 1.3},
	}
	plan := planner.Plan(zones)
	if len(plan) != 2 || plan[1].MPM <= plan[0].MPM {
		t.Fatalf("plan %+v", plan)
	}
	outlet := planner.OutletMPM(zones)
	if outlet != 100*1.2*1.3 {
		t.Fatalf("outlet %v", outlet)
	}
}
