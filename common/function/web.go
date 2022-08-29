package function

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jau1jz/cornus"
	slog "github.com/jau1jz/cornus/commons/log"
)

var config struct {
	ETHContract struct {
		Client string `yaml:"monitor_client"`
		Mint   string `yaml:"mint"`
		Assets string `yaml:"assets"`
		Ship   string `yaml:"ship"`
		Market string `yaml:"market"`
	} `yaml:"eth_contract"`
	BSCContract struct {
		Client string `yaml:"monitor_client"`
		Mint   string `yaml:"mint"`
		Assets string `yaml:"assets"`
		Ship   string `yaml:"ship"`
		Market string `yaml:"market"`
	} `yaml:"bsc_contract"`
	Chain struct {
		ETH uint8 `yaml:"eth"`
		BSC uint8 `yaml:"bsc"`
	} `yaml:"chain"`
}
var ethClient *ethclient.Client
var bscClient *ethclient.Client

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

}

func JudgeChain(chain uint8) (mint, ship, market, assets string, client *ethclient.Client, err error) {
	if chain == config.Chain.ETH {
		mint = config.ETHContract.Mint
		ship = config.ETHContract.Ship
		market = config.ETHContract.Market
		assets = config.ETHContract.Assets
		client = ethClient
		return
	} else if chain == config.Chain.BSC {
		mint = config.BSCContract.Mint
		ship = config.BSCContract.Ship
		market = config.ETHContract.Market
		assets = config.BSCContract.Assets
		client = bscClient
		return
	} else {
		return "", "", "", "", nil, errors.New("chain type error")
	}
}
