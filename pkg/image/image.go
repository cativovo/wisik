package image

import "github.com/google/uuid"

type Image struct {
	Label string
	Src   string
	Id    uuid.UUID
}
