// ddd_app/internal/api/article_handler.go
package api

import (
	"ddd_app/internal/application"
	"ddd_app/internal/domain/entity"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type BlogHandler struct {
	service *application.BlogService
}

func NewBlogHandler(s *application.BlogService) *BlogHandler {
	return &BlogHandler{service: s}
}

func (h *BlogHandler) WriteArticle(r *http.Request) (int, error) {
	var article entity.Article
	err := json.NewDecoder(r.Body).Decode(&article)
	if err != nil {
		return http.StatusBadRequest, err
	}
	fmt.Println("req article: ", article)
	err = h.service.WriteArticle(r.Context(), article)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.StatusCreated, nil
}

func (h *BlogHandler) GetArticlesByUserID(r *http.Request) ([]entity.Article, int, error) {
	userID, err := strconv.Atoi(r.URL.Query().Get("userID"))
	if err != nil {
		return nil, http.StatusBadRequest, err
	}
	fmt.Println("req userID: ", userID)
	articles, err := h.service.GetArticlesByUserID(r.Context(), userID)
	if err != nil {
		return nil, http.StatusInternalServerError, err
	}
	fmt.Println("rsp articles: ", articles)
	return articles, http.StatusOK, nil
}
