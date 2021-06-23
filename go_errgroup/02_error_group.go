package main

import (
	"fmt"
	"sync"

	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/singleflight"
)

//errgroup包在sync.WaitGroup功能的基础上，增加了错误传递，以及在发生不可恢复的错误时取消整个goroutine集合，或者等待超时
func main() {

	citySz := "sz"
	citySh := "sh"
	info, err := Cities(citySz, citySh)
	if err != nil {
		fmt.Println("Get errors: ", err)
		return
	}
	for _, infoCity := range info {
		fmt.Println(*infoCity)
	}
}

type Info struct {
	cityID   int
	cityName string
}

var group singleflight.Group

func fetchWeatherFromDB(city string) (info *Info, err error) {
	info = new(Info)
	info.cityID = 510810
	info.cityName = city
	return info, nil
}
func City(city string) (*Info, error) {
	results, err, _ := group.Do(city, func() (interface{}, error) {
		info, err := fetchWeatherFromDB(city) // mock slow operation
		return info, err
	})
	if err != nil {
		return nil, fmt.Errorf("weather.City %s: %w", city, err)
	}
	return results.(*Info), nil
}

func Cities(cities ...string) ([]*Info, error) {
	var g errgroup.Group
	var mu sync.Mutex
	res := make([]*Info, len(cities)) // res[i] corresponds to cities[i]

	for i, city := range cities {
		i, city := i, city // create locals for closure below
		g.Go(func() error {
			info, err := City(city)
			mu.Lock()
			res[i] = info
			mu.Unlock()
			return err
		})
	}
	if err := g.Wait(); err != nil {
		return nil, err
	}
	return res, nil
}

// ref: https://encore.dev/blog/advanced-go-concurrency
