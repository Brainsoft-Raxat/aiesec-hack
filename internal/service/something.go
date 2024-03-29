package service

import (
	"context"
	"time"

	"github.com/Brainsoft-Raxat/aiesec-hack/internal/repository"
	"github.com/Brainsoft-Raxat/aiesec-hack/pkg/data"
)

type someService struct {
	someRepo *repository.Repository
}

func (s *someService) DoSomething(ctx context.Context, req data.DoSomethingRequest) (resp data.DoSomethingResponse, err error) {
	ctx, cancel := context.WithTimeout(ctx, 20*time.Second)
	defer cancel()

	return
}

func NewSomeService(repo *repository.Repository) SomeService {
	return &someService{
		someRepo: repo,
	}
}
