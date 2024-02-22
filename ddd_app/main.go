// ddd_app/main.go
package main

import (
	"fmt"
	"encoding/json"
	"net/http"
	"ddd_app/internal/api"
	"ddd_app/internal/application"
	"ddd_app/internal/infrastructure/persistence"
)

func main() {
	fmt.Println("app running")
	articleRepo := persistence.NewArticleMapRepo()
	articleService := application.NewArticleService(articleRepo)
	articleHandler := api.NewArticleHandler(articleService)

	http.HandleFunc("/articles", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			fmt.Println("app post")
			status, err := articleHandler.WriteArticle(r)
			if err != nil {
				http.Error(w, err.Error(), status)
				return
			}
			w.WriteHeader(status)
			w.Write([]byte("Article created successfully"))
		case http.MethodGet:
			fmt.Println("app get")
			articles, status, err := articleHandler.GetArticlesByUserID(r)
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