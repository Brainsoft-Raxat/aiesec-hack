package repository

import (
	"context"
	"sync"

	"github.com/Brainsoft-Raxat/aiesec-hack/internal/models"
	"github.com/Brainsoft-Raxat/aiesec-hack/pkg/apperror"
	"github.com/Brainsoft-Raxat/aiesec-hack/pkg/errcodes"
)

type jerryStore struct {
	mu      sync.Mutex
	jerries []models.Jerry
}

// GetAllJerries implements JerryStore.
func (s *jerryStore) GetAllJerries(ctx context.Context) ([]models.Jerry, error) {
	return s.jerries, nil
}

// GetJerryByID implements JerryStore.
func (s *jerryStore) GetJerryByID(ctx context.Context, id string) (models.Jerry, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, jerry := range s.jerries {
		if jerry.ID == id {
			return jerry, nil
		}
	}

	return models.Jerry{}, apperror.NewErrorInfo(ctx, errcodes.JerryNotFound, "Jerry not found")
}

// TODO: fix to local cities
func NewJerryStore() JerryStore {
	return &jerryStore{
		jerries: []models.Jerry{
			{ID: "jerry1", Latitude: 43.261067, Longitude: 76.930945, City: "Almaty"},
			{ID: "jerry2", Latitude: 43.244243, Longitude: 76.959526, City: "Almaty"},
			{ID: "jerry3", Latitude: 51.129181, Longitude: 71.430724, City: "Astana"},
			{ID: "jerry4", Latitude: 51.089835, Longitude: 71.415226, City: "Astana"},
		},
	}
}
