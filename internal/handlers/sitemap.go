package handlers

import (
	"encoding/xml"
	"fmt"
	"github.com/dek0valev/niwa/internal/content"
	"net/http"
	"time"
)

type SitemapHandler struct {
	store   *content.Store
	baseURL string
}

func NewSitemapHandler(store *content.Store, baseURL string) *SitemapHandler {
	return &SitemapHandler{
		store:   store,
		baseURL: baseURL,
	}
}

func (h *SitemapHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	sitemapXML := h.generateDynamicSitemap()

	w.Header().Set("Content-Type", "application/xml")
	if _, err := w.Write(sitemapXML); err != nil {
		http.Error(w, "Не удалось записать данные", http.StatusInternalServerError)
		return
	}
}

type URLSet struct {
	XMLName xml.Name `xml:"urlset"`
	Xmlns   string   `xml:"xmlns,attr"`
	URLs    []URL    `xml:"url"`
}

type URL struct {
	Loc     string `xml:"loc"`
	LastMod string `xml:"lastmod"`
}

type Page struct {
	URL        string
	ModifiedAt time.Time
}

var pages = []Page{
	{
		URL:        "/",
		ModifiedAt: time.Date(2025, 3, 6, 0, 0, 0, 0, time.UTC),
	},
	{
		URL:        "/blog",
		ModifiedAt: time.Date(2025, 3, 6, 0, 0, 0, 0, time.UTC),
	},
	{
		URL:        "/portfolio",
		ModifiedAt: time.Date(2025, 3, 6, 0, 0, 0, 0, time.UTC),
	},
}

func (h *SitemapHandler) generateDynamicSitemap() []byte {
	articles := h.store.Articles()

	urlset := URLSet{
		Xmlns: "http://www.sitemaps.org/schemas/sitemap/0.9",
		URLs:  []URL{},
	}

	for _, page := range pages {
		urlset.URLs = append(urlset.URLs, URL{
			Loc:     h.baseURL + page.URL,
			LastMod: page.ModifiedAt.Format("2006-01-02"),
		})
	}

	for _, article := range articles {
		urlset.URLs = append(urlset.URLs, URL{
			Loc:     h.baseURL + "/blog/" + article.Slug,
			LastMod: article.ModifiedAt.Format("2006-01-02"),
		})
	}

	output, err := xml.MarshalIndent(urlset, "", "  ")
	if err != nil {
		fmt.Println("Ошибка при генерации XML:", err)
		return []byte{}
	}

	return append([]byte(xml.Header), output...)
}
