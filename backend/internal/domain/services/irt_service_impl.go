package services

import (
	"math"
)

type IRTServiceImpl struct {
	// IRT参数配置
	maxIterations int     // 最大迭代次数
	convergence   float64 // 收敛阈值
}

func NewIRTService() IRTService {
	return &IRTServiceImpl{
		maxIterations: 50,
		convergence:   0.001,
	}
}

// EstimateAbility 使用三参数IRT模型估计考生能力值
// P(θ) = c + (1-c)/(1 + e^(-Da(θ-b)))
// θ: 能力值
// a: 区分度
// b: 难度
// c: 猜测参数
// D: 常数(通常取1.7)
func (s *IRTServiceImpl) EstimateAbility(currentAbility, difficulty, discrimination, guessParameter float64, isCorrect bool) float64 {
	const D = 1.7

	// 使用最大似然估计（MLE）方法更新能力值
	var newAbility float64 = currentAbility
	var lastAbility float64

	for i := 0; i < s.maxIterations; i++ {
		lastAbility = newAbility

		// 计算作答正确的概率
		p := s.calculateProbability(newAbility, difficulty, discrimination, guessParameter)

		// 计算信息函数
		info := s.calculateInformation(newAbility, difficulty, discrimination, guessParameter, p)

		// 计算得分函数
		score := s.calculateScore(isCorrect, p)

		// 更新能力值估计
		if info != 0 {
			newAbility = lastAbility + score/info
		}

		// 检查是否收敛
		if math.Abs(newAbility-lastAbility) < s.convergence {
			break
		}
	}

	// 限制能力值范围在 [-3, 3] 之间
	if newAbility < -3 {
		newAbility = -3
	} else if newAbility > 3 {
		newAbility = 3
	}

	return newAbility
}

// GetNextQuestionDifficulty 根据当前能力值确定下一题的难度
func (s *IRTServiceImpl) GetNextQuestionDifficulty(currentAbility float64) float64 {
	// 在自适应测试中，通常选择难度接近当前能力值的题目
	return currentAbility
}

// calculateProbability 计算作答正确的概率
func (s *IRTServiceImpl) calculateProbability(ability, difficulty, discrimination, guessParameter float64) float64 {
	const D = 1.7
	z := D * discrimination * (ability - difficulty)
	return guessParameter + (1-guessParameter)/(1+math.Exp(-z))
}

// calculateInformation 计算信息函数
func (s *IRTServiceImpl) calculateInformation(ability, difficulty, discrimination, guessParameter, probability float64) float64 {
	const D = 1.7
	q := 1 - probability
	numerator := math.Pow(D*discrimination*(probability-guessParameter), 2)
	denominator := (probability - guessParameter) * q
	return numerator / denominator
}

// calculateScore 计算得分函数
func (s *IRTServiceImpl) calculateScore(isCorrect bool, probability float64) float64 {
	if isCorrect {
		return 1 - probability
	}
	return -probability
}
