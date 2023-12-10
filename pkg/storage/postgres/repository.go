package postgres

import (
	"context"
	"log"

	"github.com/cativovo/wisik/pkg/image"
	postgres "github.com/cativovo/wisik/pkg/storage/postgres/sqlc_generated"
	"github.com/jackc/pgx/v5"
)

type PostgresRepository struct {
	ctx     context.Context
	queries *postgres.Queries
}

func NewPostgresRepository() *PostgresRepository {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "postgres://postgres:1234@127.0.0.1/wisik?sslmode=disable")
	if err != nil {
		log.Fatal("Can't connect to the db")
	}

	queries := postgres.New(conn)

	return &PostgresRepository{
		ctx:     ctx,
		queries: queries,
	}
}

func (p *PostgresRepository) ListImages() ([]image.Image, error) {
	result, err := p.queries.ListImages(p.ctx)
	if err != nil {
		return nil, err
	}

	images := make([]image.Image, 0, len(result))

	for _, v := range result {
		images = append(images, image.Image{
			Id:    v.ID.Bytes,
			Label: v.Label,
			Src:   v.Src,
		})
	}

	return images, nil
}

func (p *PostgresRepository) AddImage(label, src string) (image.Image, error) {
	params := postgres.AddImageParams{
		Label: label,
		Src:   src,
	}

	result, err := p.queries.AddImage(p.ctx, params)
	if err != nil {
		return image.Image{}, err
	}

	return image.Image{
		Id:    result.ID.Bytes,
		Label: result.Label,
		Src:   result.Src,
	}, nil
}

func (p *PostgresRepository) DeleteImage(imgId string) error {
	return nil
}
