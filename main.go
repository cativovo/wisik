package main

import (
	"github.com/cativovo/wisik/pkg/http"
	"github.com/cativovo/wisik/pkg/image"
	"github.com/cativovo/wisik/pkg/storage/postgres"
)

func main() {
	postgresRepository := postgres.NewPostgresRepository()
	imageService := image.NewService(postgresRepository)

	server := http.NewServer(imageService)
	server.ListenAndServe("127.0.0.1:4000")
}
