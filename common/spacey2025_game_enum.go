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

// Rarity Type
type RarityType uint8

const (
	Common    RarityType = 1
	Uncommon  RarityType = 2
	Rare      RarityType = 3
	Epic      RarityType = 4
	Legendary RarityType = 5
	Junk      RarityType = 6
)

const (
	CommonString    = "Common"
	UncommonString  = "Uncommon"
	RareString      = "Rare"
	EpicString      = "Epic"
	LegendaryString = "Legendary"
	JunkString      = "Junk"
)
