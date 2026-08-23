package config

import (
	"time"

	"github.com/lacsar712/filmpull/internal/model"
)

type Config struct {
	LineID           string
	ZoneCount        int
	NipCount         int
	ProcessTickMs    int
	DefaultGrade     string
	TensionTolerance float64
	MaxDrawRatio     float64
	CoastDecelMPM    float64
	AlarmBufferSize  int
	ThreadingSpeed     float64
	SpeedRampSteps     int
	PreheatPlanSteps   int
	SetWindowMin       int
}

func Default() Config {
	return Config{
		LineID: "line-a", ZoneCount: 4, NipCount: 3, ProcessTickMs: 100,
		DefaultGrade: "FG-001", TensionTolerance: 8, MaxDrawRatio: 5.5,
		CoastDecelMPM: 2.0, AlarmBufferSize: 64, ThreadingSpeed: 15,
		SpeedRampSteps: 40, PreheatPlanSteps: 60, SetWindowMin: 5,
	}
}

func (c Config) ProcessTick() time.Duration {
	if c.ProcessTickMs <= 0 {
		return 100 * time.Millisecond
	}
	return time.Duration(c.ProcessTickMs) * time.Millisecond
}

func (c Config) Validate() error {
	if c.ZoneCount < 1 || c.ZoneCount > 12 {
		return model.Wrap("config", "zone_count", model.ErrInvalid)
	}
	if c.NipCount < 1 || c.NipCount > 8 {
		return model.Wrap("config", "nip_count", model.ErrInvalid)
	}
	if _, ok := model.LookupGrade(model.FilmGradeID(c.DefaultGrade)); !ok {
		return model.Wrap("config", "default_grade", model.ErrNotFound)
	}
	return nil
}