package repository

import (
	"fubuki-go/model"
)

type ChatCacheRepositoryInterface interface {
	AppendChatStream(key string, value model.Chat) error
	GetChatStream(key string) ([]model.Chat, error)
	ClearChatStream(key string) int64
}
