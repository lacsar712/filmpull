package model

import (
	"fmt"
	"strings"
)

type LineID string
type ZoneID string
type NipID string
type RollID string
type SensorID string
type ScheduleID string
type AlarmCode string
type FilmGradeID string

func ParseLineID(raw string) (LineID, error) {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return "", Wrap("model", "parse_line", ErrInvalid)
	}
	return LineID(raw), nil
}

func ParseZoneID(raw string) (ZoneID, error) {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return "", Wrap("model", "parse_zone", ErrInvalid)
	}
	return ZoneID(raw), nil
}

func ParseNipID(raw string) (NipID, error) {
	raw = strings.TrimSpace(raw)
	if raw == "" {
		return "", Wrap("model", "parse_nip", ErrInvalid)
	}
	return NipID(raw), nil
}

func (id LineID) String() string { return string(id) }
func (id ZoneID) String() string { return string(id) }
func (id NipID) String() string  { return string(id) }

func FormatSensor(zone ZoneID, idx int) SensorID {
	return SensorID(fmt.Sprintf("%s-tension-%d", zone, idx))
}