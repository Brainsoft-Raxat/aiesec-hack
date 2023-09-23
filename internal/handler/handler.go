package handler

import (
	"net/http"

	"github.com/Brainsoft-Raxat/aiesec-hack/internal/service"
	"github.com/Brainsoft-Raxat/aiesec-hack/pkg/apperror"
	"github.com/labstack/echo/v4"
)

const timeout = 60

type handler struct {
	service *service.Service
}

type Handler interface {
	Register(e *echo.Echo)
}

func New(services *service.Service) Handler {
	return &handler{service: services}
}

func (h *handler) Register(e *echo.Echo) {
	e.Use()
	api := e.Group("/api")
	{
		event := api.Group("/event")
		{
			event.POST("/create", h.CreateEvent)
			event.GET("/filter", h.GetEventsFiltered)
		}
		photoShoot := api.Group("/photo-shoot")
		{
			photoShoot.POST("/send", h.SendShoot)
		}
		promotion := api.Group("/promotion")
		{
			promotion.POST("/create", h.CreatePromotion)
			promotion.GET("/filter", h.GetPromotionsFiltered)
		}

	}
	auth := e.Group("/auth")
	{
		auth.POST("/sign-in", h.SignIn)
	}
}

func HandleEcho(c echo.Context, err error) error {
	if err == nil {
		return nil
	}

	if appErr := apperror.AsErrorInfo(err); appErr != nil {
		return c.JSON(appErr.Status, appErr)
	}

	return c.JSON(http.StatusInternalServerError, err)
}
