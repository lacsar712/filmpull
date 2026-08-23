package fsm

import (
	"testing"

	"github.com/lacsar712/filmpull/internal/model"
)

func TestHistoryRecord(t *testing.T) {
	h := NewHistory(2)
	h.Record(model.LineIdle, model.LineThreading, "thread")
	h.Record(model.LineThreading, model.LinePreheat, "preheat")
	h.Record(model.LinePreheat, model.LineStretch, "stretch")
	entries := h.Entries()
	if len(entries) != 2 {
		t.Fatalf("len %d", len(entries))
	}
	last, ok := h.Last()
	if !ok || last.To != model.LineStretch {
		t.Fatalf("last %+v", last)
	}
}
