package model

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
