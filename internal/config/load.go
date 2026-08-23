package config

import (
	"encoding/json"
	"os"
)

type fileShape struct {
	LineID           string  `json:"line_id"`
	ZoneCount        int     `json:"zone_count"`
	NipCount         int     `json:"nip_count"`
	ProcessTickMs    int     `json:"process_tick_ms"`
	DefaultGrade     string  `json:"default_grade"`
	TensionTolerance float64 `json:"tension_tolerance"`
	MaxDrawRatio     float64 `json:"max_draw_ratio"`
	CoastDecelMPM    float64 `json:"coast_decel_mpm"`
	AlarmBufferSize  int     `json:"alarm_buffer_size"`
	ThreadingSpeed   float64 `json:"threading_speed"`
}

func LoadFile(path string) (Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}
	var raw fileShape
	if err := json.Unmarshal(data, &raw); err != nil {
		return Config{}, err
	}
	cfg := Default()
	if raw.LineID != "" {
		cfg.LineID = raw.LineID
	}
	if raw.ZoneCount > 0 {
		cfg.ZoneCount = raw.ZoneCount
	}
	if raw.NipCount > 0 {
		cfg.NipCount = raw.NipCount
	}
	if raw.ProcessTickMs > 0 {
		cfg.ProcessTickMs = raw.ProcessTickMs
	}
	if raw.DefaultGrade != "" {
		cfg.DefaultGrade = raw.DefaultGrade
	}
	if raw.TensionTolerance > 0 {
		cfg.TensionTolerance = raw.TensionTolerance
	}
	if raw.MaxDrawRatio > 0 {
		cfg.MaxDrawRatio = raw.MaxDrawRatio
	}
	if raw.CoastDecelMPM > 0 {
		cfg.CoastDecelMPM = raw.CoastDecelMPM
	}
	if raw.AlarmBufferSize > 0 {
		cfg.AlarmBufferSize = raw.AlarmBufferSize
	}
	if raw.ThreadingSpeed > 0 {
		cfg.ThreadingSpeed = raw.ThreadingSpeed
	}
	return cfg, cfg.Validate()
}