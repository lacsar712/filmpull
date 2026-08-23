# filmpull

Film stretching line controller for biaxial orientation plants. Coordinates unwind tension,
stretch-zone draw ratio, nip roll pressure, and line speed with process-clock driven
interlocks and state machines.

## Build

```bash
go build -o bin/filmpull ./cmd/filmpull
go test ./... -count=1
```

## Packages

- `internal/stretch` 鈥?draw ratio and zone speed coordination
- `internal/tension` 鈥?load-cell regulation and hold windows
- `internal/nip` 鈥?nip roll pressure and threading sequence
- `internal/fsm` 鈥?line and nip state machines
- `internal/clock` 鈥?process clock vs wall clock
- `internal/interlock` 鈥?speed/tension/nip guards
- `internal/store` 鈥?in-memory snapshots and schedules
- `internal/stats` 鈥?runtime counters and tension histograms
- `internal/app` 鈥?orchestration entry point