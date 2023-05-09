package main

import (
	"context"
	"fmt"
	"github.com/agiledragon/gomonkey"
	"net/http"
	"reflect"
	"testing"
	"time"
)

func Test_testRequest(t *testing.T) {
	// Create a context with a very short timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond)
	defer cancel()

	// Create the request with it
	r, _ := http.NewRequest("GET", "http://example.com", nil)
	r = r.WithContext(ctx)

	// Create a fake http client
	fakeClient := &http.Client{}

	// Use gomonkey to replace the Do method of the fake client with a custom implementation
	patches := gomonkey.ApplyMethod(reflect.TypeOf(fakeClient), "Do", func(_ *http.Client, _ *http.Request) (*http.Response, error) {
		return nil, fmt.Errorf("Get http://example.com: %w", context.DeadlineExceeded)
	})
	defer patches.Reset()

	// Call the testRequest function with the fake client
	result := testRequest(fakeClient)

	// Check if the result is true, which means the error is context.DeadlineExceeded
	if !result {
		t.Errorf("Expected context.DeadlineExceeded, but got a different error")
	}
}
