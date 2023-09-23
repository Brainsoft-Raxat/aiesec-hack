package handler

import (
	"net/http"

	"github.com/Brainsoft-Raxat/aiesec-hack/pkg/data"
	"github.com/labstack/echo/v4"
)

func (h *handler) HandleSomething(c echo.Context) error {
	var req data.DoSomethingRequest
	err := c.Bind(&req)
	if err != nil {
		return HandleEcho(c, err)
	}

	ctx := c.Request().Context()
	resp, err := h.service.SomeService.DoSomething(ctx, req)
	if err != nil {
		return HandleEcho(c, err)
	}

	return c.JSON(http.StatusOK, resp)
}
