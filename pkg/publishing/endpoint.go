package publishing

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	GetArticleByIDEndpoint endpoint.Endpoint
	CreateArticleEndpoint  endpoint.Endpoint
	GetAllArticlesEndpoint endpoint.Endpoint
	UpdateArticleEndpoint  endpoint.Endpoint
	DeleteArticleEndpoint  endpoint.Endpoint
}

func MakeEndpoints(s ArticleService) Endpoints {
	return Endpoints{
		GetArticleByIDEndpoint: makeGetArticleByIDEndpoint(s),
		CreateArticleEndpoint:  makeCreateArticleEndpoint(s),
		GetAllArticlesEndpoint: makeGetAllArticlesEndpoint(s),
		UpdateArticleEndpoint:  makeUpdateArticleEndpoint(s),
		DeleteArticleEndpoint:  makeDeleteArticleEndpoint(s),
	}
}

type GetAllArticlesRequest struct{}

type GetAllArticlesResponse struct {
	Articles []*Article `json:"articles"`
	Err      string     `json:"error,omitempty"`
}

func makeGetAllArticlesEndpoint(s ArticleService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		_ = request.(GetAllArticlesRequest)
		articles, err := s.GetAllArticles()
		if err != nil {
			return GetAllArticlesResponse{articles, err.Error()}, err
		}
		return GetAllArticlesResponse{articles, ""}, nil
	}
}

type GetArticleByIDRequest struct {
	ID string `json:"id"`
}

// GetArticleByIDResponse represents the response for the GetArticleByID endpoint.
type GetArticleByIDResponse struct {
	Article *Article `json:"article"`
	Err     string   `json:"error,omitempty"`
}

func makeGetArticleByIDEndpoint(s ArticleService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetArticleByIDRequest)
		article, err := s.GetArticleByID(req.ID)
		if err != nil {
			return GetArticleByIDResponse{article, err.Error()}, err
		}
		return GetArticleByIDResponse{article, ""}, nil
	}
}

type CreateArticleRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

type CreateArticleResponse struct {
	ID  string `json:"id"`
	Err string `json:"error,omitempty"`
}

func makeCreateArticleEndpoint(s ArticleService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateArticleRequest)
		id, err := s.CreateArticle(req.Title, req.Content, req.Author)
		if err != nil {
			return CreateArticleResponse{id, err.Error()}, nil
		}
		return CreateArticleResponse{id, ""}, nil
	}
}

type UpdateArticleRequest struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

type UpdateArticleResponse struct {
	Msg string `json:"msg"`
	Err string `json:"error,omitempty"`
}

func makeUpdateArticleEndpoint(s ArticleService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(UpdateArticleRequest)
		msg, err := s.UpdateArticle(req.ID, req.Title, req.Content, req.Author)
		if err != nil {
			return UpdateArticleResponse{msg, err.Error()}, err
		}
		return UpdateArticleResponse{msg, ""}, nil
	}
}

type DeleteArticleRequest struct {
	ID string `json:"id"`
}

type DeleteArticleResponse struct {
	Msg string `json:"msg"`
	Err string `json:"error,omitempty"`
}

func makeDeleteArticleEndpoint(s ArticleService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(DeleteArticleRequest)
		msg, err := s.DeleteArticle(req.ID)
		if err != nil {
			return DeleteArticleResponse{msg, err.Error()}, err
		}
		return DeleteArticleResponse{msg, ""}, nil
	}
}
