package function

import (
	"errors"
	"github.com/blockfishio/metaspace-backend/common"
)

func GetCategoryString(mType common.AssetType) string {
	switch mType {
	case common.Chest:
		return common.ChestString
	case common.Ticket:
		return common.TicketString
	case common.Land:
		return common.LandString
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

func GetSubcategoryString(Category string, subCategory string) (string, error) {
	if Category == string(common.Dummy) { //Dummy
		return "", errors.New(" subCategory data Not found ")
	} else if Category == string(common.Chest) { //Chest
		return "", errors.New(" subCategory data Not found ")
	} else if Category == string(common.Ticket) { //Ticket
		return "", errors.New(" subCategory data Not found ")
	} else if Category == string(common.Land) { //Land
		return "", errors.New(" subCategory data Not found ")
	} else if Category == string(common.Building) { //Building
		switch subCategory {
		case string(common.BuildingType1):
			return common.BuildingType1String, nil
		case string(common.BuildingType2):
			return common.BuildingType2String, nil
		case string(common.BuildingType3):
			return common.BuildingType3String, nil
		case string(common.BuildingType4):
			return common.BuildingType4String, nil
		case string(common.BuildingType5):
			return common.BuildingType5String, nil
		default:
			return "", errors.New(" subCategory data Not found ")
		}
	} else if Category == string(common.Tower) { //Tower
		switch subCategory {
		case string(common.TowerType1):
			return common.TowerType1String, nil
		case string(common.TowerType2):
			return common.TowerType2String, nil
		case string(common.TowerType3):
			return common.TowerType3String, nil
		case string(common.TowerType4):
			return common.TowerType4String, nil
		case string(common.TowerType5):
			return common.TowerType5String, nil
		default:
			return "", errors.New(" subCategory data Not found ")
		}
	} else if Category == string(common.Trap) { //Trap
		switch subCategory {
		case string(common.TrapType1):
			return common.TrapType1String, nil
		case string(common.TrapType2):
			return common.TrapType2String, nil
		case string(common.TrapType3):
			return common.TrapType3String, nil
		case string(common.TrapType4):
			return common.TrapType4String, nil
		case string(common.TrapType5):
			return common.TrapType5String, nil
		default:
			return "", errors.New(" subCategory data Not found ")
		}
	} else {
		return "", errors.New(" Category data Not found")
	}
}

func StringCheck(ins ...string) bool {
	for _, in := range ins {
		if len(in) <= 0 {
			return false
		}
	}
	return true
}
