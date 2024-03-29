package service

import (
	"context"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/Brainsoft-Raxat/aiesec-hack/internal/models"
	"github.com/Brainsoft-Raxat/aiesec-hack/internal/repository"
	"github.com/Brainsoft-Raxat/aiesec-hack/pkg/apperror"
	"github.com/Brainsoft-Raxat/aiesec-hack/pkg/data"
	"github.com/Brainsoft-Raxat/aiesec-hack/pkg/errcodes"
	"github.com/google/uuid"
)

type eventService struct {
	repo *repository.Repository
}

// FetchAndCache implements EventService.
func (s *eventService) FetchAndCache(ctx context.Context) error {
	jerries, _ := s.repo.JerryStore.GetAllJerries(ctx)

	for _, jerry := range jerries {
		eventsGroup, err := s.GetEventsFiltered(ctx, data.GetEventsFilteredRequest{
			JerryID:    jerry.ID,
			City:       jerry.City,
			Categories: []string{},
		})
		if err != nil {
			return err
		}

		err = s.repo.Redis.CacheEvents(ctx, jerry.ID, eventsGroup.Events)
		if err != nil {
			return err
		}
	}

	return nil
}

// DeleteEvent implements EventService.
func (s *eventService) DeleteEvent(ctx context.Context, request data.DeleteEventRequest) (data.DeleteEventResponse, error) {
	panic("unimplemented")
}

// GetEvent implements EventService.festivali
func (s *eventService) GetEvent(ctx context.Context, request data.GetEventRequest) (data.GetEventResponse, error) {
	panic("unimplemented")
}

func (s *eventService) GiveSuggestion(ctx context.Context, request data.GiveSuggestionRequest) (resp data.GiveSuggestionResponse, err error) {
	events, err := s.repo.GetEventsForJerryID(ctx, request.JerryID)
	if err != nil {
		return
	}

	resp.Message, err = s.repo.GPT.SendPrompt(ctx, events)
	if err != nil {
		return
	}

	return
}

func (s *eventService) GetEventsFiltered(ctx context.Context, request data.GetEventsFilteredRequest) (resp data.GetEventsFilteredResponse, err error) {
	jerry, err := s.repo.JerryStore.GetJerryByID(ctx, request.JerryID)
	if err != nil {
		return
	}

	events, err := s.repo.Postgres.GetEventsFiltered(ctx, jerry.City, request.Categories)
	if err != nil {
		return resp, apperror.NewErrorInfo(ctx, errcodes.InternalServerError, err.Error())
	}

	// TODO: Do filtering by jerries location

	for i, event := range events {
		coords := strings.Split(event.Location, " ")

		dst := models.Coordinates{}

		dst.Latitude, err = strconv.ParseFloat(coords[0], 64)
		if err != nil {
			return resp, apperror.NewErrorInfo(ctx, errcodes.InternalServerError, "Unable to parse dst.Latitude")
		}

		dst.Longitude, err = strconv.ParseFloat(coords[1], 64)
		if err != nil {
			return resp, apperror.NewErrorInfo(ctx, errcodes.InternalServerError, "Unable to parse dst.Longitude")
		}

		events[i].Distance = Haversine(models.Coordinates{
			Latitude:  jerry.Latitude,
			Longitude: jerry.Longitude,
		},
			dst,
		)

		events[i].DistanceKM = fmt.Sprintf("%.1f km", events[i].Distance)

		events[i].Latitude = dst.Latitude
		events[i].Longitude = dst.Longitude
	}

	// Group events by date
	dateGroups := make(map[time.Time][]models.Event)
	for _, event := range events {
		date := event.Datetime.Truncate(24 * time.Hour)
		dateGroups[date] = append(dateGroups[date], event)
	}

	// Sort events within each date group
	for date, group := range dateGroups {
		sort.Slice(group, func(i, j int) bool {
			scoreI := score(group[i])
			scoreJ := score(group[j])

			// Sort in descending order, higher score first.
			return scoreI > scoreJ
		})
		dateGroups[date] = group
	}

	// Flatten the sorted date groups back into a single events slice
	var sortedEvents []models.Event
	for _, group := range dateGroups {
		sortedEvents = append(sortedEvents, group...)
	}

	resp.Events = sortedEvents

	return
}

// UpdateEvent implements EventService.
func (s *eventService) UpdateEvent(ctx context.Context, request data.UpdateEventRequest) (data.UpdateEventResponse, error) {
	panic("unimplemented")
}

// UpdateEvent implements EventService.
func (s *eventService) UpdateEventCount(ctx context.Context, request data.UpdateEventCountRequest) (resp data.UpdateEventCountResponse, err error) {
	err = s.repo.Postgres.UpdateEventCount(ctx, request.EventID)
	if err != nil {
		return resp, apperror.NewErrorInfo(ctx, errcodes.InternalServerError, err.Error())
	}

	resp.Success = true
	resp.Message = "ok"

	return
}

// CreateEvent implements EventService.
func (s *eventService) CreateEvent(ctx context.Context, request data.CreateEventRequest) (resp data.CreateEventResponse, err error) {
	resp.EventID, err = s.repo.Postgres.CreateEvent(ctx, models.Event{
		ID:          uuid.New(),
		Title:       request.Title,
		Description: request.Description,
		BannerURL:   request.BannerURL,
		Category:    request.Category,
		Author:      request.Author,
		Datetime:    request.Datetime,
		Address:     request.Address,
		Location:    request.Location,
		City:        request.City,
	})
	if err != nil {
		return
	}

	return
}

func Haversine(src, dst models.Coordinates) float64 {
	const EarthRadius = 6371.0

	lat1 := degToRad(src.Latitude)
	lon1 := degToRad(src.Longitude)
	lat2 := degToRad(dst.Latitude)
	lon2 := degToRad(dst.Longitude)

	// Differences in coordinates
	dLat := lat2 - lat1
	dLon := lon2 - lon1

	// Haversine formula
	a := math.Pow(math.Sin(dLat/2), 2) + math.Cos(lat1)*math.Cos(lat2)*math.Pow(math.Sin(dLon/2), 2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	distance := EarthRadius * c

	return distance
}

// Convert degrees to radians
func degToRad(deg float64) float64 {
	return deg * (math.Pi / 180)
}

func NewEventService(repo *repository.Repository) EventService {
	return &eventService{
		repo: repo,
	}
}

func score(event models.Event) float64 {
	// You can adjust these weight values as needed to prioritize distance and datetime.
	distanceWeight := 0.6
	datetimeWeight := 0.4

	// Normalize distance to a value between 0 and 1 (assuming non-negative distances).
	normalizedDistance := 1.0 / (event.Distance + 1.0)

	// Normalize datetime to a value between 0 and 1 (assuming Unix timestamps).
	normalizedDatetime := float64(event.Datetime.Unix()) / float64(time.Now().Unix())

	// Check for potential errors or edge cases.
	if normalizedDistance <= 0.0 {
		normalizedDistance = 1.0 // Set a default value or handle the error.
	}

	// Calculate the score based on the normalized and weighted values.
	return (normalizedDistance * distanceWeight) + (normalizedDatetime * datetimeWeight)
}
