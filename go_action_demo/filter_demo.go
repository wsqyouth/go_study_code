package main

import (
	"errors"
	"fmt"

	"gopkg.in/yaml.v3"
)

func main() {
	//1. 解析插件配置
	// 2. 插件注册
	typ := "app"
	name := "app"
	cfg := yaml.Node{
		Tag:   "app_key",
		Value: "app_val",
	}
	Register("app", AppPlugin)
	//3. 插件初始化
	factory := Get(typ, name)
	factory.Setup(name, &YamlNodeDecoder{Node: &cfg})
	//4. 插件内部方法调用
	fmt.Println("hello")
}

// Factory
type Factory interface {
	// Type 插件类型
	Type() string
	// Setup 装载插件，需要用户先定义好插件数据结构
	Setup(name string, dec Decoder) error
}

// Register 注册插件工厂
func Register(name string, f Factory) {
	factories, ok := plugins[f.Type()]
	if !ok {
		plugins[f.Type()] = map[string]Factory{
			name: f,
		}
		return
	}
	factories[name] = f
}

// Get 根据插件类型,插件名字获取插件工厂。
func Get(typ string, name string) Factory {
	factories, ok := plugins[typ]
	if !ok {
		return nil
	}
	return factories[name]
}

// Decoder
type Decoder interface {
	Decode(cfg interface{}) error
}

// YamlNodeDecoder yaml
type YamlNodeDecoder struct {
	Node *yaml.Node
}

// Decode 解析
func (d *YamlNodeDecoder) Decode(cfg interface{}) error {
	if d.Node == nil {
		return errors.New("yaml node empty")
	}
	return d.Node.Decode(cfg)
}

var (
	plugins = make(map[string]map[string]Factory)
)

// AppPlugin FFlag 插件
var AppPlugin = &appPlugin{}

type appPlugin struct {
	AppKey    string `yaml:"app_key"`
	AppSecret string `yaml:"app_secret"`
}

// Type 插件类型
func (p *appPlugin) Type() string {
	return "app"
}

// Setup 方法
func (p *appPlugin) Setup(name string, decoder Decoder) error {
	/*
		err := decoder.Decode(p)
		if err != nil {
			fmt.Println("decode conf error", err)
			return err
		}
	*/
	fmt.Println(p)
	fmt.Println("i am new plgin, has set up")
	return nil
}
