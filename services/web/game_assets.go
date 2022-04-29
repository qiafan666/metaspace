package web

import (
	"encoding/json"
	"strconv"
	"strings"

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
	//NftApiUrl = "http://0.0.0.0:5000/v1/nfts?owner="
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

	assetsNum := 0

	var user model.User
	err = p.dao.WithContext(info.Ctx).First([]string{model.UserColumns.UUID, model.UserColumns.WalletAddress}, map[string]interface{}{
		model.UserColumns.UUID: info.BasePortalRequest.BaseUUID,
	}, nil, &user)
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "gameAssetsServiceImp failed to fetch UUID. Error: %s", err.Error())
		return out, 0, err
	}

	vWalletAddress := strings.ToLower(user.WalletAddress)
	// TODO: Find all ingame NFT assets (unmited assets)
	var vAssets []model.Assets
	err = p.dao.WithContext(info.Ctx).Find([]string{}, map[string]interface{}{
		model.AssetsColumns.Uid: vWalletAddress,
	}, nil, &vAssets)
	if err != nil {
		slog.Slog.ErrorF(info.Ctx, "gameAssetsServiceImp find assets Error: %s", err.Error())
		return out, 0, err
	}
	for _, vAsset := range vAssets {
		SubcategoryString, err := function.GetSubcategoryString(vAsset.Category, vAsset.Type)
		if err != nil {
			slog.Slog.ErrorF(info.Ctx, "gameAssetServiceImp SubcategoryString Error: %s", err.Error())
			return out, 0, err
		}

		out.Assets = append(out.Assets, response.AssetBody{
			IsNft:           false,
			TokenId:         strconv.FormatInt(vAsset.Id, 10),
			ContractAddress: "0xxxxx",
			ContrainChain:   "BSC",
			Name:            vAsset.Name,
			Image:           vAsset.Image,
			Description:     vAsset.Description,
			Category:        function.GetCategoryString(vAsset.Category),
			CategoryId:      vAsset.Category,
			Rarity:          function.GetRarityString(vAsset.Rarity),
			RarityId:        vAsset.Rarity,
			MintSignature:   vAsset.MintSignature,
			SubcategoryId:   vAsset.Type,
			Subcategory:     SubcategoryString,
		})
		assetsNum++
	}

	// Find all onchain assets
	// Test: "0x47bfEf1eed74f02feBe37F39D3119dcc3feDa3A2"
	var URL = NftApiUrl + "0x47bfEf1eed74f02feBe37F39D3119dcc3feDa3A2"
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

	//iterate the NftResp to generate the response
	for _, vNftDetail := range vNftResp.Data {

		SubcategoryString, err := function.GetSubcategoryByNftString(vNftDetail.Nft.Category, vNftDetail.Nft.Subcategory)
		if err != nil {
			slog.Slog.ErrorF(info.Ctx, "gameAssetServiceImp By Nft SubcategoryString Error: %s", err.Error())
			return out, 0, err
		}
		out.Assets = append(out.Assets, response.AssetBody{
			IsNft:           true,
			TokenId:         vNftDetail.Nft.TokenId,
			ContractAddress: vNftDetail.Nft.ContractAddress,
			ContrainChain:   "BSC",
			Name:            vNftDetail.Nft.Name,
			Image:           vNftDetail.Nft.Image,
			Description:     vNftDetail.Nft.Description,
			Category:        vNftDetail.Nft.Category,
			CategoryId:      function.GetCategoryId(vNftDetail.Nft.Category),
			Rarity:          function.GetRarityString(vNftDetail.Nft.Data.Tower.Rarity),
			RarityId:        vNftDetail.Nft.Data.Tower.Rarity,
			Subcategory:     SubcategoryString,
			SubcategoryId:   vNftDetail.Nft.Subcategory,
		})
		assetsNum++
	}
	out.AssetNum = assetsNum
	return
}
