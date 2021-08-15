package main

import (
	"fmt"

	"github.com/golang/groupcache/lru"
)

func main() {
	cache := lru.New(2)
	cache.Add("x", "x0")
	cache.Add("y", "y0")
	yval, ok := cache.Get("y")
	if ok {
		fmt.Printf("y is %v\n", yval)

	}
	cache.Add("z", "z0")

	fmt.Printf("cache length is %d\n", cache.Len())
	_, ok = cache.Get("x")
	if !ok {
		fmt.Printf("x key was weeded out\n")

	}

}

//go run lru_example.go
//go fmt -s lru_example.go
