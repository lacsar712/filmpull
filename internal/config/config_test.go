package config

import "testing"

func TestDefaultValid(t *testing.T) {
	if err := Default().Validate(); err != nil {
		t.Fatal(err)
	}
}

func TestInvalidZoneCount(t *testing.T) {
	cfg := Default()
	cfg.ZoneCount = 0
	if err := cfg.Validate(); err == nil {
		t.Fatal("expected error")
	}
}