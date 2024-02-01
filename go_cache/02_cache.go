package main

import (
	"errors"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/bluele/gcache"
)

var studentCache gcache.Cache // Global cache variable

// Initialize a new ARC Cache with size of 10 items and expiry after 1 minute
func initStudentScoreCache() {
	studentCache = gcache.New(10).
		ARC().
		Expiration(time.Minute).
		LoaderExpireFunc(loadAndCalculateScoresFromDatabase).
		Build()
}

type Student struct {
	Name   string
	Age    int
	Grade  float64
	Gender bool
}

// This function simulates loading data from database or any external source
func loadAndCalculateScoresFromDatabase(key interface{}) (value interface{}, duration *time.Duration, e error) {
	var score = 89.7 // Assume we get this value from DB for given student name
	timer := 10 * time.Hour
	duration = &timer
	name, ok := key.(string)
	if !ok {
		errMsg := "Key type not supported."
		return nil, nil, errors.New(errMsg)
	} else if len(name) == 0 {
		errMsg := "Empty Key received."
		return nil, nil, errors.New(errMsg)
	}
	s := Student{
		Name:   name,
		Age:    18,
		Grade:  score,
		Gender: true,
	}
	return s, duration, nil
}

func main() {
	initStudentScoreCache()
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			scoreInterface, err := studentCache.Get("John Doe-" + strconv.Itoa(i))
			if err != nil || scoreInterface == nil {
				fmt.Println("Failed to fetch score from cache.")
			}
			switch v := scoreInterface.(type) {
			case Student:
				_ = v
			default:
				fmt.Println("Type mismatch! Expected 'Student' but got ", v)
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("All goroutines finished.")
}
