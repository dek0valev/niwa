package main

import (
	"flag"
	"github.com/dek0valev/niwa/internal/config"
	"github.com/dek0valev/niwa/pkg/logger"
	"github.com/dek0valev/niwa/pkg/must"
	"log/slog"
)

func main() {
	var configPath string

	flag.StringVar(&configPath, "config", "", "path to config file")
	flag.Parse()

	cfg := must.Must(config.NewConfig(configPath))

	log := logger.NewLogger(cfg.Env)
	log.Info("庭 | Niwa", slog.String("env", cfg.Env))
}
