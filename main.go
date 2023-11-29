package main

import "github.com/cativovo/wisik/pkg/http"

func main() {
	server := http.NewServer()
	server.ListenAndServe("127.0.0.1:4000")
}
