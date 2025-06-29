package repository

import "context"

type CacheRepositoryInterface interface {
	GetJSON(ctx context.Context, key string, dest interface{}) error
	SetJSON(ctx context.Context, key string, value interface{}) error
	Delete(ctx context.Context, key string) error
}
