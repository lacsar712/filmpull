package model

import "testing"

func TestTensionSetpointWithin(t *testing.T) {
	sp := TensionSetpoint{Newtons: 100, TolerancePct: 10}
	if !sp.Within(105) {
		t.Fatal("expected within band")
	}
	if sp.Within(120) {
		t.Fatal("expected outside band")
	}
}

func TestStretchScheduleClone(t *testing.T) {
	orig := StretchSchedule{ID: "sch-1", Version: 2, Entries: []StretchScheduleEntry{{ID: "e1", Zone: "z1"}}}
	cl := orig.Clone()
	cl.Entries[0].ID = "mutated"
	if orig.Entries[0].ID == "mutated" {
		t.Fatal("clone should not alias entries")
	}
}

func TestLookupGrade(t *testing.T) {
	g, ok := LookupGrade("FG-001")
	if !ok || g.ID != "FG-001" {
		t.Fatalf("lookup failed: %+v", g)
	}
}