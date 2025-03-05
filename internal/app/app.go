package app

import (
	"github.com/dek0valev/niwa/internal/config"
	"log/slog"
)

type App struct {
	cfg *config.Config
	log *slog.Logger
}

func NewApp(cfg *config.Config, log *slog.Logger) *App {
	return &App{
		cfg: cfg,
		log: log,
	}
}

func (a *App) Run() {
	a.log.Info("Запуск приложения")
}
