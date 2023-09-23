package service

import (
	"context"
	"fmt"
	"sort"

	"github.com/Brainsoft-Raxat/aiesec-hack/internal/models"
	"github.com/Brainsoft-Raxat/aiesec-hack/internal/repository"
	"github.com/Brainsoft-Raxat/aiesec-hack/pkg/apperror"
	"github.com/Brainsoft-Raxat/aiesec-hack/pkg/data"
	"github.com/Brainsoft-Raxat/aiesec-hack/pkg/errcodes"
	"github.com/google/uuid"
)

type promotionService struct {
	repo *repository.Repository
}

// // DeletePromotion implements PromotionService.
// func (s *promotionService) DeletePromotion(ctx context.Context, request data.DeletePromotionRequest) (data.DeletePromotionResponse, error) {
// 	panic("unimplemented")
// }

// // GetPromotion implements PromotionService.festivali
// func (s *promotionService) GetPromotion(ctx context.Context, request data.GetPromotionRequest) (data.GetPromotionResponse, error) {
// 	panic("unimplemented")
// }

func (s *promotionService) GetPromotionsFiltered(ctx context.Context, request data.GetPromotionsFilteredRequest) (resp data.GetPromotionsFilteredResponse, err error) {
	jerry, err := s.repo.JerryStore.GetJerryByID(ctx, request.JerryID)
	if err != nil {
		return
	}

	promotions, err := s.repo.Postgres.GetPromotionsFiltered(ctx, jerry.City)
	if err != nil {
		return resp, apperror.NewErrorInfo(ctx, errcodes.InternalServerError, err.Error())
	}

	for i, promotion := range promotions {

		promotions[i].Distance = Haversine(models.Coordinates{
			Latitude:  jerry.Latitude,
			Longitude: jerry.Longitude,
		},
			models.Coordinates{
				Latitude:  promotion.Latitude,
				Longitude: promotion.Longitude,
			},
		)

		promotions[i].DistanceKM = fmt.Sprintf("%.1f km", promotions[i].Distance)
	}

	sort.Slice(promotions, func(i, j int) bool {
		return promotions[i].Distance < promotions[j].Distance
	})

	resp.Promotions = promotions

	return
}

// // UpdatePromotion implements PromotionService.
// func (s *promotionService) UpdatePromotion(ctx context.Context, request data.UpdatePromotionRequest) (data.UpdatePromotionResponse, error) {
// 	panic("unimplemented")
// }

// CreatePromotion implements PromotionService.
func (s *promotionService) CreatePromotion(ctx context.Context, request data.CreatePromotionRequest) (resp data.CreatePromotionResponse, err error) {
	resp.PromotionID, err = s.repo.Postgres.CreatePromotion(ctx, models.Promotion{
		ID:            uuid.New(),
		Title:         request.Title,
		BannerURL:     request.BannerURL,
		ReviewsNumber: request.ReviewsNumber,
		ReviewsRate:   request.ReviewsRate,
		Expires:       request.Expires,
		Discount:      request.Discount,
		City:          request.City,
		Address:       request.Address,
		Latitude:      request.Latitude,
		Longitude:     request.Longitude,
		Price:         request.Price,
	})
	if err != nil {
		return
	}

	return
}

func NewPromotionService(repo *repository.Repository) PromotionService {
	return &promotionService{
		repo: repo,
	}
}
