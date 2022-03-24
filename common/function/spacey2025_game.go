package function

import "github.com/blockfishio/metaspace-backend/common"

func GetCategoryString(mType common.AssetType) string {
	switch mType {
	case common.Chest:
		return common.ChestString
	case common.Building:
		return common.BuildingString
	case common.Tower:
		return common.TowerString
	case common.Trap:
		return common.TrapString
	}

	return ""
}

func GetRarityString(mType common.RarityType) string {
	switch mType {
	case common.Common:
		return common.CommonString
	case common.Uncommon:
		return common.UncommonString
	case common.Rare:
		return common.RareString
	case common.Epic:
		return common.EpicString
	case common.Legendary:
		return common.LegendaryString
	case common.Junk:
		return common.JunkString
	}

	return ""
}
