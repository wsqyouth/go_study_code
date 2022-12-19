package main

import (
	"context"
	"fmt"
)

func main() {
	// getKeyDemo()
	params := map[string]interface{}{
		"id": "123",
	}
	ctx := setContextKey(context.Background(), params)
	getContextKey(ctx)
}

func getKeyDemo() {
	type contextKey string

	getValFunc := func(ctx context.Context, k contextKey) {
		if v := ctx.Value(k); v != nil {
			fmt.Println("found value:", v)
			return
		}
		fmt.Println("key not found:", k)
	}

	k := contextKey("language")
	ctx := context.WithValue(context.Background(), k, "Go")

	getValFunc(ctx, k)
	getValFunc(ctx, contextKey("color"))
}

var contextKey string = "inner_context_key"

func setContextKey(ctx context.Context, params map[string]interface{}) context.Context {
	return context.WithValue(ctx, contextKey, params)
}

func getContextKey(ctx context.Context) (params map[string]interface{}, ok bool) {
	params, ok = ctx.Value(contextKey).(map[string]interface{})
	fmt.Println("params:%v,ok:%v", params, ok)
	return
}
