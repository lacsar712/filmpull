package app

import (
	"context"
	"fmt"
	"time"

	"github.com/lacsar712/filmpull/internal/clock"
	"github.com/lacsar712/filmpull/internal/config"
	"github.com/lacsar712/filmpull/internal/fsm"
	"github.com/lacsar712/filmpull/internal/interlock"
	"github.com/lacsar712/filmpull/internal/model"
	"github.com/lacsar712/filmpull/internal/nip"
	"github.com/lacsar712/filmpull/internal/stats"
	"github.com/lacsar712/filmpull/internal/store"
	"github.com/lacsar712/filmpull/internal/stretch"
	"github.com/lacsar712/filmpull/internal/tension"
)

type App struct {
	cfg      config.Config
	clk      *clock.ProcessClock
	mem      *store.Memory
	sched    *store.ScheduleStore
	lineFSM  *fsm.LineFSM
	zones    *stretch.ZoneTable
	stretch  *stretch.Controller
	sensors  *tension.SensorBank
	tension  *tension.Regulator
	nips     *nip.Coordinator
	guard    *interlock.Guard
	speed    *interlock.SpeedLock
	stats    *stats.Collector
	registry *stats.Registry
	grade    model.FilmGrade
	lineID   model.LineID
	speedMPM float64
}

func New(cfg config.Config) (*App, error) {
	if err := cfg.Validate(); err != nil {
		return nil, err
	}
	lineID, err := model.ParseLineID(cfg.LineID)
	if err != nil {
		return nil, err
	}
	grade, ok := model.LookupGrade(model.FilmGradeID(cfg.DefaultGrade))
	if !ok {
		return nil, model.Wrap("app", "grade", model.ErrNotFound)
	}
	start := time.Date(2024, 8, 1, 0, 0, 0, 0, time.UTC)
	clk := clock.NewProcessClock(start, cfg.ProcessTick())
	mem := store.NewMemory()
	zones, err := stretch.NewZoneTable(cfg.ZoneCount, grade)
	if err != nil {
		return nil, err
	}
	nipIDs := make([]model.NipID, cfg.NipCount)
	for i := 0; i < cfg.NipCount; i++ {
		nipIDs[i] = model.NipID(fmt.Sprintf("nip-%d", i+1))
	}
	a := &App{
		cfg: cfg, clk: clk, mem: mem, sched: store.NewScheduleStore(mem, clk),
		zones: zones, stretch: stretch.NewController(zones, clk, cfg.MaxDrawRatio),
		sensors: tension.NewSensorBank(), tension: nil,
		nips: nip.NewCoordinator(nipIDs, 600),
		guard: interlock.NewGuard(buildGuardMap(zones.Zones(), nipIDs)),
		speed: interlock.NewSpeedLock(float64(grade.MaxLineSpeedMPM)),
		stats: stats.NewCollector(), registry: stats.NewRegistry(),
		grade: grade, lineID: lineID,
	}
	a.tension = tension.NewRegulator(a.sensors, clk)
	a.lineFSM = fsm.NewLineFSM(lineID, a.onLineTransition)
	a.persistSnapshot()
	return a, nil
}

func buildGuardMap(zones []model.StretchZone, nips []model.NipID) map[model.ZoneID]model.NipID {
	m := make(map[model.ZoneID]model.NipID)
	for i, z := range zones {
		nipIdx := i % len(nips)
		m[z.ID] = nips[nipIdx]
	}
	return m
}

func (a *App) onLineTransition(ctx context.Context, line model.LineID, from, to model.LineState) error {
	a.stats.IncEvent("line_" + string(to))
	if to == model.LineFault {
		a.registry.Set("fault_active", 1)
	}
	if to == model.LineStretch {
		a.tension.BeginHold(5*a.cfg.ProcessTick(), model.TensionSetpoint{Newtons: a.grade.NominalTensionN, TolerancePct: a.cfg.TensionTolerance})
	}
	_ = ctx
	_ = line
	_ = from
	return nil
}

func (a *App) persistSnapshot() {
	b := store.NewSnapshotBuilder(a.lineID).State(a.lineFSM.State()).Grade(a.grade.ID).Speed(a.speedMPM)
	for _, z := range a.zones.Zones() {
		b.Zone(z)
	}
	for _, n := range a.nips.Snapshots() {
		b.Nip(n)
	}
	b.Draw(stretch.EffectiveDraw(a.zones.Zones()))
	a.mem.PutSnapshot(b.Build(a.clk.Now()))
}

func (a *App) Thread(ctx context.Context) error {
	if err := a.lineFSM.Apply(ctx, "thread"); err != nil {
		return err
	}
	a.speedMPM = a.cfg.ThreadingSpeed
	a.persistSnapshot()
	return nil
}

func (a *App) StartStretch(ctx context.Context) error {
	if err := a.lineFSM.Apply(ctx, "preheat"); err != nil {
		return err
	}
	if err := a.lineFSM.Apply(ctx, "stretch"); err != nil {
		return err
	}
	target := float64(a.grade.MaxLineSpeedMPM) * 0.8
	if err := a.rampSpeed(ctx, target); err != nil {
		return err
	}
	a.persistSnapshot()
	return nil
}

func (a *App) rampSpeed(ctx context.Context, target float64) error {
	nipSnaps := a.nips.Snapshots()
	if err := a.speed.PermitSpeedChange(a.speedMPM, target, nipSnaps); err != nil {
		return err
	}
	a.speedMPM = target
	a.registry.Set("line_speed_mpm", target)
	a.stats.IncEvent("speed_change")
	return nil
}

func (a *App) IngestTension(zone model.ZoneID, sensor model.SensorID, n float64) error {
	a.sensors.Ingest(model.TensionReading{Sensor: sensor, Newtons: n, At: a.clk.Now()})
	a.stats.RecordTension(zone, n, a.clk.Now())
	return a.zones.RecordTension(zone, n)
}

func (a *App) CloseNip(ctx context.Context, id model.NipID, kpa float64) error {
	zones := a.guard.ZonesFor(id)
	for _, z := range zones {
		if err := a.guard.Permit(z, id); err != nil {
			return err
		}
	}
	return a.nips.Close(ctx, id, kpa)
}

func (a *App) Coast(ctx context.Context) error {
	if err := a.lineFSM.Apply(ctx, "coast"); err != nil {
		return err
	}
	for a.speedMPM > 0 {
		select {
		case <-ctx.Done():
			return context.Cause(ctx)
		default:
		}
		a.speedMPM -= a.cfg.CoastDecelMPM * 0.1
		if a.speedMPM < 0 {
			a.speedMPM = 0
		}
		a.clk.Step()
	}
	a.persistSnapshot()
	return nil
}

func (a *App) Snapshot() (model.LineSnapshot, bool) {
	return a.mem.Snapshot(a.lineID)
}

func (a *App) LineState() model.LineState { return a.lineFSM.State() }
func (a *App) Stats() *stats.Collector    { return a.stats }
func (a *App) Registry() *stats.Registry  { return a.registry }