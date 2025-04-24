package repository

import (
	"context"
	"fubuki-go/model"
	"log"

	"github.com/redis/go-redis/v9"
)

type ChatCacheRepository struct {
	redisClient *redis.Client
	context.Context
}

func NewChatCacheRepository(client *redis.Client, ctx context.Context) *ChatCacheRepository {
	return &ChatCacheRepository{client, ctx}
}

func (r *ChatCacheRepository) AppendChatStream(key string, value model.Chat) error {
	_, err := r.redisClient.XAdd(r.Context, &redis.XAddArgs{
		Stream: key,
		Values: value,
		MaxLen: 100,
	}).Result()

	if err != nil {
		log.Println("#ERROR " + err.Error())
		return err
	}
	return nil
}

func (r *ChatCacheRepository) GetChatStream(key string) ([]model.Chat, error) {
	chatStream, err := r.redisClient.XRead(r.Context, &redis.XReadArgs{
		Streams: []string{key, "0"},
		Count:   100,
		Block:   0,
	}).Result()

	if err != nil {
		log.Println("#ERROR " + err.Error())
		return nil, err
	}

	var chats []model.Chat
	for _, stream := range chatStream {
		for _, message := range stream.Messages {
			chat := model.Chat{
				UserQuestion: message.Values["UserQuestion"].(string),
				ModelAnswer:  message.Values["ModelAnswer"].(string),
			}
			chats = append(chats, chat)
		}
	}

	return chats, nil

}

func (r *ChatCacheRepository) ClearChatStream(key string) int64 {
	delCount := r.redisClient.Del(r.Context, key).Val()
	return delCount

}
