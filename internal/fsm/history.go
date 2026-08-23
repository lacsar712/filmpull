package fsm

import "github.com/lacsar712/filmpull/internal/model"

// History records recent line state transitions for diagnostics export.
type History struct {
	entries []TransitionRecord
	limit   int
}

type TransitionRecord struct {
	From  model.LineState
	To    model.LineState
	Event string
}

func NewHistory(limit int) *History {
	if limit <= 0 {
		limit = 32
	}
	return &History{limit: limit}
}

func (h *History) Record(from, to model.LineState, event string) {
	h.entries = append(h.entries, TransitionRecord{From: from, To: to, Event: event})
	if len(h.entries) > h.limit {
		h.entries = h.entries[len(h.entries)-h.limit:]
	}
}

func (h *History) Entries() []TransitionRecord {
	out := make([]TransitionRecord, len(h.entries))
	copy(out, h.entries)
	return out
}

func (h *History) Last() (TransitionRecord, bool) {
	if len(h.entries) == 0 {
		return TransitionRecord{}, false
	}
	return h.entries[len(h.entries)-1], true
}
