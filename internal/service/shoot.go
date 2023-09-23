package service

import (
	"context"
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/Brainsoft-Raxat/aiesec-hack/internal/repository"
	"github.com/Brainsoft-Raxat/aiesec-hack/pkg/apperror"
	"github.com/Brainsoft-Raxat/aiesec-hack/pkg/data"
	"github.com/Brainsoft-Raxat/aiesec-hack/pkg/errcodes"
)

type shootService struct {
	repo *repository.Repository
}

// SendShoot implements ShootService.
func (s *shootService) SendShoot(ctx context.Context, request data.SendShootRequest) (resp data.SendShootResponse, err error) {
	err = s.repo.SMTP.SendEmailWithAttachment(ctx, request.FileData, generateFileName(request.FileName, time.Now()), request.ToEmail)
	if err != nil {
		return resp, apperror.NewErrorInfo(ctx, errcodes.InternalServerError, fmt.Sprintf("unable to send email: %s", err.Error()))
	}

	resp.Status = "ok"

	return
}

func NewShootService(repo *repository.Repository) ShootService {
	return &shootService{
		repo: repo,
	}
}

func generateFileName(originalName string, currentDate time.Time) string {
	// Format the current date as a string, for example: "2006-01-02"
	dateString := currentDate.Format("2006-01-02")

	// Remove any file extension from the original name
	extension := filepath.Ext(originalName)
	baseName := strings.TrimSuffix(originalName, extension)

	// Combine the formatted date and the original file name
	fileName := fmt.Sprintf("%s_%s%s", baseName, dateString, extension)

	return fileName
}
