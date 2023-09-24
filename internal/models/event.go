package models

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	BannerURL   string    `json:"banner_url"`
	Category    string    `json:"category"`
	Author      string    `json:"author"`
	Datetime    time.Time `json:"datetime"`
	Address     string    `json:"address"`
	Location    string    `json:"location"`
	Latitude    float64   `json:"latitude"`
	Longitude   float64   `json:"longitude"`
	City        string    `json:"city"`

	DistanceKM string  `json:"distance"`
	Distance   float64 `json:"-"`

	Count int `json:"count"`
}
