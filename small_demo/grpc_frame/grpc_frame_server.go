package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type GetShortURLReq struct {
	Id        string `json:"id"`
	IsEncrypt bool   `json:"is_encrypt"`
}

type GetShortURLResp struct {
	ShortURL string `json:"short_url"`
	ExpireTs int64  `json:"expire_ts"`
}

type ShortenProxyService interface {
	GetShortURL(ctx context.Context, req *GetShortURLReq) (*GetShortURLResp, error)
}

type Server struct {
	service ShortenProxyService
}

func NewServer(service ShortenProxyService) *Server {
	return &Server{service: service}
}

func (s *Server) GetShortURLHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		req := &GetShortURLReq{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		resp, err := s.service.GetShortURL(r.Context(), req)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	})
}

type MyService struct{}

func (s *MyService) GetShortURL(ctx context.Context, req *GetShortURLReq) (*GetShortURLResp, error) {
	// 这里只是一个示例，实际的实现可能需要查询数据库或者其他的服务
	shortURL := "http://example.com/" + req.Id
	return &GetShortURLResp{ShortURL: shortURL, ExpireTs: time.Now().Unix()}, nil
}

func main() {
	service := &MyService{}
	server := NewServer(service)

	mux := http.NewServeMux()
	mux.Handle("/GetShortURL", server.GetShortURLHandler())
	fmt.Println("sever running...")
	http.ListenAndServe(":8080", mux)
}

/*
curl -X POST -H "Content-Type: application/json" -d '{"id":"123", "is_encrypt":false}' http://localhost:8080/GetShortURL

一个server内包含了这个service，在回调函数内同样使用代理模式实现对注入的业务逻辑进行调用
这里抽离这个demo是像研究怎么实现一个server框架的，为后面阅读gin/grpc源码框架做铺垫
*/
