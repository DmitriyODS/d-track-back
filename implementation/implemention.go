package implementation

import "gitlab.com/ddda/d-track/d-track-back/repository"

type BasicService struct {
	rep repository.Repository
}

func NewBasicService(rep repository.Repository) *BasicService {
	return &BasicService{rep: rep}
}
