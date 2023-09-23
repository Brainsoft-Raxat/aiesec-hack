package data

import (
	"time"

	"github.com/Brainsoft-Raxat/aiesec-hack/internal/models"
	"github.com/google/uuid"
)

type CreatePromotionRequest struct {
	Title         string    `json:"title"`
	BannerURL     string    `json:"banner_url"`
	ReviewsNumber int       `json:"reviews_number"`
	ReviewsRate   float64   `json:"reviews_rate"`
	Expires       time.Time `json:"expires"` // You can use a specific time format if needed
	Discount      int       `json:"discount"`
	City          string    `json:"city"`
	Address       string    `json:"address"`
	Latitude      float64   `json:"latitude"`
	Longitude     float64   `json:"longitude"`
	Price         float64   `json:"price"`
}

type CreatePromotionResponse struct {
	PromotionID uuid.UUID `json:"promotion_id"`
}

type GetPromotionsFilteredRequest struct {
	JerryID    string   `json:"jerry_id"`
	City       string   `json:"city"`
	Categories []string `json:"categories"`
}

type GetPromotionsFilteredResponse struct {
	Promotions []models.Promotion `json:"promotions"`
}
