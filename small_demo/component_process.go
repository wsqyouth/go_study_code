package main

import (
	"context"
	"errors"
	"fmt"
)

// YourStruct 是一个示例结构体，用于请求参数
type YourStruct struct {
	Value string
}

// Config 是用于校验的配置结构体
type Config struct {
	RequiredValue string
}

// ElementProcessor 接口定义了校验和转换两个方法
type ElementProcessor interface {
	Validate(ctx context.Context, params map[string]YourStruct, config Config) error
	Transform(ctx context.Context, params map[string]YourStruct) (map[string]YourStruct, error)
}

// MockComponent 实现了 ElementProcessor 接口
type MockComponent struct {
	name string
}

// Validate 方法模拟校验过程
func (m *MockComponent) Validate(ctx context.Context, params map[string]YourStruct, config Config) error {
	if m.name == "" {
		return errors.New("component name is empty")
	}
	if val, ok := params[m.name]; ok {
		if val.Value != config.RequiredValue {
			return fmt.Errorf("validation failed for %s: expected %s, got %s", m.name, config.RequiredValue, val.Value)
		}
	} else {
		return fmt.Errorf("parameter for %s not found", m.name)
	}
	fmt.Printf("Validating component: %s\n", m.name)
	return nil
}

// Transform 方法模拟转换过程
func (m *MockComponent) Transform(ctx context.Context, params map[string]YourStruct) (map[string]YourStruct, error) {
	fmt.Printf("Transforming component: %s\n", m.name)
	if val, ok := params[m.name]; ok {
		val.Value = "Transformed " + val.Value
		params[m.name] = val
	} else {
		return nil, fmt.Errorf("parameter for %s not found", m.name)
	}
	return params, nil
}

// NewElementProcessor 是一个简单工厂方法，用于创建 ElementProcessor 实例
func NewElementProcessor(name string) ElementProcessor {
	return &MockComponent{name: name}
}

func main() {
	ctx := context.Background()
	params := map[string]YourStruct{
		"ComponentA": {Value: "ValueA"},
		"ComponentB": {Value: "ValueA"},
	}
	config := Config{RequiredValue: "ValueA"}

	// 获取所有元素处理器
	elementProcessors := []ElementProcessor{
		NewElementProcessor("ComponentA"),
		NewElementProcessor("ComponentB"),
	}

	// 校验所有元素处理器
	for _, processor := range elementProcessors {
		err := processor.Validate(ctx, params, config)
		if err != nil {
			fmt.Printf("Validation error: %s\n", err)
			return
		}
	}

	// 转换所有元素处理器
	for _, processor := range elementProcessors {
		transformedParams, err := processor.Transform(ctx, params)
		if err != nil {
			fmt.Printf("Transformation error: %s\n", err)
			return
		}
		fmt.Printf("Transformed parameters: %+v\n", transformedParams)
	}
}
