package handlers

import (
	"fmt"
	"github.com/dek0valev/niwa/internal/content"
	"net/http"
	"strings"
)

type RobotsHandler struct {
	store   *content.Store
	baseURL string
}

func NewRobotsHandler(store *content.Store, baseURL string) *RobotsHandler {
	return &RobotsHandler{
		store:   store,
		baseURL: baseURL,
	}
}

func (h *RobotsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var sb strings.Builder

	sb.WriteString("User-agent: *\n")
	sb.WriteString("Allow: /\n")
	sb.WriteString(fmt.Sprintf("Sitemap: %s", h.baseURL+"/sitemap.xml"))

	w.Header().Set("Content-Type", "text/plain")
	if _, err := w.Write([]byte(sb.String())); err != nil {
		http.Error(w, "Не удалось записать данные", http.StatusInternalServerError)
		return
	}
}
