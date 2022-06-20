package common

// Asset Type
type AssetType uint8

const (
	Dummy    AssetType = 0
	Chest    AssetType = 1
	Ticket   AssetType = 2
	Land     AssetType = 3
	Building AssetType = 4
	Tower    AssetType = 5
	Trap     AssetType = 6
)

const (
	DummyString    = "Dummy"
	ChestString    = "Chest"
	TicketString   = "Ticket"
	LandString     = "Land"
	BuildingString = "Building"
	TowerString    = "Tower"
	TrapString     = "Trap"
)

const (
	CategoryId0 int64 = 0
	CategoryId1 int64 = 1
	CategoryId2 int64 = 2
	CategoryId3 int64 = 3
	CategoryId4 int64 = 4
	CategoryId5 int64 = 5
	CategoryId6 int64 = 6
)

const (
	SubCategoryId0 int64 = 0
	SubCategoryId1 int64 = 1
	SubCategoryId2 int64 = 2
	SubCategoryId3 int64 = 3
	SubCategoryId4 int64 = 4
	SubCategoryId5 int64 = 5
)

const (
	CommonString    = "Common"
	UncommonString  = "Uncommon"
	RareString      = "Rare"
	EpicString      = "Epic"
	LegendaryString = "Legendary"
	JunkString      = "Junk"
)

// CategoryId Type
type CategoryIdInt uint8

const (
	CategoryId CategoryIdInt = 5
)

// RarityId Type
type RarityIdInt uint8

const (
	RarityId RarityIdInt = 5
)

// Land Type
const (
	LandTypeString = "Land"
)

// Building Type
type BuildingType uint8

const (
	BuildingType1 BuildingType = 1
	BuildingType2 BuildingType = 2
	BuildingType3 BuildingType = 3
	BuildingType4 BuildingType = 4
	BuildingType5 BuildingType = 5
)

const (
	BuildingType1String = "BuildingType1"
	BuildingType2String = "BuildingType2"
	BuildingType3String = "BuildingType3"
	BuildingType4String = "BuildingType4"
	BuildingType5String = "BuildingType5"
)

// Tower Type
type TowerType uint8

const (
	TowerType1 TowerType = 1
	TowerType2 TowerType = 2
	TowerType3 TowerType = 3
	TowerType4 TowerType = 4
	TowerType5 TowerType = 5
)

const (
	TowerType1String = "TowerType1"
	TowerType2String = "TowerType2"
	TowerType3String = "TowerType3"
	TowerType4String = "TowerType4"
	TowerType5String = "TowerType5"
)

// Trap Type
type TrapType uint8

const (
	TrapType1 TrapType = 1
	TrapType2 TrapType = 2
	TrapType3 TrapType = 3
	TrapType4 TrapType = 4
	TrapType5 TrapType = 5
)

const (
	TrapType1String = "TrapType1"
	TrapType2String = "TrapType2"
	TrapType3String = "TrapType3"
	TrapType4String = "TrapType4"
	TrapType5String = "TrapType5"
)

const (
	TowerTypeConfigs1 = "1"
	TowerTypeConfigs2 = "2"
	TowerTypeConfigs3 = "3"
	TowerTypeConfigs4 = "4"
	TowerTypeConfigs5 = "5"
)

const (
	RarityConfigs1 = "1"
	RarityConfigs2 = "2"
	RarityConfigs3 = "3"
	RarityConfigs4 = "4"
	RarityConfigs5 = "5"
	RarityConfigs6 = "6"
)
