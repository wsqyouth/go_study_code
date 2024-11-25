package main

import (
	"errors"
	"fmt"
	"reflect"
)

// ErrDetail represents the detailed information of an error
type ErrDetail struct {
	Info string
}

// PostmenErr represents a custom error type
type PostmenErr struct {
	Code     int
	Message  string
	Details  []ErrDetail
	originErr error
}

// Error implements the error interface for PostmenErr
func (e *PostmenErr) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s, Details: %v", e.Code, e.Message, e.Details)
}

// Equal compares two PostmenErr objects for equality
func (e *PostmenErr) Equal(other *PostmenErr) bool {
	if e == nil || other == nil {
		return e == other
	}
	return e.Code == other.Code &&
		e.Message == other.Message &&
		reflect.DeepEqual(e.Details, other.Details)
}

func main() {
	// Create two PostmenErr objects
	err1 := &PostmenErr{
		Code:    404,
		Message: "Not Found",
		Details: []ErrDetail{{Info: "Resource not found"}},
	}

	err2 := &PostmenErr{
		Code:    404,
		Message: "Not Found",
		Details: []ErrDetail{{Info: "Resource not found"}},
	}

	// Method 1: Using reflect.DeepEqual
	if reflect.DeepEqual(err1, err2) {
		fmt.Println("Errors are equal using reflect.DeepEqual")
	} else {
		fmt.Println("Errors are not equal using reflect.DeepEqual")
	}

	// Method 2: Using custom Equal method
	if err1.Equal(err2) {
		fmt.Println("Errors are equal using custom Equal method")
	} else {
		fmt.Println("Errors are not equal using custom Equal method")
	}

	// Demonstrating errors.As
	var target *PostmenErr
	if errors.As(err1, &target) {
		fmt.Println("Error is of type *PostmenErr using errors.As")
	}

	// Demonstrating errors.Is
	if errors.Is(err1, err2) {
		fmt.Println("Errors are the same using errors.Is")
	} else {
		fmt.Println("Errors are not the same using errors.Is")
	}
}


/*
学习如何使用 errors.Is 和 errors.As，以及如何通过 reflect.DeepEqual 和自定义 Equal 方法比较自定义错误对象
*/
