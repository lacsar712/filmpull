package model

import "testing"

func TestGradeTableLookup(t *testing.T) {
	tbl := DefaultGradeTable()
	if tbl.Count() < 10 {
		t.Fatalf("count %d", tbl.Count())
	}
	g, ok := tbl.Lookup("FG-001")
	if !ok || g.ID != "FG-001" {
		t.Fatalf("lookup %+v", g)
	}
}

func TestGradeTableBestForSpeed(t *testing.T) {
	tbl := DefaultGradeTable()
	g, ok := tbl.BestForSpeed(200)
	if !ok {
		t.Fatal("expected grade")
	}
	if g.MaxLineSpeedMPM < 200 {
		t.Fatalf("speed %v", g.MaxLineSpeedMPM)
	}
}

func TestGradeTableNearestThickness(t *testing.T) {
	tbl := DefaultGradeTable()
	g, ok := tbl.NearestThickness(20)
	if !ok {
		t.Fatal("expected grade")
	}
	if g.ThicknessUM < 10 {
		t.Fatalf("thickness %d", g.ThicknessUM)
	}
}

func TestGradeTableFilterDraw(t *testing.T) {
	tbl := DefaultGradeTable()
	out := tbl.FilterDrawRange(3.0, 3.5)
	if len(out) == 0 {
		t.Fatal("expected matches")
	}
}
