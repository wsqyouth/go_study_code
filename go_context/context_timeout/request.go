package main

import (
	"context"
	"errors"
	"net/http"
	"time"
)

func main() {
	testRequest(http.DefaultClient)
}

func testRequest(client *http.Client) bool {
	// Create a context with a very short timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
	defer cancel()

	// Create the request with it
	r, _ := http.NewRequest("GET", "http://example.com", nil)
	r = r.WithContext(ctx)

	// Do it, it will fail because the request will take longer than 1ms
	_, err := client.Do(r)

	// Check if the error is context.DeadlineExceeded
	return errors.Is(err, context.DeadlineExceeded)
}
