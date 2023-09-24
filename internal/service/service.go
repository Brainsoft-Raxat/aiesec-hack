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
	UpdateEventCount(ctx context.Context, request data.UpdateEventCountRequest) (resp data.UpdateEventCountResponse, err error)

	// Worker methods
	FetchAndCache(ctx context.Context) error
	GiveSuggestion(ctx context.Context, request data.GiveSuggestionRequest) (resp data.GiveSuggestionResponse, err error)
}

type ShootService interface {
	SendShoot(ctx context.Context, request data.SendShootRequest) (data.SendShootResponse, error)
}

type PromotionService interface {
	GetPromotionsFiltered(ctx context.Context, request data.GetPromotionsFilteredRequest) (resp data.GetPromotionsFilteredResponse, err error)
	CreatePromotion(ctx context.Context, request data.CreatePromotionRequest) (resp data.CreatePromotionResponse, err error)
}

type Service struct {
	SomeService
	AuthService
	EventService
	ShootService
	PromotionService
}

func New(repos *repository.Repository) *Service {
	return &Service{
		SomeService:      NewSomeService(repos),
		AuthService:      NewAuthService(repos),
		EventService:     NewEventService(repos),
		ShootService:     NewShootService(repos),
		PromotionService: NewPromotionService(repos),
	}
}
