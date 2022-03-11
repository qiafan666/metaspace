package redis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/blockfishio/metaspace-backend/common"
	"github.com/blockfishio/metaspace-backend/pojo/inner"
	"github.com/go-redis/redis/v8"
	"github.com/jau1jz/cornus"
	"time"
)

var Nil = errors.New("redis: nil")

type Dao interface {
	SetNonce(ctx context.Context, nonce inner.Nonce, expire time.Duration) (err error)
	GetNonce(ctx context.Context, uuid string) (out inner.Nonce, err error)
	DelNonce(ctx context.Context, uuid string) (err error)
}

type Imp struct {
	redis *redis.Client
}

func (i Imp) GetNonce(ctx context.Context, address string) (out inner.Nonce, err error) {
	result := i.redis.Get(ctx, fmt.Sprintf(common.UserNonce, address))
	if result.Err() != nil {
		return out, result.Err()
	}
	err = json.Unmarshal([]byte(result.Val()), &out)
	return
}

func (i Imp) SetNonce(ctx context.Context, nonce inner.Nonce, expire time.Duration) (err error) {
	marshal, err := json.Marshal(nonce)
	if err != nil {
		return err
	}
	return i.redis.SetEX(ctx, fmt.Sprintf(common.UserNonce, nonce.Address), marshal, expire).Err()
}

func (i Imp) DelNonce(ctx context.Context, uuid string) (err error) {
	return i.redis.Del(ctx, fmt.Sprintf(common.UserNonce, uuid)).Err()
}
func Instance() Dao {
	return &Imp{redis: cornus.GetCornusInstance().Redis("metaspace")}
}
