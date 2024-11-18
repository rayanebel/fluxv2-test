package main

import "net/http"

func Hello() {
	println("Hello from the package")
}

func main() {

	mu := http.NewServeMux()
	mu.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from the main"))
	})
	http.ListenAndServe(":8080", mu)
}
