package handlers

import (
	"github.com/dek0valev/niwa/internal/content"
	"html/template"
	"net/http"
)

type ArticleHandler struct {
	store *content.Store
}

func NewArticleHandler(store *content.Store) *ArticleHandler {
	return &ArticleHandler{
		store: store,
	}
}

func (h *ArticleHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")

	preTmpl := template.New("base").Funcs(template.FuncMap{
		"byteToHTML": func(b []byte) template.HTML {
			return template.HTML(b)
		},
	})

	tmpl := template.Must(preTmpl.ParseFiles(
		"web/templates/layouts/base/base.gohtml",
		"web/templates/pages/article/article.gohtml",
	))

	article, err := h.store.ArticleBySlug(slug)
	if err != nil {
		http.Error(w, "Не удалось найти статью", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	if err := tmpl.ExecuteTemplate(w, "base", article); err != nil {
		http.Error(w, "Не удалось отрисовать шаблон", http.StatusInternalServerError)
		return
	}
}
