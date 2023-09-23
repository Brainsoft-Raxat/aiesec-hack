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
			{ID: "jerry3", Latitude: 34.0522, Longitude: -118.2437, City: "Kokshetau"},
			{ID: "jerry4", Latitude: 41.8781, Longitude: -87.6298, City: "Shymkent"},
			{ID: "jerry5", Latitude: 51.5074, Longitude: -0.1278, City: "Aktau"},
		},
	}
}
