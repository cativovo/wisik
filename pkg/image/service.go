package image

type Repository interface {
	ListImages() ([]Image, error)
	AddImage(label, src string) (Image, error)
	DeleteImage(id string) error
}

type Service struct {
	repository Repository
}

func NewService(r Repository) *Service {
	return &Service{
		repository: r,
	}
}

func (s *Service) ListImages() ([]Image, error) {
	return s.repository.ListImages()
}

func (s *Service) AddImage(label, src string) (Image, error) {
	return s.repository.AddImage(label, src)
}

func (s *Service) DeleteImage(id string) error {
	return s.repository.DeleteImage(id)
}
