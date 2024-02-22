// ddd_app/internal/api/article_handler.go
package api

import (
	"fmt"
	"encoding/json"
	"net/http"
	"strconv"
	"ddd_app/internal/application"
	"ddd_app/internal/domain/entity"
)

type ArticleHandler struct {
	service *application.ArticleService
}

func NewArticleHandler(s *application.ArticleService) *ArticleHandler {
	return &ArticleHandler{service: s}
}

func (h *ArticleHandler) WriteArticle(r *http.Request) (int, error) {
	var article entity.Article
	err := json.NewDecoder(r.Body).Decode(&article)
	if err != nil {
		return http.StatusBadRequest, err
	}
	fmt.Println("req article: ",article)
	err = h.service.WriteArticle(r.Context(), article)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusCreated, nil
}

func (h *ArticleHandler) GetArticlesByUserID(r *http.Request) ([]entity.Article, int, error) {
	userID, err := strconv.Atoi(r.URL.Query().Get("userID"))
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	fmt.Println("req userID: ",userID)
	articles, err := h.service.GetArticlesByUserID(r.Context(), userID)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	fmt.Println("rsp articles: ",articles)
	return articles, http.StatusOK, nil
}