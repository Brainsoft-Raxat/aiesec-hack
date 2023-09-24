package conn

import (
	"github.com/go-redis/redis/v8"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Conn struct {
	DB *pgxpool.Pool
	RedisClient *redis.Client
}
