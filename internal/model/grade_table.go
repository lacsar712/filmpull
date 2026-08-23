package model

import "sort"

// GradeTable indexes film grades for lookup by thickness and speed capability.
type GradeTable struct {
	byID        map[FilmGradeID]FilmGrade
	byThickness map[int][]FilmGrade
	sorted      []FilmGrade
}

func NewGradeTable(profiles []FilmGrade) *GradeTable {
	t := &GradeTable{
		byID:        make(map[FilmGradeID]FilmGrade, len(profiles)),
		byThickness: make(map[int][]FilmGrade),
		sorted:      make([]FilmGrade, len(profiles)),
	}
	copy(t.sorted, profiles)
	sort.Slice(t.sorted, func(i, j int) bool {
		return t.sorted[i].ThicknessUM < t.sorted[j].ThicknessUM
	})
	for _, g := range profiles {
		t.byID[g.ID] = g
		t.byThickness[g.ThicknessUM] = append(t.byThickness[g.ThicknessUM], g)
	}
	return t
}

func DefaultGradeTable() *GradeTable {
	return NewGradeTable(BuiltinProfiles())
}

func (t *GradeTable) Lookup(id FilmGradeID) (FilmGrade, bool) {
	g, ok := t.byID[id]
	return g, ok
}

func (t *GradeTable) ByThickness(um int) []FilmGrade {
	out := t.byThickness[um]
	if len(out) == 0 {
		return nil
	}
	cp := make([]FilmGrade, len(out))
	copy(cp, out)
	return cp
}

func (t *GradeTable) BestForSpeed(minMPM float64) (FilmGrade, bool) {
	var best FilmGrade
	found := false
	for _, g := range t.sorted {
		if float64(g.MaxLineSpeedMPM) < minMPM {
			continue
		}
		if !found || g.MaxLineSpeedMPM < best.MaxLineSpeedMPM {
			best = g
			found = true
		}
	}
	return best, found
}

func (t *GradeTable) NearestThickness(um int) (FilmGrade, bool) {
	if len(t.sorted) == 0 {
		return FilmGrade{}, false
	}
	best := t.sorted[0]
	bestDist := absInt(best.ThicknessUM - um)
	for _, g := range t.sorted[1:] {
		d := absInt(g.ThicknessUM - um)
		if d < bestDist {
			best = g
			bestDist = d
		}
	}
	return best, true
}

func (t *GradeTable) FilterDrawRange(minDraw, maxDraw float64) []FilmGrade {
	var out []FilmGrade
	for _, g := range t.sorted {
		if g.DrawRatio >= minDraw && g.DrawRatio <= maxDraw {
			out = append(out, g)
		}
	}
	return out
}

func (t *GradeTable) Count() int { return len(t.sorted) }

func (t *GradeTable) All() []FilmGrade {
	out := make([]FilmGrade, len(t.sorted))
	copy(out, t.sorted)
	return out
}

func absInt(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
