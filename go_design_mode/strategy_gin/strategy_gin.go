package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// AuthStrategy defines the set of methods used to do resource authentication.
type AuthStrategy interface {
	AuthFunc(w http.ResponseWriter, r *http.Request)
}

// BasicStrategy implements the basic authentication strategy.
type BasicStrategy struct{}

func (b BasicStrategy) AuthFunc(w http.ResponseWriter, r *http.Request) {
	// Simulate Basic authentication logic
	log.Println("Using Basic Authentication Strategy")
	// Here you would parse and verify the Basic token
	w.Write([]byte("Basic Authentication Successful\n"))
}

// JWTStrategy implements the JWT (Bearer) authentication strategy.
type JWTStrategy struct{}

func (j JWTStrategy) AuthFunc(w http.ResponseWriter, r *http.Request) {
	// Simulate JWT authentication logic
	log.Println("Using JWT (Bearer) Authentication Strategy")
	// Here you would parse and verify the JWT token
	w.Write([]byte("JWT Authentication Successful\n"))
}

// AutoStrategy defines authentication strategy which can automatically choose between Basic and Bearer
type AutoStrategy struct {
	basic AuthStrategy
	jwt   AuthStrategy
}

func NewAutoStrategy(basic, jwt AuthStrategy) AutoStrategy {
	return AutoStrategy{
		basic: basic,
		jwt:   jwt,
	}
}

func (a AutoStrategy) AuthFunc(w http.ResponseWriter, r *http.Request) {
	authHeader := strings.SplitN(r.Header.Get("Authorization"), " ", 2)

	if len(authHeader) != 2 {
		log.Println("Authorization header format is wrong.")
		http.Error(w, "Authorization header format is wrong.", http.StatusBadRequest)
		return
	}

	switch authHeader[0] {
	case "Basic":
		a.basic.AuthFunc(w, r)
	case "Bearer":
		a.jwt.AuthFunc(w, r)
	default:
		log.Println("Unrecognized Authorization header.")
		http.Error(w, "Unrecognized Authorization header.", http.StatusUnauthorized)
		return
	}
}

func main() {
	basicStrategy := BasicStrategy{}
	jwtStrategy := JWTStrategy{}
	autoStrategy := NewAutoStrategy(basicStrategy, jwtStrategy)

	http.HandleFunc("/auth", func(w http.ResponseWriter, r *http.Request) {
		autoStrategy.AuthFunc(w, r)
	})

	fmt.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
