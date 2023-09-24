package connection

import (
	"context"

	"github.com/Brainsoft-Raxat/aiesec-hack/internal/app/config"
	"github.com/go-redis/redis/v8"
)

func DialRedis(ctx context.Context, cfg config.Redis) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     cfg.Addr,
		Password: cfg.Pass,
		DB:       cfg.DB,
	})
}
