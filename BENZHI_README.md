# filmpull

filmpull 是一套薄膜拉伸线工业过程控制系统，用于放卷张力、拉伸比、压辊压力与线速联锁。

## Requirements

- Go 1.22+ (container image uses golang:1.22)
- `GOTOOLCHAIN=local` recommended on host when using a pinned toolchain

## Build

```bash
export GOTOOLCHAIN=local
go build ./...
```

## Run

```bash
export GOTOOLCHAIN=local
go run ./cmd/filmpull
```

Open http://127.0.0.1:8080/

Default ingest secret: `dev-ingest-secret`.

## Test

```bash
export GOTOOLCHAIN=local
go test ./... -count=1
```

## Docker (benzhi)

Must build **linux/amd64** and **linux/arm64**:

```bash
chmod +x build_benzhi_docker.sh
./build_benzhi_docker.sh filmpull linux/amd64
./build_benzhi_docker.sh filmpull linux/arm64
docker run -it filmpull:latest
# inside container:
export GOTOOLCHAIN=local
go version
go build ./...
go test ./... -count=1
```
