package impl

import (
	"github.com/redis/go-redis/v9"
)

type CacheRepository struct {
	*redis.Client
}

func NewCacheRepository(client *redis.Client) *CacheRepository {
	return &CacheRepository{client}
}
