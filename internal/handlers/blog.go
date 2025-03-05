package handlers

import (
	"github.com/dek0valev/niwa/internal/content"
	"html/template"
	"net/http"
)

type BlogHandler struct {
	store *content.Store
}

func NewBlogHandler(store *content.Store) *BlogHandler {
	return &BlogHandler{
		store: store,
	}
}

func (h *BlogHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles(
		"web/templates/layouts/base/base.gohtml",
		"web/templates/pages/blog/blog.gohtml",
	))

	articles := h.store.Articles()

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := tmpl.ExecuteTemplate(w, "base", articles); err != nil {
		http.Error(w, "Не удалось отрисовать шаблон", http.StatusInternalServerError)
		return
	}
}
