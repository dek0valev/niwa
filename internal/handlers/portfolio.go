package handlers

import (
	"github.com/dek0valev/niwa/internal/content"
	"html/template"
	"net/http"
)

type PortfolioHandler struct {
	store *content.Store
}

func NewPortfolioHandler(store *content.Store) *PortfolioHandler {
	return &PortfolioHandler{
		store: store,
	}
}

func (h *PortfolioHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(
		"web/templates/layouts/base/base.gohtml",
		"web/templates/layouts/base/partials/header.gohtml",
		"web/templates/pages/portfolio/portfolio.gohtml",
	))

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := tmpl.ExecuteTemplate(w, "base", nil); err != nil {
		http.Error(w, "Не удалось отрисовать шаблон", http.StatusInternalServerError)
		return
	}
}
