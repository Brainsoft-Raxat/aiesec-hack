package data

import (
	"time"

	"github.com/Brainsoft-Raxat/aiesec-hack/internal/models"
	"github.com/google/uuid"
)

type CreateEventRequest struct {
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Category    string    `json:"category"`
	Author      string    `json:"author"`
	Datetime    time.Time `json:"datetime"`
	Location    string    `json:"location"`
	City        string    `json:"city"`
}

type CreateEventResponse struct {
	EventID uuid.UUID `json:"event_id"`
}

type UpdateEventRequest struct {
	EventID     int    `json:"event_id,omitempty"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Author      string `json:"author"`
	Datetime    string `json:"datetime"`
	Location    string `json:"location"`
	City        string `json:"city"`
}

type UpdateEventResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type GetEventRequest struct {
	EventID int `json:"event_id"`
}

type GetEventResponse struct {
	Event models.Event `json:"event"`
}

type GetEventsFilteredRequest struct {
	JerryID    string   `json:"jerry_id"`
	City       string   `json:"city"`
	Categories []string `json:"categories"`
}

type GetEventsFilteredResponse struct {
	Events []models.Event `json:"events"`
}

type DeleteEventRequest struct {
	EventID int `json:"event_id"`
}
type DeleteEventResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
