.PHONY: build test benzhi-docker
build:
	go build -o bin/filmpull ./cmd/filmpull
test:
	go test ./... -count=1
benzhi-docker:
	sh build_benzhi_docker.sh