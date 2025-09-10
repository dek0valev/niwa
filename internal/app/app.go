package app

import (
	"context"
	"github.com/dek0valev/niwa/internal/config"
	"github.com/dek0valev/niwa/internal/content"
	"github.com/dek0valev/niwa/internal/handlers"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/extension"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type App struct {
	cfg *config.Config
	log *slog.Logger
	hs  *http.Server
}

func NewApp(cfg *config.Config, log *slog.Logger) *App {
	md := goldmark.New(
		goldmark.WithExtensions(meta.Meta, extension.GFM, highlighting.NewHighlighting(
			highlighting.WithStyle("github-dark"),
		)),
	)

	store := content.NewStore()

	parser := content.NewParser(md)
	if err := parser.ParseDirectory(cfg.Content.ArticlesDir, store); err != nil {
		log.Error("Не удалось прочитать директорию: " + err.Error())
		os.Exit(1)
	}

	sitemapHandler := handlers.NewSitemapHandler(store, cfg.BaseURL)
	robotsHandler := handlers.NewRobotsHandler(store, cfg.BaseURL)

	homeHandler := handlers.NewHomeHandler(store)
	blogHandler := handlers.NewBlogHandler(store)
	articleHandler := handlers.NewArticleHandler(store)
	portfolioHandler := handlers.NewPortfolioHandler(store)

	hfs := http.FileServer(http.Dir("web/static"))

	r := http.NewServeMux()

	r.Handle("GET /static/", http.StripPrefix("/static/", hfs))

	r.Handle("GET /sitemap.xml", sitemapHandler)
	r.Handle("GET /robots.txt", robotsHandler)

	r.Handle("GET /{$}", homeHandler)
	r.Handle("GET /blog", blogHandler)
	r.Handle("GET /blog/{slug}", articleHandler)
	r.Handle("GET /portfolio", portfolioHandler)

	hs := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	return &App{
		cfg: cfg,
		log: log,
		hs:  hs,
	}
}

func (a *App) Run() {
	a.log.Info("Запуск приложения")

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := a.hs.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			a.log.Error("Не удалось запустить сервер: " + err.Error())
			os.Exit(1)
		}
	}()

	<-stopChan

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := a.hs.Shutdown(ctx); err != nil {
		a.log.Error("Не удалось остановить сервер: " + err.Error())
		os.Exit(1)
	}

	a.log.Info("Сервер успешно остановлен")
}
