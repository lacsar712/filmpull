package store

import (
	"testing"

	"github.com/lacsar712/filmpull/internal/model"
)

func TestCase(t *testing.T) {
	orig := ProfileSnapshot{
		Line:  model.LineID("line-a"),
		Steps: []float64{120.0, 130.0},
	}
	clone := CloneProfileSnapshot(orig)
	clone.Steps[0] = 99.0
	if orig.Steps[0] == 99.0 {
		t.Fatal("clone mutated original profile Steps backing array")
	}
}
