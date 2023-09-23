package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/Brainsoft-Raxat/aiesec-hack/pkg/data"
	"github.com/labstack/echo/v4"
)

func (h *handler) CreatePromotion(c echo.Context) error {
	var req data.CreatePromotionRequest
	err := c.Bind(&req)
	if err != nil {
		return HandleEcho(c, err)
	}

	ctx, cancel := context.WithTimeout(c.Request().Context(), timeout*time.Second)
	defer cancel()

	resp, err := h.service.PromotionService.CreatePromotion(ctx, req)
	if err != nil {
		return HandleEcho(c, err)
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *handler) GetPromotionsFiltered(c echo.Context) error {
	var req data.GetPromotionsFilteredRequest

	req.JerryID = c.QueryParam("jerry_id")

	ctx, cancel := context.WithTimeout(c.Request().Context(), timeout*time.Second)
	defer cancel()

	resp, err := h.service.PromotionService.GetPromotionsFiltered(ctx, req)
	if err != nil {
		return HandleEcho(c, err)
	}

	return c.JSON(http.StatusOK, resp)
}