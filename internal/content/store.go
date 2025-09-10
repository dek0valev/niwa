package content

import (
	"fmt"
	"github.com/dek0valev/niwa/internal/models"
	"sort"
	"sync"
)

type Store struct {
	mu       sync.RWMutex
	articles map[string]models.Article
}

func NewStore() *Store {
	return &Store{
		articles: make(map[string]models.Article),
	}
}

func (s *Store) Articles() []models.Article {
	s.mu.RLock()
	defer s.mu.RUnlock()

	articles := make([]models.Article, 0, len(s.articles))
	for _, article := range s.articles {
		articles = append(articles, article)
	}

	sort.Slice(articles, func(i, j int) bool {
		return articles[i].PublishedAt.After(articles[j].PublishedAt)
	})

	return articles
}

func (s *Store) ArticleBySlug(slug string) (models.Article, error) {
	const op = "content.store.ArticleBySlug"

	s.mu.RLock()
	defer s.mu.RUnlock()

	article, ok := s.articles[slug]
	if !ok {
		return models.Article{}, fmt.Errorf("%s: не удалось найти статью", op)
	}

	return article, nil
}

func (s *Store) UpdateArticles(articles map[string]models.Article) {
	s.mu.Lock()
	defer s.mu.Unlock()

	newArticles := make(map[string]models.Article, len(articles))

	for slug, article := range articles {
		newArticles[slug] = article
	}

	s.articles = newArticles
}
