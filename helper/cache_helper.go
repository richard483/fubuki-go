package helper

import (
	"context"

	repository "fubuki-go/repository"
)

func GetTyped[T any](r repository.CacheRepositoryInterface, ctx context.Context, key string) (T, error) {
	var result T
	err := r.GetJSON(ctx, key, &result)
	return result, err
}
