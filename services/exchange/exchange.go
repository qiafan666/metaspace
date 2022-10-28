package exchange

import (
	"context"
	"github.com/blockfishio/metaspace-backend/common"
	"github.com/blockfishio/metaspace-backend/redis"
	slog "github.com/qiafan666/quickweb/commons/log"
	"sync"
	"time"
)

var ins Exchange
var exchangeOnce sync.Once

type BaseExchange interface {
	ExchangeRate(ctx context.Context, quote string, base string) (float64, error)
}

type Exchange struct {
	exchange map[string]BaseExchange
	redis    redis.Dao
}

func (e *Exchange) Rate(ctx context.Context, quote string, base string) (float64, error) {

	result, err := e.redis.GetExchangePrice(ctx, quote, base)
	if err != nil {
		slog.Slog.InfoF(ctx, "Exchange redis error %s", err.Error())
	}
	if result.ExpireTime.After(time.Now()) {
		return result.Price, nil
	} else {
		_, err = e.exchange[common.CoinMarket].ExchangeRate(ctx, quote, base)
		if err != nil {
			return 0, err
		}
		return result.Price, nil
	}

}

func NewExchangeService() Exchange {
	exchangeOnce.Do(func() {
		ins.redis = redis.Instance()
		ins.exchange = map[string]BaseExchange{}
		ins.exchange[common.CoinMarket] = NewCoinMarketService()
	})

	return ins

}
