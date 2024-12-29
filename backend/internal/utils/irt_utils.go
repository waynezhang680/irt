package utils

import (
	"math"
)

// CalculateItemInformation 计算题目信息量
func CalculateItemInformation(ability, difficulty, discrimination, guessParameter float64) float64 {
	const D = 1.7
	p := CalculateProbability(ability, difficulty, discrimination, guessParameter)
	q := 1 - p
	numerator := math.Pow(D*discrimination*(p-guessParameter), 2)
	denominator := (p - guessParameter) * q
	return numerator / denominator
}

// CalculateProbability 计算作答正确概率
func CalculateProbability(ability, difficulty, discrimination, guessParameter float64) float64 {
	const D = 1.7
	z := D * discrimination * (ability - difficulty)
	return guessParameter + (1-guessParameter)/(1+math.Exp(-z))
}

// CalculateStandardError 计算能力值的标准误
func CalculateStandardError(information float64) float64 {
	if information <= 0 {
		return math.Inf(1)
	}
	return 1 / math.Sqrt(information)
}
