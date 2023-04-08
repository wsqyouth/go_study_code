package main

import (
	"fmt"
	"github.com/jasonkayzk/consistent-hashing-demo/core"
	"github.com/jasonkayzk/consistent-hashing-demo/proxy"
	"net/http"
)

var (
	port = "18888"

	p = proxy.NewProxy(core.NewConsistent(10, nil))
)

func main() {
	stopChan := make(chan interface{})
	startServer(port)
	<-stopChan
}

/*
启动代理服务器,监听端口18888
*/
func startServer(port string) {
	http.HandleFunc("/register", registerHost)
	http.HandleFunc("/unregister", unregisterHost)
	http.HandleFunc("/key", getKey)
	http.HandleFunc("/key_least", getKeyLeast)

	fmt.Printf("start proxy server: %s\n", port)

	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		panic(err)
	}
}

// 注册缓存服务器
func registerHost(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()

	err := p.RegisterHost(r.Form["host"][0])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprintf(w, err.Error())
		return
	}

	_, _ = fmt.Fprintf(w, fmt.Sprintf("register host: %s success", r.Form["host"][0]))
}

// 注销缓存服务器
func unregisterHost(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()

	err := p.UnregisterHost(r.Form["host"][0])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprintf(w, err.Error())
		return
	}

	_, _ = fmt.Fprintf(w, fmt.Sprintf("unregister host: %s success", r.Form["host"][0]))
}

// 查询缓存key
func getKey(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()

	val, err := p.GetKey(r.Form["key"][0])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprintf(w, err.Error())
		return
	}

	_, _ = fmt.Fprintf(w, fmt.Sprintf("key: %s, val: %s", r.Form["key"][0], val))
}

func getKeyLeast(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()

	val, err := p.GetKeyLeast(r.Form["key"][0])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = fmt.Fprintf(w, err.Error())
		return
	}

	_, _ = fmt.Fprintf(w, fmt.Sprintf("key: %s, val: %s", r.Form["key"][0], val))
}
