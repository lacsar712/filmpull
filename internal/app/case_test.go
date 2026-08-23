package app

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/lacsar712/filmpull/internal/config"
	"github.com/lacsar712/filmpull/internal/model"
)

func TestCase(t *testing.T) {
	a, err := New(config.Default())
	if err != nil {
		t.Fatal(err)
	}
	anchor := a.clk.Now().Add(-10 * time.Minute)
	err = a.ConfirmSetHold(context.Background(), anchor)
	if err == nil {
		t.Fatal("expected set hold error")
	}
	if !errors.Is(err, model.ErrSetHold) {
		t.Fatalf("expected ErrSetHold, got %v", err)
	}
}
