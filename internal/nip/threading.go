package nip

import (
	"github.com/lacsar712/filmpull/internal/model"
)

// ThreadingSequence defines the order nip rolls close during web insertion.
type ThreadingSequence struct {
	order []model.NipID
}

func NewThreadingSequence(ids []model.NipID) ThreadingSequence {
	cp := make([]model.NipID, len(ids))
	copy(cp, ids)
	return ThreadingSequence{order: cp}
}

func (s ThreadingSequence) Order() []model.NipID {
	out := make([]model.NipID, len(s.order))
	copy(out, s.order)
	return out
}

func (s ThreadingSequence) Next(after model.NipID) (model.NipID, bool) {
	if len(s.order) == 0 {
		return "", false
	}
	if after == "" {
		return s.order[0], true
	}
	for i, id := range s.order {
		if id == after && i+1 < len(s.order) {
			return s.order[i+1], true
		}
	}
	return "", false
}

func (s ThreadingSequence) Contains(id model.NipID) bool {
	for _, n := range s.order {
		if n == id {
			return true
		}
	}
	return false
}
