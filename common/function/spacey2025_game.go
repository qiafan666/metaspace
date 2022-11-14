package function

import (
	"errors"
	"github.com/qiafan666/metaspace/common"
)

func GetCategoryString(mType int64) string {
	switch mType {
	case common.CategoryId0:
		return common.DummyString
	case common.CategoryId1:
		return common.ChestString
	case common.CategoryId2:
		return common.TicketString
	case common.CategoryId3:
		return common.LandString
	case common.CategoryId4:
		return common.BuildingString
	case common.CategoryId5:
		return common.TowerString
	case common.CategoryId6:
		return common.TrapString
	case common.CategoryId7:
		return common.ShipString
	}

	return ""
}

func GetRarityString(mType int64) string {
	switch mType {
	case common.RarityId0:
		return common.CommonString
	case common.RarityId1:
		return common.UncommonString
	case common.RarityId2:
		return common.RareString
	case common.RarityId3:
		return common.EpicString
	case common.RarityId4:
		return common.LegendaryString
	case common.RarityId5:
		return common.JunkString
	}

	return ""
}

func GetSubcategoryString(Category int64, subCategory int64) (string, error) {
	if Category == common.CategoryId0 { //Dummy
		return common.DummyString, nil
	} else if Category == common.CategoryId1 { //Chest
		return common.ChestString, nil
	} else if Category == common.CategoryId2 { //Ticket
		return common.TicketString, nil
	} else if Category == common.CategoryId3 { //Land
		return common.LandString, nil
	} else if Category == common.CategoryId4 { //Building
		switch subCategory {
		case int64(common.BuildingType1):
			return common.BuildingType1String, nil
		case int64(common.BuildingType2):
			return common.BuildingType2String, nil
		case int64(common.BuildingType3):
			return common.BuildingType3String, nil
		case int64(common.BuildingType4):
			return common.BuildingType4String, nil
		case int64(common.BuildingType5):
			return common.BuildingType5String, nil
		default:
			return "", errors.New(" subCategory data Not found ")
		}
	} else if Category == common.CategoryId5 { //Tower
		switch subCategory {
		case int64(common.TowerType1):
			return common.TowerType1String, nil
		case int64(common.TowerType2):
			return common.TowerType2String, nil
		case int64(common.TowerType3):
			return common.TowerType3String, nil
		case int64(common.TowerType4):
			return common.TowerType4String, nil
		case int64(common.TowerType5):
			return common.TowerType5String, nil
		case int64(common.TowerType6):
			return common.TowerType6String, nil
		case int64(common.TowerType7):
			return common.TowerType7String, nil
		case int64(common.TowerType8):
			return common.TowerType8String, nil
		default:
			return "", errors.New(" subCategory data Not found ")
		}
	} else if Category == common.CategoryId6 { //Trap
		switch subCategory {
		case int64(common.TrapType1):
			return common.TrapType1String, nil
		case int64(common.TrapType2):
			return common.TrapType2String, nil
		case int64(common.TrapType3):
			return common.TrapType3String, nil
		case int64(common.TrapType4):
			return common.TrapType4String, nil
		case int64(common.TrapType5):
			return common.TrapType5String, nil
		default:
			return "", errors.New(" subCategory data Not found ")
		}
	} else if Category == common.CategoryId7 { //Land
		return common.ShipString, nil
	} else {
		return "", errors.New(" GetSubcategoryString data By db Not found")
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
