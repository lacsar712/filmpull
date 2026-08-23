package store

import "github.com/lacsar712/filmpull/internal/model"

type ProfileSnapshot struct {
	Line  model.LineID
	Steps []float64
}

func CloneProfileSnapshot(s ProfileSnapshot) ProfileSnapshot {
	out := ProfileSnapshot{Line: s.Line}
	out.Steps = make([]float64, len(s.Steps))
	copy(out.Steps, s.Steps)
	return out
}
