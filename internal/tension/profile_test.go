package tension

import (
	"testing"

	"github.com/lacsar712/filmpull/internal/model"
)

func TestProfileZoneSetpoints(t *testing.T) {
	grade := model.DefaultGrade()
	p := NewProfile(grade, 4)
	all := p.All()
	if len(all) != 4 {
		t.Fatal()
	}
	if all[1].Newtons <= all[0].Newtons {
		t.Fatalf("expected increasing tension %+v", all)
	}
	pre := p.PreheatHold()
	if pre.Newtons >= grade.NominalTensionN {
		t.Fatal()
	}
}
