package handler

import (
	"context"
	"net/http"
	"time"

	"github.com/Brainsoft-Raxat/aiesec-hack/pkg/data"
	"github.com/labstack/echo/v4"
)

func (h *handler) CreateEvent(c echo.Context) error {
	var req data.CreateEventRequest
	err := c.Bind(&req)
	if err != nil {
		return HandleEcho(c, err)
	}

	ctx, cancel := context.WithTimeout(c.Request().Context(), timeout*time.Second)
	defer cancel()

	resp, err := h.service.EventService.CreateEvent(ctx, req)
	if err != nil {
		return HandleEcho(c, err)
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *handler) GetEventsFiltered(c echo.Context) error {
	var req data.GetEventsFilteredRequest

	req.JerryID = c.QueryParam("jerry_id")
	req.City = c.QueryParam("city")
	req.Categories = c.QueryParams()["categories"]

	ctx, cancel := context.WithTimeout(c.Request().Context(), timeout*time.Second)
	defer cancel()

	resp, err := h.service.EventService.GetEventsFiltered(ctx, req)
	if err != nil {
		return HandleEcho(c, err)
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *handler) UpdateEventCount(c echo.Context) error {
	var req data.UpdateEventCountRequest
	if err := c.Bind(&req); err != nil {
		return HandleEcho(c, err)
	}

	ctx, cancel := context.WithTimeout(c.Request().Context(), timeout*time.Second)
	defer cancel()

	resp, err := h.service.EventService.UpdateEventCount(ctx, req)
	if err != nil {
		return HandleEcho(c, err) // Handle service errors
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *handler) GetSuggestion(c echo.Context) error {
	var req data.GiveSuggestionRequest

	req.JerryID = c.QueryParam("jerry_id")

	ctx, cancel := context.WithTimeout(c.Request().Context(), timeout*time.Second)
	defer cancel()

	resp, err := h.service.EventService.GiveSuggestion(ctx, req)
	if err != nil {
		return HandleEcho(c, err)
	}

	return c.JSON(http.StatusOK, resp)
}