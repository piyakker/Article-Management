package publishing

import (
	"time"

	"github.com/go-kit/log"
)

func NewLoggingMiddleware(logger log.Logger, next ArticleService) logmw {
	return logmw{logger, next}
}

type logmw struct {
	logger         log.Logger
	articleService ArticleService
}

func (mw logmw) GetAllArticles() (articles []*Article, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "GetAllArticles",
			"input", "",
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	articles, err = mw.articleService.GetAllArticles()
	return
}

func (mw logmw) CreateArticle(title, content, author string) (str string, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "CreateArticle",
			"input", title+", "+content+", "+author,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	str, err = mw.articleService.CreateArticle(title, content, author)
	return
}

func (mw logmw) GetArticleByID(id string) (article *Article, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "GetArticleByID",
			"input", id,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	article, err = mw.articleService.GetArticleByID(id)
	return
}

func (mw logmw) UpdateArticle(id, title, content, author string) (str string, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "UpdateArticle",
			"input", id+", "+title+", "+content+", "+author,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	str, err = mw.articleService.UpdateArticle(id, title, content, author)
	return
}

func (mw logmw) DeleteArticle(id string) (str string, err error) {
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "DeleteArticle",
			"input", id,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	str, err = mw.articleService.DeleteArticle(id)
	return
}
