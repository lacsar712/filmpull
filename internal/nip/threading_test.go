package nip

import (
	"testing"

	"github.com/lacsar712/filmpull/internal/model"
)

func TestThreadingSequence(t *testing.T) {
	seq := NewThreadingSequence([]model.NipID{"nip-1", "nip-2", "nip-3"})
	first, ok := seq.Next("")
	if !ok || first != "nip-1" {
		t.Fatal()
	}
	second, ok := seq.Next("nip-1")
	if !ok || second != "nip-2" {
		t.Fatal()
	}
	if _, ok := seq.Next("nip-3"); ok {
		t.Fatal("expected end")
	}
	if !seq.Contains("nip-2") {
		t.Fatal()
	}
}
