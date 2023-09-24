package repository

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Brainsoft-Raxat/aiesec-hack/internal/models"
	"github.com/go-redis/redis/v8"
)

type redisRepo struct {
	client *redis.Client
}

func NewRedisRepository(client *redis.Client) Redis {
	return &redisRepo{
		client: client,
	}
}

func (r *redisRepo) CacheEvents(ctx context.Context, jerryID string, events []models.Event) error {
    pipe := r.client.Pipeline()
    defer pipe.Close()

    for _, event := range events {
        eventKey := fmt.Sprintf("event:%d", event.ID)
        eventJSON, err := json.Marshal(event)
        if err != nil {
            return err
        }

        pipe.HSet(ctx, eventKey, "data", eventJSON)

        jerrySetKey := fmt.Sprintf("jerry:%s", jerryID)
        pipe.SAdd(ctx, jerrySetKey, eventKey)
    }

    _, err := pipe.Exec(ctx)
    return err
}


func (r *redisRepo) GetEventsForJerryID(ctx context.Context, jerryID string) ([]models.Event, error) {
	jerrySetKey := fmt.Sprintf("jerry:%s", jerryID)
	eventKeys, err := r.client.SMembers(ctx, jerrySetKey).Result()
	if err != nil {
		return nil, err
	}

	var events []models.Event

	for _, eventKey := range eventKeys {
		eventData, err := r.client.HGet(ctx, eventKey, "data").Result()
		if err != nil {
			return nil, err
		}

		var event models.Event
		if err := json.Unmarshal([]byte(eventData), &event); err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}
