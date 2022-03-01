package redis

import (
	"context"
	"fmt"
	"github.com/blockfishio/metaspace-backend/common"
	"github.com/go-redis/redis/v8"
	"github.com/jau1jz/cornus"
	"time"
)

type Dao interface {
	SetNonce(ctx context.Context, uuid string, nonce string, expire time.Duration) (err error)
	GetNonce(ctx context.Context, uuid string) (nonce string, err error)
}

type Imp struct {
	redis *redis.Client
}

func (i Imp) GetNonce(ctx context.Context, uuid string) (nonce string, err error) {
	result := i.redis.Get(ctx, fmt.Sprintf(common.UserNonce, uuid))
	return result.String(), result.Err()
}

func (i Imp) SetNonce(ctx context.Context, uuid string, nonce string, expire time.Duration) (err error) {
	return i.redis.SetEX(ctx, fmt.Sprintf(common.UserNonce, uuid), nonce, expire).Err()
}

func Instance() Dao {
	return &Imp{redis: cornus.GetCornusInstance().Redis("metaspace")}
}
