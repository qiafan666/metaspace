package model

type TowerStats struct {
	Attack       int     `json:"Attack"`
	Fire_Rate    float32 `json:"FireRate"`
	Attack_Range int     `json:"AttackRange"`
	Durability   int     `json:"Durability"`
}

type TowerConfig struct {
	AttackFactors       [3]float32
	AttackSpeedFactors  [3]float32
	AttackRangeFactors  [3]float32
	DurabilityFactors   [3]float32
	DefaultAttackPerSec float32
}

type RarityConfig struct {
	AttackFactor     float32
	DurabilityFactor float32
}
