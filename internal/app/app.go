package app

import (
	"github.com/dek0valev/niwa/internal/config"
	"github.com/dek0valev/niwa/internal/content"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
	meta "github.com/yuin/goldmark-meta"
	"log/slog"
	"os"
)

type App struct {
	cfg *config.Config
	log *slog.Logger
}

func NewApp(cfg *config.Config, log *slog.Logger) *App {
	md := goldmark.New(
		goldmark.WithExtensions(meta.Meta, highlighting.NewHighlighting(
			highlighting.WithStyle("github-dark"),
		)),
	)

	store := content.NewStore()

	parser := content.NewParser(md)
	if err := parser.ParseDirectory(cfg.Content.ArticlesDir, store); err != nil {
		log.Error("Не удалось прочитать директорию: " + err.Error())
		os.Exit(1)
	}

	return &App{
		cfg: cfg,
		log: log,
	}
}

func (a *App) Run() {
	a.log.Info("Запуск приложения")
}
