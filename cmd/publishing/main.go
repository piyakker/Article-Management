package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-kit/log"
	"github.com/piyakker/ArticleManagement/pkg/publishing"
)

func main() {
	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stderr)
	logger = log.With(logger, "listen", "8080", "caller", log.DefaultCaller)

	inMemoryRepo := publishing.NewInMemoryArticleRepository()
	repo := inMemoryRepo

	articleService := publishing.NewArticleService(repo)
	articleService = publishing.NewLoggingMiddleware(logger, articleService)
	endpoints := publishing.MakeEndpoints(articleService)
	handler := publishing.NewHttpHandler(endpoints, logger)

	port := 8080
	addr := fmt.Sprintf(":%d", port)
	logger.Log("msg", "HTTP server is starting", "port", port)
	logger.Log("err", http.ListenAndServe(addr, handler))
}
