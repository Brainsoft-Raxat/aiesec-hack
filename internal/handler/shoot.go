package handler

import (
	"context"
	"io"
	"net/http"
	"time"

	"github.com/Brainsoft-Raxat/aiesec-hack/pkg/apperror"
	"github.com/Brainsoft-Raxat/aiesec-hack/pkg/data"
	"github.com/Brainsoft-Raxat/aiesec-hack/pkg/errcodes"
	"github.com/labstack/echo/v4"
)

func (h *handler) SendShoot(c echo.Context) error {
    ctx, cancel := context.WithTimeout(c.Request().Context(), timeout*time.Second)
    defer cancel()

    var req data.SendShootRequest

    // Retrieve the "file" field from the form data
    file, err := c.FormFile("file")
    if err != nil {
        return HandleEcho(c, apperror.NewErrorInfo(ctx, errcodes.InvalidFile, err.Error()))
    }

    // Retrieve the "to_email" field from the form data
    toEmail := c.FormValue("to_email")
    if toEmail == "" {
        return HandleEcho(c, apperror.NewErrorInfo(ctx, errcodes.ValidationError, "Missing 'to_email' field"))
    }

    req.FileName = file.Filename
    req.ToEmail = toEmail

    src, err := file.Open()
    if err != nil {
        return HandleEcho(c, apperror.NewErrorInfo(ctx, errcodes.InternalServerError, err.Error()))
    }
    defer src.Close()

    req.FileData, err = io.ReadAll(src)
    if err != nil {
        return HandleEcho(c, apperror.NewErrorInfo(ctx, errcodes.InternalServerError, err.Error()))
    }

    resp, err := h.service.ShootService.SendShoot(ctx, req)
    if err != nil {
        return HandleEcho(c, err)
    }

    return c.JSON(http.StatusOK, resp)
}

