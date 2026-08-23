# filmpull — project notes

Film stretching line controller for biaxial orientation (BOPET/BOPP) plants.

## Domain

- **Unwind / threading**: low-speed nip sequencing with tension hold windows
- **Preheat**: zone enablement before draw
- **Stretch**: multi-zone draw ratio with per-zone tension regulation
- **Anneal / coast**: controlled deceleration with interlocks

## Architecture

```
cmd/filmpull ──► internal/app
                    ├── stretch (zones, draw, speed planner)
                    ├── tension (sensors, regulator, profiles)
                    ├── nip (rollers, pressure, coordinator)
                    ├── fsm (line + nip state machines)
                    ├── clock (process vs wall)
                    ├── interlock (speed/tension/nip guards)
                    ├── store (snapshots, schedules)
                    └── stats (samples, gauges)
```

## Film grades

Built-in profiles (`FG-001` … `FG-080`) cover thickness, web width, draw ratio,
nominal tension, and max line speed for recipe selection.
