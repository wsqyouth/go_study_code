// ddd_app/main.go
package main

import (
	"ddd_app/internal/api"
	"ddd_app/internal/application"
	"ddd_app/internal/infrastructure/persistence"
	"encoding/json"
	"fmt"
	"net/http"

	// 引入providers
	_ "ddd_app/internal/providers"
)

func main() {
	fmt.Println("app running")
	articleRepo := persistence.NewArticleMapRepo()
	UserRepo := persistence.NewUserInfoRepo()
	blogService := application.NewBlogService(articleRepo, UserRepo)
	blogHandler := api.NewBlogHandler(blogService)

	http.HandleFunc("/miniblog", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			fmt.Println("app post")
			status, err := blogHandler.WriteArticle(r)
			if err != nil {
				http.Error(w, err.Error(), status)
				return
			}
			w.WriteHeader(status)
			w.Write([]byte("Article created successfully"))
		case http.MethodGet:
			fmt.Println("app get")
			articles, status, err := blogHandler.GetArticlesByUserID(r)
			if err != nil {
				http.Error(w, err.Error(), status)
				return
			}
			json.NewEncoder(w).Encode(articles)
		default:
			http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		}
	})
	http.ListenAndServe(":8080", nil)
}
