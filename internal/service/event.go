package service

import (
	"context"

	"github.com/Brainsoft-Raxat/aiesec-hack/internal/models"
	"github.com/Brainsoft-Raxat/aiesec-hack/internal/repository"
	"github.com/Brainsoft-Raxat/aiesec-hack/pkg/data"
	"github.com/google/uuid"
)

type eventService struct {
	repo *repository.Repository
}

// DeleteEvent implements EventService.
func (s *eventService) DeleteEvent(ctx context.Context, request data.DeleteEventRequest) (data.DeleteEventResponse, error) {
	panic("unimplemented")
}

// GetEvent implements EventService.
func (s *eventService) GetEvent(ctx context.Context, request data.GetEventRequest) (data.GetEventResponse, error) {
	panic("unimplemented")
}

// GetEventsFiltered implements EventService.
func (s *eventService) GetEventsFiltered(ctx context.Context, request data.GetEventsFilteredRequest) (resp data.GetEventsFilteredResponse, err error) {
	jerry, err := s.repo.JerryStore.GetJerryByID(ctx, request.JerryID)
	if err != nil {
		return
	}

	events, err := s.repo.Postgres.GetEventsFiltered(ctx, jerry.City, request.Categories)
	if err != nil {
		return
	}

	resp.Events = events

	// TODO: Do filtering by jerries location

	return
}

// UpdateEvent implements EventService.
func (s *eventService) UpdateEvent(ctx context.Context, request data.UpdateEventRequest) (data.UpdateEventResponse, error) {
	panic("unimplemented")
}

// CreateEvent implements EventService.
func (s *eventService) CreateEvent(ctx context.Context, request data.CreateEventRequest) (resp data.CreateEventResponse, err error) {
	resp.EventID, err = s.repo.Postgres.CreateEvent(ctx, models.Event{
		ID:          uuid.New(),
		Title:       request.Title,
		Description: request.Description,
		Category:    request.Category,
		Author:      request.Author,
		Datetime:    request.Datetime,
		Location:    request.Location,
		City:        request.City,
	})
	if err != nil {
		return
	}

	return
}

func NewEventService(repo *repository.Repository) EventService {
	return &eventService{
		repo: repo,
	}
}
