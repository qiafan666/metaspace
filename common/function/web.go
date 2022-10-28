package function

import (
	"context"
	"errors"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/qiafan666/quickweb"
	slog "github.com/qiafan666/quickweb/commons/log"
)

var config struct {
	ETHContract struct {
		Client string `yaml:"monitor_client"`
		Mint   string `yaml:"mint"`
		Assets string `yaml:"assets"`
		Ship   string `yaml:"ship"`
		Market string `yaml:"market"`
		Spay   string `yaml:"spay"`
	} `yaml:"eth_contract"`
	BSCContract struct {
		Client string `yaml:"monitor_client"`
		Mint   string `yaml:"mint"`
		Assets string `yaml:"assets"`
		Ship   string `yaml:"ship"`
		Market string `yaml:"market"`
		Spay   string `yaml:"spay"`
	} `yaml:"bsc_contract"`
	PolygonContract struct {
		Client string `yaml:"monitor_client"`
		Mint   string `yaml:"mint"`
		Assets string `yaml:"assets"`
		Ship   string `yaml:"ship"`
		Market string `yaml:"market"`
	} `yaml:"polygon_contract"`
	Chain struct {
		ETH     uint64 `yaml:"eth"`
		BSC     uint64 `yaml:"bsc"`
		Polygon uint64 `yaml:"polygon"`
	} `yaml:"chain"`
}
var ethClient *ethclient.Client
var bscClient *ethclient.Client
var polygonClient *ethclient.Client

func init() {
	cornus.GetCornusInstance().LoadCustomizeConfig(&config)

	var err error
	ethClient, err = ethclient.Dial(config.ETHContract.Client)
	if err != nil {
		slog.Slog.ErrorF(context.Background(), "eth client connect error %s", err.Error())
		panic(err.Error())
	}

	bscClient, err = ethclient.Dial(config.BSCContract.Client)
	if err != nil {
		slog.Slog.ErrorF(context.Background(), "eth client connect error %s", err.Error())
		panic(err.Error())
	}

	polygonClient, err = ethclient.Dial(config.PolygonContract.Client)
	if err != nil {
		slog.Slog.ErrorF(context.Background(), "eth client connect error %s", err.Error())
		panic(err.Error())
	}
}

func JudgeChain(chain uint64) (mint, ship, market, assets, spay string, client *ethclient.Client, err error) {
	if chain == config.Chain.ETH {
		mint = config.ETHContract.Mint
		ship = config.ETHContract.Ship
		market = config.ETHContract.Market
		assets = config.ETHContract.Assets
		spay = config.ETHContract.Spay
		client = ethClient
		return
	} else if chain == config.Chain.BSC {
		mint = config.BSCContract.Mint
		ship = config.BSCContract.Ship
		market = config.BSCContract.Market
		assets = config.BSCContract.Assets
		spay = config.BSCContract.Spay
		client = bscClient
		return
	} else if chain == config.Chain.Polygon {
		mint = config.PolygonContract.Mint
		ship = config.PolygonContract.Ship
		market = config.PolygonContract.Market
		assets = config.PolygonContract.Assets
		client = polygonClient
		return
	} else {
		return "", "", "", "", "", nil, errors.New("chain type error")
	}
}
