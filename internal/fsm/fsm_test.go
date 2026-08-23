package fsm

import (
	"context"
	"testing"

	"github.com/lacsar712/filmpull/internal/model"
)

func TestLineFSMThreading(t *testing.T) {
	f := NewLineFSM("line-a", nil)
	ctx := context.Background()
	if err := f.Apply(ctx, "thread"); err != nil {
		t.Fatal(err)
	}
	if f.State() != model.LineThreading {
		t.Fatalf("state %s", f.State())
	}
}

func TestIllegalTransition(t *testing.T) {
	f := NewLineFSM("line-a", nil)
	if err := f.Apply(context.Background(), "stretch"); err == nil {
		t.Fatal("expected error from idle")
	}
}

func TestNipCloseSequence(t *testing.T) {
	f := NewNipFSM("nip-1", nil)
	ctx := context.Background()
	if err := f.Apply(ctx, "close"); err != nil {
		t.Fatal(err)
	}
	if err := f.Apply(ctx, "seated"); err != nil {
		t.Fatal(err)
	}
	if f.State() != model.NipClosed {
		t.Fatalf("state %s", f.State())
	}
}