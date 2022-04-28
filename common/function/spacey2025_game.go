package function

import (
	"errors"
	"github.com/blockfishio/metaspace-backend/common"
	"strings"
	"unsafe"
)

func GetCategoryString(mType int64) string {
	switch mType {
	case common.CategoryId0:
		return common.ChestString
	case common.CategoryId1:
		return common.TicketString
	case common.CategoryId2:
		return common.LandString
	case common.CategoryId3:
		return common.BuildingString
	case common.CategoryId4:
		return common.TowerString
	case common.CategoryId5:
		return common.TrapString
	case common.CategoryId6:
		return common.TrapString
	}

	return ""
}

func GetCategoryId(Category string) int64 {
	switch Category {
	case strings.ToLower(common.DummyString):
		return common.CategoryId0
	case strings.ToLower(common.ChestString):
		return common.CategoryId1
	case strings.ToLower(common.TicketString):
		return common.CategoryId2
	case strings.ToLower(common.LandString):
		return common.CategoryId3
	case strings.ToLower(common.BuildingString):
		return common.CategoryId4
	case strings.ToLower(common.TowerString):
		return common.CategoryId5
	case strings.ToLower(common.TrapString):
		return common.CategoryId6
	}

	return -1
}

func GetRarityString(mType int64) string {
	switch mType {
	case common.SubCategoryId0:
		return common.CommonString
	case common.SubCategoryId1:
		return common.UncommonString
	case common.SubCategoryId2:
		return common.RareString
	case common.SubCategoryId3:
		return common.EpicString
	case common.SubCategoryId4:
		return common.LegendaryString
	case common.SubCategoryId5:
		return common.JunkString
	}

	return ""
}

func GetSubcategoryString(Category int64, subCategory int64) (string, error) {
	if Category == common.CategoryId0 { //Dummy
		return "", errors.New(" subCategory data Not found ")
	} else if Category == common.CategoryId1 { //Chest
		return "", errors.New(" subCategory data Not found ")
	} else if Category == common.CategoryId2 { //Ticket
		return "", errors.New(" subCategory data Not found ")
	} else if Category == common.CategoryId3 { //Land
		return "", errors.New(" subCategory data Not found ")
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
	} else {
		return "", errors.New(" GetSubcategoryString data By db Not found")
	}
}

func GetSubcategoryByNftString(Category string, subCategory int64) (string, error) {
	if Category == strings.ToLower(common.DummyString) { //Dummy
		return "", errors.New(" subCategory data By Nft Not found ")
	} else if Category == strings.ToLower(common.ChestString) { //Chest
		return "", errors.New(" subCategory data By Nft Not found ")
	} else if Category == strings.ToLower(common.TicketString) { //Ticket
		return "", errors.New(" subCategory data By Nft Not found ")
	} else if Category == strings.ToLower(common.DummyString) { //Land
		return "", errors.New(" subCategory data By Nft Not found ")
	} else if Category == strings.ToLower(common.BuildingString) { //Building
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
			return "", errors.New(" subCategory data By Nft Not found ")
		}
	} else if Category == strings.ToLower(common.TowerString) { //Tower
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
		default:
			return "", errors.New(" subCategory data By Nft Not found ")
		}
	} else if Category == strings.ToLower(common.TrapString) { //Trap
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
			return "", errors.New(" subCategory data By Nft Not found ")
		}
	} else {
		return "", errors.New(" GetSubcategoryString data By Nft Not found")
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

func Byte2(b []byte) [32]byte {

	return *(*[32]byte)(unsafe.Pointer(&b))
}
