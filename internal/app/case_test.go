package app

import (
	"context"
	"errors"
	"testing"

	"github.com/lacsar712/filmpull/internal/config"
	"github.com/lacsar712/filmpull/internal/model"
)

func TestCase(t *testing.T) {
	a, err := New(config.Default())
	if err != nil {
		t.Fatal(err)
	}
	err = a.HandleGripSlip(context.Background(), 25.0)
	if err == nil {
		t.Fatal("expected grip slip error")
	}
	if !errors.Is(err, model.ErrGripSlip) {
		t.Fatalf("expected ErrGripSlip, got %v", err)
	}
}
