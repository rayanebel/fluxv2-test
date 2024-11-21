package main

import (
	"fmt"
	"net/http"
)

func Hello() {
	println("Hello from the package")
}

func main() {
	mu := http.NewServeMux()
	mu.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello world"))
	})
	panic("Something went wrong")

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", mu)
}
