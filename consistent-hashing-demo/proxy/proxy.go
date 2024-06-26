package proxy

import (
	"fmt"
	"github.com/jasonkayzk/consistent-hashing-demo/core"
	"io/ioutil"
	"net/http"
	"time"
)

type Proxy struct {
	consistent *core.Consistent
}

/*
代理服务器选择具体使用哪个缓存服务器处理请求
*/
// NewProxy creates a new Proxy
func NewProxy(consistent *core.Consistent) *Proxy {
	proxy := &Proxy{
		consistent: consistent,
	}
	return proxy
}

/*
v0版本策略:根据key计算并采用一定策略得到特定服务器host,使用http请求向其发出业务请求获取内容
*/
func (p *Proxy) GetKey(key string) (string, error) {

	host, err := p.consistent.GetKey(key)
	if err != nil {
		return "", err
	}

	resp, err := http.Get(fmt.Sprintf("http://%s?key=%s", host, key))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Printf("Response from host %s: %s\n", host, string(body))

	return string(body), nil
}

/*
v1版本策略:根据key计算并采用一定策略得到特定服务器host,使用http请求向其发出业务请求获取内容
*/
func (p *Proxy) GetKeyLeast(key string) (string, error) {

	host, err := p.consistent.GetKeyLeast(key)
	if err != nil {
		return "", err
	}
	p.consistent.Inc(host)

	time.AfterFunc(time.Second*10, func() { // drop the host after 10 seconds(for testing)!
		fmt.Printf("dropping host: %s after 10 second\n", host)
		p.consistent.Done(host)
	})

	resp, err := http.Get(fmt.Sprintf("http://%s?key=%s", host, key))
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	fmt.Printf("Response from host %s: %s\n", host, string(body))

	return string(body), nil
}

func (p *Proxy) RegisterHost(host string) error {

	err := p.consistent.RegisterHost(host)
	if err != nil {
		return err
	}

	fmt.Println(fmt.Sprintf("register host: %s success", host))
	return nil
}

func (p *Proxy) UnregisterHost(host string) error {
	err := p.consistent.UnregisterHost(host)
	if err != nil {
		return err
	}

	fmt.Println(fmt.Sprintf("unregister host: %s success", host))
	return nil
}
