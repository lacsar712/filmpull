package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/lacsar712/filmpull/internal/app"
	"github.com/lacsar712/filmpull/internal/config"
)

func main() {
	cfgPath := flag.String("config", "", "optional TOML config path")
	tickMs := flag.Int("tick-ms", 100, "process clock tick in milliseconds")
	flag.Parse()
	cfg := config.Default()
	cfg.ProcessTickMs = *tickMs
	if *cfgPath != "" {
		loaded, err := config.LoadFile(*cfgPath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "config: %v\n", err)
			os.Exit(2)
		}
		cfg = loaded
	}
	a, err := app.New(cfg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "init: %v\n", err)
		os.Exit(2)
	}
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	if err := a.Run(ctx); err != nil && err != context.Canceled {
		fmt.Fprintf(os.Stderr, "run: %v\n", err)
		os.Exit(1)
	}
}