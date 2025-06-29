package impl

import (
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
)

type CacheRepository struct {
	*redis.Client
}

func NewCacheRepository(client *redis.Client) *CacheRepository {
	return &CacheRepository{client}
}

func (r *CacheRepository) GetJSON(ctx context.Context, key string, dest interface{}) error {
	result := r.Client.Get(ctx, key)
	if err := result.Err(); err != nil {
		return err
	}
	return json.Unmarshal([]byte(result.Val()), dest)
}

func (r *CacheRepository) SetJSON(ctx context.Context, key string, value interface{}) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return r.Client.Set(ctx, key, data, 0).Err()
}

func (r *CacheRepository) Delete(ctx context.Context, key string) error {
	return r.Client.Del(ctx, key).Err()
}
