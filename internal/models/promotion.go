package models

import (
    "github.com/google/uuid"
    "time"
)

// Promotion represents a promotion entry.
type Promotion struct {
    ID            uuid.UUID       `json:"id"`
    Title         string          `json:"title"`
    BannerURL     string          `json:"banner_url"`
    ReviewsNumber int             `json:"reviews_number"`
    ReviewsRate   float64         `json:"reviews_rate"`
    Expires       time.Time       `json:"expires"`
    Discount      int             `json:"discount"`
    City          string          `json:"city"`
    Address       string          `json:"address"`
    Latitude      float64         `json:"latitude"`
    Longitude     float64         `json:"longitude"`
    Price         float64         `json:"price"`
	
	DistanceKM string  `json:"distance"`
	Distance   float64 `json:"-"`
}
