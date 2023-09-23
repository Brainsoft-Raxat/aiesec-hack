package repository

import (
	"context"

	"github.com/Brainsoft-Raxat/aiesec-hack/internal/app/config"
	"github.com/Brainsoft-Raxat/aiesec-hack/internal/app/conn"
	"github.com/Brainsoft-Raxat/aiesec-hack/internal/models"
	"github.com/google/uuid"
)

type Repository struct {
	Postgres
	JerryStore
	SMTP
}

type Postgres interface {
	DoSomething(ctx context.Context)
	GetUserCredsByLogin(login string) (int64, string, error)
	CreateEvent(ctx context.Context, event models.Event) (uuid.UUID, error)
	// GetEventByID(ctx context.Context, id int) (models.Event, error)
	GetEventsFiltered(ctx context.Context, city string, categories []string) ([]models.Event, error)
	UpdateEvent(ctx context.Context, event models.Event) error
	DeleteEvent(ctx context.Context, id int) error
	CreatePromotion(ctx context.Context, promotion models.Promotion) (uuid.UUID, error)
	GetPromotionsFiltered(ctx context.Context, city string) ([]models.Promotion, error)
}

type JerryStore interface {
	GetJerryByID(ctx context.Context, id string) (models.Jerry, error)
}

type SMTP interface {
	SendEmailWithAttachment(ctx context.Context, fileData []byte, fileName, toEmail string) error
}

func New(conn conn.Conn, cfg *config.Config) *Repository {
	return &Repository{
		Postgres:   NewPostgresRepository(conn.DB, cfg.Postgres),
		JerryStore: NewJerryStore(),
		SMTP:       NewSMTPRepository(cfg.SMTP),
	}
}
