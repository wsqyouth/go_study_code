package main

import (
	"context"
	"fmt"
)

type GetShortURLReq struct {
	Id        string `json:"id"`
	IsEncrypt bool   `json:"is_encrypt"`
}

type GetShortURLResp struct {
	ShortURL string `json:"short_url"`
	ExpireTs int64  `json:"expire_ts"`
}

type ShortenProxyClient interface {
	GetShortURL(ctx context.Context, req *GetShortURLReq) (*GetShortURLResp, error)
}

type RealShortenProxyClient struct{}

func (c *RealShortenProxyClient) GetShortURL(ctx context.Context, req *GetShortURLReq) (*GetShortURLResp, error) {
	// 这里只是一个示例，实际的实现可能需要查询数据库或者其他的服务
	shortURL := "http://example.com/" + req.Id
	return &GetShortURLResp{ShortURL: shortURL, ExpireTs: 1234567890}, nil
}

type ShortenProxyClientProxy struct {
	client ShortenProxyClient
}

func (p *ShortenProxyClientProxy) GetShortURL(ctx context.Context, req *GetShortURLReq) (*GetShortURLResp, error) {
	fmt.Println("Proxy: Before calling the real client")
	resp, err := p.client.GetShortURL(ctx, req)
	fmt.Println("Proxy: After calling the real client")
	return resp, err
}

func main() {
	// 创建实际的客户端对象
	client := &RealShortenProxyClient{}

	// 创建代理客户端对象
	proxy := &ShortenProxyClientProxy{
		client: client,
	}

	// 创建请求对象
	req := &GetShortURLReq{
		Id:        "123",
		IsEncrypt: false,
	}

	// 调用代理方法
	rsp, err := proxy.GetShortURL(context.Background(), req)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Response: %+v\n", rsp)
	}
}

/*
总结：
这个其实就是代理模式，这个代理封装下，实现对实现类功能的调用
*/
