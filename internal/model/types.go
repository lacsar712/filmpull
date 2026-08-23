package model

import "time"

type LineState string

const (
	LineIdle      LineState = "idle"
	LineThreading LineState = "threading"
	LinePreheat   LineState = "preheat"
	LineStretch   LineState = "stretch"
	LineAnneal    LineState = "anneal"
	LineCoast     LineState = "coast"
	LineFault     LineState = "fault"
	LineShutdown  LineState = "shutdown"
)

type NipState string

const (
	NipOpen    NipState = "open"
	NipClosing NipState = "closing"
	NipClosed  NipState = "closed"
	NipFault   NipState = "fault"
)

type TensionSetpoint struct {
	Newtons     float64
	TolerancePct float64
}

func (t TensionSetpoint) Within(actual float64) bool {
	if t.Newtons <= 0 {
		return actual <= 0
	}
	lo := t.Newtons * (1 - t.TolerancePct/100)
	hi := t.Newtons * (1 + t.TolerancePct/100)
	return actual >= lo && actual <= hi
}

type SpeedSetpoint struct {
	MPM         float64
	RampMPMPerS float64
}

type StretchZone struct {
	ID         ZoneID
	DrawRatio  float64
	SpeedMPM   float64
	Tension    TensionSetpoint
	Enabled    bool
	LastTension float64
}

type NipRoll struct {
	ID        NipID
	State     NipState
	PressureKPa float64
	MaxPressure float64
	SpeedMPM  float64
}

type TensionReading struct {
	Sensor SensorID
	Newtons float64
	At      time.Time
}

type LineSnapshot struct {
	ID        LineID
	State     LineState
	Zones     []StretchZone
	Nips      []NipRoll
	SpeedMPM  float64
	DrawRatio float64
	Grade     FilmGradeID
	UpdatedAt time.Time
}

type StretchScheduleEntry struct {
	ID       ScheduleID
	Zone     ZoneID
	Start    time.Time
	End      time.Time
	Draw     float64
	Tension  TensionSetpoint
}

type StretchSchedule struct {
	ID      ScheduleID
	Entries []StretchScheduleEntry
	Version int64
}

func (s StretchSchedule) Clone() StretchSchedule {
	out := StretchSchedule{ID: s.ID, Version: s.Version}
	if len(s.Entries) == 0 {
		return out
	}
	out.Entries = make([]StretchScheduleEntry, len(s.Entries))
	copy(out.Entries, s.Entries)
	return out
}

type AlarmEvent struct {
	Code     AlarmCode
	Message  string
	Line     LineID
	RaisedAt time.Time
	Severity int
}

type FilmGrade struct {
	ID              FilmGradeID
	ThicknessUM     int
	WebWidthMM      int
	DrawRatio         float64
	NominalTensionN float64
	MaxLineSpeedMPM float64
	AnnealTempC     int
	PreheatTempC    int
}

type ZoneRoute struct {
	From     ZoneID
	To       ZoneID
	Nip      NipID
	Priority int
}