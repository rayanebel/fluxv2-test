package main

import (
	"fmt"
	"net/http"
)

func Add(a int, b int) int {
	return a + b
}
func main() {
	mu := http.NewServeMux()
	mu.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	})
	mu.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API Server"))
	})
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8000", mu)
}
