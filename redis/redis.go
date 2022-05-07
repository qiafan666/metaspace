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

	SetPublicKey(ctx context.Context, publicKey inner.PublicKey, expire time.Duration) (err error)
	GetPublicKey(ctx context.Context, apiKey string) (out inner.PublicKey, err error)
	DelPublicKey(ctx context.Context, apiKey string) (err error)

	SetRand(ctx context.Context, publicKey inner.Rand, expire time.Duration) (err error)
	GetRand(ctx context.Context, apiKey string) (out inner.Rand, err error)
	DelRand(ctx context.Context, apiKey string) (err error)

	SetAuthCode(ctx context.Context, publicKey inner.AuthCode, expire time.Duration) (err error)
	GetAuthCode(ctx context.Context, authCode string) (out inner.AuthCode, err error)
	DelAuthCode(ctx context.Context, authCode string) (err error)

	SetThirdPartyToken(ctx context.Context, thirdPartyToken inner.ThirdPartyToken, expire time.Duration) (err error)
	GetThirdPartyToken(ctx context.Context, thirdPartyId string) (out inner.ThirdPartyToken, err error)
	DelThirdPartyToken(ctx context.Context, thirdPartyId string) (err error)

	SetTokenUser(ctx context.Context, TokenUser inner.TokenUser, expire time.Duration) (err error)
	GetTokenUser(ctx context.Context, thirdPartyLoginId string, token string) (out inner.TokenUser, err error)
	DelTokenUser(ctx context.Context, token string) (err error)
}

type Imp struct {
	redis *redis.Client
}

func Instance() Dao {
	return &Imp{redis: cornus.GetCornusInstance().Redis("metaspace")}
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

func (i Imp) GetPublicKey(ctx context.Context, apiKey string) (out inner.PublicKey, err error) {
	result := i.redis.Get(ctx, fmt.Sprintf(common.ThirdPartyPublicKey, apiKey))
	if result.Err() != nil {
		return out, result.Err()
	}
	err = json.Unmarshal([]byte(result.Val()), &out)
	return
}

func (i Imp) SetPublicKey(ctx context.Context, publicKey inner.PublicKey, expire time.Duration) (err error) {
	marshal, err := json.Marshal(publicKey)
	if err != nil {
		return err
	}
	return i.redis.SetNX(ctx, fmt.Sprintf(common.ThirdPartyPublicKey, publicKey.ApiKey), marshal, expire).Err()
}

func (i Imp) DelPublicKey(ctx context.Context, apiKey string) (err error) {
	return i.redis.Del(ctx, fmt.Sprintf(common.ThirdPartyPublicKey, apiKey)).Err()
}

func (i Imp) GetRand(ctx context.Context, apiKey string) (out inner.Rand, err error) {
	result := i.redis.Get(ctx, fmt.Sprintf(common.ThirdPartyRand, apiKey))
	if result.Err() != nil {
		return out, result.Err()
	}
	err = json.Unmarshal([]byte(result.Val()), &out)
	return
}

func (i Imp) SetRand(ctx context.Context, publicKey inner.Rand, expire time.Duration) (err error) {
	marshal, err := json.Marshal(publicKey)
	if err != nil {
		return err
	}
	return i.redis.SetEX(ctx, fmt.Sprintf(common.ThirdPartyRand, publicKey.ApiKey), marshal, expire).Err()
}

func (i Imp) DelRand(ctx context.Context, apiKey string) (err error) {
	return i.redis.Del(ctx, fmt.Sprintf(common.ThirdPartyRand, apiKey)).Err()
}

func (i Imp) GetAuthCode(ctx context.Context, authCode string) (out inner.AuthCode, err error) {
	result := i.redis.Get(ctx, fmt.Sprintf(common.ThirdPartyAuthCode, authCode))
	if result.Err() != nil {
		return out, result.Err()
	}
	err = json.Unmarshal([]byte(result.Val()), &out)
	return
}

func (i Imp) SetAuthCode(ctx context.Context, authCode inner.AuthCode, expire time.Duration) (err error) {
	marshal, err := json.Marshal(authCode)
	if err != nil {
		return err
	}
	return i.redis.SetEX(ctx, fmt.Sprintf(common.ThirdPartyAuthCode, authCode.AuthCode), marshal, expire).Err()
}

func (i Imp) DelAuthCode(ctx context.Context, authCode string) (err error) {
	return i.redis.Del(ctx, fmt.Sprintf(common.ThirdPartyAuthCode, authCode)).Err()
}

func (i Imp) GetThirdPartyToken(ctx context.Context, thirdPartyId string) (out inner.ThirdPartyToken, err error) {
	result := i.redis.Get(ctx, fmt.Sprintf(common.ThirdPartyToken, thirdPartyId))
	if result.Err() != nil {
		return out, result.Err()
	}
	err = json.Unmarshal([]byte(result.Val()), &out)
	return
}

func (i Imp) SetThirdPartyToken(ctx context.Context, thirdPartyToken inner.ThirdPartyToken, expire time.Duration) (err error) {
	marshal, err := json.Marshal(thirdPartyToken)
	if err != nil {
		return err
	}
	return i.redis.SetEX(ctx, fmt.Sprintf(common.ThirdPartyToken, thirdPartyToken.ThirdPartyPublicId), marshal, expire).Err()
}

func (i Imp) DelThirdPartyToken(ctx context.Context, thirdPartyId string) (err error) {
	return i.redis.Del(ctx, fmt.Sprintf(common.ThirdPartyToken, thirdPartyId)).Err()
}

func (i Imp) GetTokenUser(ctx context.Context, thirdPartyLoginId string, token string) (out inner.TokenUser, err error) {
	result := i.redis.HGet(ctx, thirdPartyLoginId, fmt.Sprintf(common.ThirdPartyTokenUser, token))
	if result.Err() != nil {
		return out, result.Err()
	}
	err = json.Unmarshal([]byte(result.Val()), &out)
	return
}

func (i Imp) SetTokenUser(ctx context.Context, tokenUser inner.TokenUser, expire time.Duration) (err error) {
	marshal, err := json.Marshal(tokenUser)
	if err != nil {
		return err
	}
	err = i.redis.Expire(ctx, tokenUser.ThirdPartyPublicId, expire).Err()
	if err != nil {
		return err
	}
	return i.redis.HSet(ctx, tokenUser.ThirdPartyPublicId, fmt.Sprintf(common.ThirdPartyTokenUser, tokenUser.Token), tokenUser.UserId, marshal).Err()
}

func (i Imp) DelTokenUser(ctx context.Context, token string) (err error) {
	return i.redis.Del(ctx, fmt.Sprintf(common.ThirdPartyTokenUser, token)).Err()
}
