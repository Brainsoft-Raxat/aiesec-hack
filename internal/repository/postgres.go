package repository

import (
	"context"
	"strconv"
	"strings"
	"time"

	"github.com/Brainsoft-Raxat/aiesec-hack/internal/app/config"
	"github.com/Brainsoft-Raxat/aiesec-hack/internal/models"
	"github.com/Brainsoft-Raxat/aiesec-hack/pkg/apperror"
	"github.com/Brainsoft-Raxat/aiesec-hack/pkg/errcodes"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4/pgxpool"
)

type postgres struct {
	db  *pgxpool.Pool
	cfg config.Postgres
}

func NewPostgresRepository(db *pgxpool.Pool, cfg config.Postgres) Postgres {
	return &postgres{
		db:  db,
		cfg: cfg,
	}
}

// GetUserCredsByLogin implements Postgres.
func (r *postgres) GetUserCredsByLogin(login string) (int64, string, error) {
	return 1, "asd", nil
}

// DoSomething implements Postgres.
func (r *postgres) DoSomething(ctx context.Context) {
	panic("unimplemented")
}

func (r *postgres) CreateEvent(ctx context.Context, event models.Event) (eventID uuid.UUID, err error) {
	query := `
        INSERT INTO events (title, description, banner_url, category, author, datetime, address, location, city)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
        RETURNING id
    `
	err = r.db.QueryRow(ctx, query,
		event.Title, event.Description, event.BannerURL, event.Category, event.Author, event.Datetime, event.Address, event.Location, event.City).
		Scan(&eventID)
	if err != nil {
		return eventID, apperror.NewErrorInfo(ctx, errcodes.CreateEventError, err.Error())
	}

	return eventID, nil
}

func (r *postgres) GetEventsFiltered(ctx context.Context, city string, categories []string) ([]models.Event, error) {
	var whereClause string
	var args []interface{}

	// Add the WHERE clause for city filtering if a city is provided
	if city != "" {
		whereClause += "WHERE city = $" + strconv.Itoa(len(args)+1) + " "
		args = append(args, city)
	}

	// Add the WHERE clause for category filtering if categories are provided
	if len(categories) > 0 {
		if city != "" {
			whereClause += "AND "
		} else {
			whereClause += "WHERE "
		}
		whereClause += "category IN ("
		for _, cat := range categories {
			whereClause += "$" + strconv.Itoa(len(args)+1) + ","
			args = append(args, cat)
		}
		whereClause = strings.TrimSuffix(whereClause, ",") + ") "
	}

	// Add the condition to filter events for today
	if whereClause != "" {
		whereClause += "AND "
	} else {
		whereClause += "WHERE "
	}
	whereClause += "datetime::date >= $" + strconv.Itoa(len(args)+1) + "::date"
	args = append(args, time.Now())

	query := `
        SELECT id, title, description, banner_url, category, author, datetime, address, location, city, count
        FROM events
    ` + whereClause + `
        ORDER BY datetime ASC
    `

	rows, err := r.db.Query(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []models.Event
	for rows.Next() {
		var event models.Event
		err := rows.Scan(
			&event.ID, &event.Title, &event.Description, &event.BannerURL, &event.Category, &event.Author,
			&event.Datetime, &event.Address, &event.Location, &event.City, &event.Count,
		)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return events, nil
}

func (r *postgres) UpdateEvent(ctx context.Context, event models.Event) error {
	query := `
        UPDATE events
        SET title = $1, description = $2, category = $3, author = $4, datetime = $5, location = $6, city = $7, banner_url = $8, address = $9  
        WHERE id = $8
    `
	_, err := r.db.Exec(ctx, query,
		event.Title, event.Description, event.Category, event.Author, event.Datetime, event.Location, event.City, event.ID, event.BannerURL, event.Address)
	return err
}

func (r *postgres) UpdateEventCount(ctx context.Context, id uuid.UUID) error {
	query := `
        UPDATE events
        SET count = count + 1
        WHERE id = $1
    `
	_, err := r.db.Exec(ctx, query, id)
	return err
}

func (r *postgres) DeleteEvent(ctx context.Context, id int) error {
	query := `DELETE FROM events WHERE id = $1`
	_, err := r.db.Exec(ctx, query, id)
	return err
}
