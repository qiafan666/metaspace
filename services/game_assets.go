package bizservice

import (
	"encoding/json"

	"github.com/blockfishio/metaspace-backend/common/function"
	"github.com/blockfishio/metaspace-backend/dao"
	"github.com/blockfishio/metaspace-backend/model"
	"github.com/blockfishio/metaspace-backend/pojo/inner"
	"github.com/blockfishio/metaspace-backend/pojo/request"
	"github.com/blockfishio/metaspace-backend/pojo/response"
	"github.com/blockfishio/metaspace-backend/redis"
	"github.com/jau1jz/cornus/commons"
	slog "github.com/jau1jz/cornus/commons/log"

	// "gorm.io/gorm"
	"sync"
)

const (
	NftApiUrl = "http://nftapi.spacey2025.com/v1/nfts?owner="
)

// PortalService service layer interface
type GameAssetsService interface {
	GetGameAssets(info request.GetGameAssets) (out response.GetGameAssets, code commons.ResponseCode, err error)
}

var gameAssetsServiceIns *gameAssetsServiceImp
var gameAssetsServiceInitOnce sync.Once

func NewGameAssetsInstance() GameAssetsService {
	gameAssetsServiceInitOnce.Do(func() {
		gameAssetsServiceIns = &gameAssetsServiceImp{
			dao:   dao.Instance(),
			redis: redis.Instance(),
		}
	})
	return gameAssetsServiceIns
}

type gameAssetsServiceImp struct {
	dao   dao.Dao
	redis redis.Dao
}

func (p gameAssetsServiceImp) GetGameAssets(info request.GetGameAssets) (out response.GetGameAssets, code commons.ResponseCode, err error) {
	var user model.User

	err = p.dao.WithContext(info.Ctx).First([]string{model.UserColumns.UUID, model.UserColumns.WalletAddress}, map[string]interface{}{
		model.UserColumns.UUID: info.BaseRequest.BaseUUID,
	}, nil, &user)
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "gameAssetsServiceImp failed to fetch UUID. Error: %s", err.Error())
		return out, 0, err
	}
	// TODO: Find all ingame assets (unmited assets)

	// Find all onchain assets
	// Test: "0x47bfEf1eed74f02feBe37F39D3119dcc3feDa3A2"
	var URL = NftApiUrl + user.WalletAddress
	vNftResp := &inner.NftResp{}
	vNftRespInJson, err := function.HandleGet(URL)
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "gameAssetServiceImp failed to fetch onchain data. Error: %s", err.Error())
		return out, 0, err
	}
	err = json.Unmarshal(vNftRespInJson, vNftResp)
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "gameAssetServiceImp failed to parse onchain data. Error: %s", err.Error())
		return out, 0, err
	}
	// iterate the NftResp to generate the response
	for _, vNftDetail := range vNftResp.Data {
		out.Assets = append(out.Assets, response.AssetBody{
			IsNft:           true,
			TokenId:         vNftDetail.Nft.TokenId,
			ContractAddress: vNftDetail.Nft.ContractAddress,
			Name:            vNftDetail.Nft.Name,
			Image:           vNftDetail.Nft.Image,
			Description:     vNftDetail.Nft.Description,
		})
	}

	return
}
