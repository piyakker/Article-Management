package publishing

import (
	"errors"
	"sync"
)

type inMemoryArticleRepository struct {
	mu       sync.RWMutex
	articles map[string]*Article
}

func NewInMemoryArticleRepository() ArticleRepository {
	return &inMemoryArticleRepository{
		articles: make(map[string]*Article),
	}
}

func (r *inMemoryArticleRepository) GetAll() ([]*Article, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var allArticles []*Article
	for _, article := range r.articles {
		allArticles = append(allArticles, article)
	}
	return allArticles, nil
}

func (r *inMemoryArticleRepository) GetByID(id string) (*Article, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	article, exist := r.articles[id]
	if !exist {
		return nil, ErrArticleNotExist
	}
	return article, nil
}

func (r *inMemoryArticleRepository) Create(article *Article) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exist := r.articles[article.ID]; exist {
		return ErrArticleAlreadyExist
	}
	r.articles[article.ID] = article
	return nil
}

func (r *inMemoryArticleRepository) Update(article *Article) (string, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	oldArticle, exist := r.articles[article.ID]
	if !exist {
		return "Update Failed", ErrArticleNotExist
	}
	r.articles[oldArticle.ID] = article
	return "Update Successfully", nil
}

func (r *inMemoryArticleRepository) Delete(id string) (string, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	_, exist := r.articles[id]
	if !exist {
		return "Delete Failed", ErrArticleNotExist
	}
	delete(r.articles, id)
	return "Delete Successfully", nil
}

var (
	ErrArticleNotExist     = errors.New("article not exist")
	ErrArticleAlreadyExist = errors.New("article with the same ID already exists")
)
