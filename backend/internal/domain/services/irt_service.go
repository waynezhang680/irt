package services

type IRTService interface {
	// 估计考生能力值
	EstimateAbility(currentAbility, difficulty, discrimination, guessParameter float64, isCorrect bool) float64
	// 获取下一题的建议难度
	GetNextQuestionDifficulty(currentAbility float64) float64
}
