package exchange

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/qiafan666/metaspace/pojo/inner"
	"github.com/qiafan666/metaspace/redis"
	"github.com/qiafan666/quickweb"
	slog "github.com/qiafan666/quickweb/commons/log"
	"github.com/valyala/fasthttp"
	"golang.org/x/sync/singleflight"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var sg singleflight.Group
var coinMarketConfig struct {
	CoinMarket struct {
		Address string `yaml:"address"`
		ApiKey  string `yaml:"api_key"`
	} `yaml:"coinMarket"`
}

func init() {
	cornus.GetCornusInstance().LoadCustomizeConfig(&coinMarketConfig)
}

type coinMarketService struct {
	redis       redis.Dao
	requestPool sync.Pool
}

func NewCoinMarketService() BaseExchange {
	imp := coinMarketService{redis: redis.Instance()}
	imp.requestPool.New = func() any {
		return CoinMarketRequest{}
	}
	sg = singleflight.Group{}
	return &imp
}
func (c *coinMarketService) ExchangeRate(ctx context.Context, quote string, base string) (code float64, err error) {

	_, err, _ = sg.Do(strconv.FormatInt(int64(time.Now().Hour()), 10), func() (interface{}, error) {

		req := fasthttp.AcquireRequest()
		defer fasthttp.ReleaseRequest(req)

		req.Header.SetMethod(http.MethodGet)
		//set authorization header
		req.Header.SetContentType("application/json")
		req.Header.Add("X-CMC_PRO_API_KEY", coinMarketConfig.CoinMarket.ApiKey)

		requestUri := fasthttp.AcquireURI()
		defer fasthttp.ReleaseURI(requestUri)
		err = requestUri.Parse(nil, []byte(coinMarketConfig.CoinMarket.Address))
		if err != nil {
			slog.Slog.ErrorF(ctx, "coinMarketService url parse error %s", err.Error())
			return 0, err
		}

		requestURL := requestUri.String() + "?symbol=" + quote + "&amount=1" + "&convert=" + base
		req.SetRequestURI(requestURL)
		resp := fasthttp.AcquireResponse()
		defer fasthttp.ReleaseResponse(resp)

		err = fasthttp.DoTimeout(req, resp, 30*time.Second)
		if err != nil {
			slog.Slog.ErrorF(ctx, "coinMarketService ExchangeRate error %s", err.Error())
			return 0, err
		}
		body := resp.Body()
		request := c.requestPool.Get().(CoinMarketRequest)
		err = json.Unmarshal(body, &request)
		if err != nil {
			slog.Slog.ErrorF(ctx, "coinMarketService Unmarshal error %s", err.Error())
			return 0, err
		}
		if request.Status.ErrorCode != 0 {
			slog.Slog.ErrorF(ctx, "coinMarketService request error %s", err.Error())
			return 0, errors.New(request.Status.ErrorMessage.(string))
		}

		err = c.redis.SetExchangePrice(ctx, inner.ExchangePrice{
			Quote:      quote,
			Base:       base,
			Price:      request.Data[0].Quote[base].Price,
			ExpireTime: time.Now().Add(15 * time.Minute),
		}, 0)
		if err != nil {
			slog.Slog.ErrorF(ctx, "coinMarketService SetExchangePrice error %s", err.Error())
			return 0, err
		}
		return request.Data[0].Quote[base].Price, nil
	})
	if err != nil {
		slog.Slog.ErrorF(ctx, "coinMarketService sg error %s", err.Error())
		return 0, err
	}

	result, err := c.redis.GetExchangePrice(ctx, quote, base)
	if err == nil && result.Price != 0 {
		sg.Forget(strconv.FormatInt(int64(time.Now().Hour()), 10))
	}
	return result.Price, nil
}
