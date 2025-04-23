package repository

import (
	"github.com/redis/go-redis/v9"
)

type ChatCacheRepository struct {
	*redis.Client
}

func NewChatCacheRepository(client *redis.Client) *ChatCacheRepository {
	return &ChatCacheRepository{client}
}

