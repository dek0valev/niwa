package content

import (
	"bytes"
	"fmt"
	"github.com/dek0valev/niwa/internal/models"
	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Parser struct {
	md goldmark.Markdown
}

func NewParser(md goldmark.Markdown) *Parser {
	return &Parser{md: md}
}

func (p *Parser) ParseDirectory(dirPath string, store *Store) error {
	const op = "content.parser.ParseDirectory"

	newArticles := make(map[string]models.Article)

	err := filepath.WalkDir(dirPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() || !strings.HasSuffix(d.Name(), ".md") {
			return nil
		}

		slug := strings.TrimSuffix(filepath.Base(path), filepath.Ext(path))

		article, err := p.parseFile(path, slug)
		if err != nil {
			return fmt.Errorf("%s: %w", op, err)
		}

		newArticles[slug] = article
		return nil
	})

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	store.UpdateArticles(newArticles)

	return nil
}

func (p *Parser) parseFile(filePath, slug string) (models.Article, error) {
	const op = "content.parser.parseFile"

	f, err := os.ReadFile(filePath)
	if err != nil {
		return models.Article{}, fmt.Errorf("%s: %w", op, err)
	}

	var buf bytes.Buffer

	context := parser.NewContext()

	if err := p.md.Convert(f, &buf, parser.WithContext(context)); err != nil {
		return models.Article{}, fmt.Errorf("%s: %w", op, err)
	}

	metaData := meta.Get(context)

	publishedAt, err := time.Parse("2006-01-02", metaData["published_at"].(string))
	if err != nil {
		return models.Article{}, fmt.Errorf("%s: %w", op, err)
	}

	modifiedAt, err := time.Parse("2006-01-02", metaData["modified_at"].(string))
	if err != nil {
		return models.Article{}, fmt.Errorf("%s: %w", op, err)
	}

	categories := make([]string, 0)

	rawCategories := metaData["categories"].([]any)
	for _, rawCategory := range rawCategories {
		categories = append(categories, rawCategory.(string))
	}

	article := models.Article{
		Slug:        slug,
		Title:       metaData["title"].(string),
		Description: metaData["description"].(string),
		IsDraft:     metaData["is_draft"].(bool),
		Categories:  categories,
		PublishedAt: publishedAt,
		ModifiedAt:  modifiedAt,
		Content:     buf.Bytes(),
	}

	return article, nil
}
