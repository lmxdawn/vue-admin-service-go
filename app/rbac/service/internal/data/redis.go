package data

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func likeKey(id int64) string {
	return fmt.Sprintf("like:%d", id)
}

func (r *adminRepo) GetArticleLike(ctx context.Context, id int64) (rv int64, err error) {
	get := r.data.rdb.Get(ctx, likeKey(id))
	rv, err = get.Int64()
	if err == redis.Nil {
		return 0, nil
	}
	return
}

func (r *adminRepo) IncArticleLike(ctx context.Context, id int64) error {
	_, err := r.data.rdb.Incr(ctx, likeKey(id)).Result()
	return err
}
