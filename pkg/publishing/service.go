package publishing

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type ArticleService interface {
	GetAllArticles() ([]*Article, error)
	GetArticleByID(id string) (*Article, error)
	CreateArticle(title, content, author string) (string, error)
	UpdateArticle(id, title, content, author string) (string, error)
	DeleteArticle(id string) (string, error)
}

type articleService struct {
	repo ArticleRepository
}

func NewArticleService(repo ArticleRepository) ArticleService {
	return &articleService{repo}
}

func (s articleService) GetAllArticles() ([]*Article, error) {
	return s.repo.GetAll()
}

func (s articleService) CreateArticle(title, content, author string) (string, error) {

	if title == "" || content == "" || author == "" {
		return "", errors.New("title, content, and author are required")
	}

	article := &Article{
		ID:      generateUniqueID(),
		Title:   title,
		Content: content,
		Author:  author,
	}

	err := s.repo.Create(article)
	if err != nil {
		return "", err
	}

	return article.ID, nil
}

func (s articleService) GetArticleByID(id string) (*Article, error) {
	article, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	return article, nil
}

func (s articleService) UpdateArticle(id, title, content, author string) (string, error) {
	return s.repo.Update(&Article{id, title, content, author})
}

func (s articleService) DeleteArticle(id string) (string, error) {
	return s.repo.Delete(id)
}

func generateUniqueID() string {
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	randomNumber := rand.Intn(1000) // Adjust the range as needed.

	return fmt.Sprintf("%d-%d", timestamp, randomNumber)
}
