package service

import (
	"context"

	"github.com/Brainsoft-Raxat/aiesec-hack/internal/repository"
	"github.com/Brainsoft-Raxat/aiesec-hack/pkg/data"
)

type SomeService interface {
	DoSomething(ctx context.Context, request data.DoSomethingRequest) (response data.DoSomethingResponse, err error)
}

type AuthService interface {
	Login(ctx context.Context, creds data.UserSignInRequest) (data.Tokens, error)
}

type EventService interface {
	CreateEvent(ctx context.Context, request data.CreateEventRequest) (data.CreateEventResponse, error)
	UpdateEvent(ctx context.Context, request data.UpdateEventRequest) (data.UpdateEventResponse, error)
	GetEvent(ctx context.Context, request data.GetEventRequest) (data.GetEventResponse, error)
	GetEventsFiltered(ctx context.Context, request data.GetEventsFilteredRequest) (data.GetEventsFilteredResponse, error)
	DeleteEvent(ctx context.Context, request data.DeleteEventRequest) (data.DeleteEventResponse, error)
}
type Service struct {
	SomeService
	AuthService
	EventService
}

func New(repos *repository.Repository) *Service {
	return &Service{
		SomeService:  NewSomeService(repos),
		AuthService:  NewAuthService(repos),
		EventService: NewEventService(repos),
	}
}
