package publishing

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"github.com/gorilla/mux"
)

var ErrBadRouting = errors.New("inconsistent mapping between route and handler (programmer error)")

func NewHttpHandler(endpoints Endpoints, logger log.Logger) http.Handler {
	m := mux.NewRouter()

	options := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
		httptransport.ServerErrorEncoder(encodeErrorResponse),
	}

	GetAllArticlesHandler := httptransport.NewServer(
		endpoints.GetAllArticlesEndpoint,
		decodeGetAllArticlesRequest,
		encodeResponse,
		options...,
	)

	GetArticleByIDHandler := httptransport.NewServer(
		endpoints.GetArticleByIDEndpoint,
		decodeGetArticleByIDRequest,
		encodeResponse,
		options...,
	)

	CreateArticleHandler := httptransport.NewServer(
		endpoints.CreateArticleEndpoint,
		decodeCreateArticleRequest,
		encodeResponse,
		options...,
	)

	UpdateArticleHandler := httptransport.NewServer(
		endpoints.UpdateArticleEndpoint,
		decodeUpdateArticleTRequest,
		encodeResponse,
		options...,
	)

	DeleteArticleHandler := httptransport.NewServer(
		endpoints.DeleteArticleEndpoint,
		decodeDeleteArticleRequest,
		encodeResponse,
		options...,
	)

	m.Handle("/Articles", GetAllArticlesHandler).Methods("GET")
	m.Handle("/Articles/{id}", GetArticleByIDHandler).Methods("GET")
	m.Handle("/Articles", CreateArticleHandler).Methods("POST")
	m.Handle("/Articles/{id}", UpdateArticleHandler).Methods("PUT")
	m.Handle("/Articles/{id}", DeleteArticleHandler).Methods("DELETE")
	return m
}

func encodeErrorResponse(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(codeFrom(err))
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func codeFrom(err error) int {
	switch err {
	case ErrArticleNotExist:
		return http.StatusNotFound
	case ErrArticleAlreadyExist:
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}

func decodeGetArticleByIDRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	return GetArticleByIDRequest{ID: id}, nil
}

func decodeCreateArticleRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request CreateArticleRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeGetAllArticlesRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return GetAllArticlesRequest{}, nil
}

func decodeUpdateArticleTRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}

	var request UpdateArticleRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	request.ID = id
	return request, nil
}

func decodeDeleteArticleRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}

	return DeleteArticleRequest{ID: id}, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
