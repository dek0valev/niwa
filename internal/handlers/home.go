package handlers

import (
	"github.com/dek0valev/niwa/internal/content"
	"html/template"
	"net/http"
)

type HomeHandler struct {
	store *content.Store
}

func NewHomeHandler(store *content.Store) *HomeHandler {
	return &HomeHandler{
		store: store,
	}
}

func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(
		"web/templates/layouts/base/base.gohtml",
		"web/templates/layouts/base/partials/header.gohtml",
		"web/templates/pages/home/home.gohtml",
	))

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := tmpl.ExecuteTemplate(w, "base", nil); err != nil {
		http.Error(w, "Не удалось отрисовать шаблон", http.StatusInternalServerError)
		return
	}
}
