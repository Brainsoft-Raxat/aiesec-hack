package repository

import (
	"context"
	"time"

	"github.com/Brainsoft-Raxat/aiesec-hack/internal/models"
	"github.com/google/uuid"
)

func (r *postgres) CreatePromotion(ctx context.Context, promotion models.Promotion) (uuid.UUID, error) {
	query := `
        INSERT INTO promotions (title, banner_url, reviews_number, reviews_rate, expires, discount, city, address, latitude, longitude, price)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
        RETURNING id
    `

	var promotionID uuid.UUID
	err := r.db.QueryRow(ctx, query,
		promotion.Title, promotion.BannerURL, promotion.ReviewsNumber, promotion.ReviewsRate, promotion.Expires,
		promotion.Discount, promotion.City, promotion.Address, promotion.Latitude, promotion.Longitude, promotion.Price).
		Scan(&promotionID)
	if err != nil {
		return promotionID, err
	}

	return promotionID, nil
}

// GetPromotionsFiltered fetches promotions filtered by city.
func (r *postgres) GetPromotionsFiltered(ctx context.Context, city string) ([]models.Promotion, error) {
	query := `
        SELECT id, title, banner_url, reviews_number, reviews_rate, expires, discount, city, address, latitude, longitude, price
        FROM promotions
        WHERE city = $1 AND expires >= $2
    `

	rows, err := r.db.Query(ctx, query, city, time.Now())
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var promotions []models.Promotion
	for rows.Next() {
		var promotion models.Promotion
		err := rows.Scan(
			&promotion.ID, &promotion.Title, &promotion.BannerURL, &promotion.ReviewsNumber, &promotion.ReviewsRate,
			&promotion.Expires, &promotion.Discount, &promotion.City, &promotion.Address, &promotion.Latitude,
			&promotion.Longitude, &promotion.Price,
		)
		if err != nil {
			return nil, err
		}
		promotions = append(promotions, promotion)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return promotions, nil
}
