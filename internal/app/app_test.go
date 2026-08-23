package app

import (
	"context"
	"testing"

	"github.com/lacsar712/filmpull/internal/config"
	"github.com/lacsar712/filmpull/internal/model"
)

func TestAppThreadAndStretch(t *testing.T) {
	cfg := config.Default()
	a, err := New(cfg)
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()
	if err := a.Thread(ctx); err != nil {
		t.Fatal(err)
	}
	if a.LineState() != model.LineThreading {
		t.Fatalf("state %s", a.LineState())
	}
	nipID := model.NipID("nip-1")
	if err := a.CloseNip(ctx, nipID, 80); err != nil {
		t.Fatal(err)
	}
	if err := a.StartStretch(ctx); err != nil {
		t.Fatal(err)
	}
	if a.LineState() != model.LineStretch {
		t.Fatalf("state %s", a.LineState())
	}
	snap, ok := a.Snapshot()
	if !ok || snap.SpeedMPM <= 0 {
		t.Fatalf("snap %+v", snap)
	}
}

func TestAppIngestTension(t *testing.T) {
	a, err := New(config.Default())
	if err != nil {
		t.Fatal(err)
	}
	zone := model.ZoneID("zone-1")
	sid := model.FormatSensor(zone, 1)
	if err := a.IngestTension(zone, sid, 110); err != nil {
		t.Fatal(err)
	}
	mean, ok := a.Stats().MeanTension(zone)
	if !ok || mean != 110 {
		t.Fatalf("mean %v", mean)
	}
}

func TestAppCoast(t *testing.T) {
	a, err := New(config.Default())
	if err != nil {
		t.Fatal(err)
	}
	ctx := context.Background()
	_ = a.Thread(ctx)
	_ = a.CloseNip(ctx, "nip-1", 50)
	_ = a.StartStretch(ctx)
	if err := a.Coast(ctx); err != nil {
		t.Fatal(err)
	}
	if a.LineState() != model.LineCoast {
		t.Fatalf("state %s", a.LineState())
	}
}