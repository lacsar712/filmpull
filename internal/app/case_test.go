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
	err = a.ReportTensionFault(context.Background(), a.grade.NominalTensionN*2)
	if err == nil {
		t.Fatal("expected tension exceeded error")
	}
	if !errors.Is(err, model.ErrTensionExceeded) {
		t.Fatalf("expected ErrTensionExceeded, got %v", err)
	}
}
