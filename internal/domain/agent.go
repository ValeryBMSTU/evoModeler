package domain

import "math/rand"

type Agent struct {
	ID           int                `json:"id"`
	FitnessValue int                `json:"fitness_value"`
	ParentID     int                `json:"parent_id"`
	GenerationID int                `json:"generation_id"`
	CodeID       int                `json:"code_id"`
	Genocode     map[string]float64 `json:"genocode"`
}

func (a *Agent) Mutate(power float64) {
	mutatedParamIndex := rand.Int() % len(a.Genocode)
	index := 0
	for k, v := range a.Genocode {
		if index == mutatedParamIndex {
			cleanChange := rand.Float64() * 2 * power
			vChange := v * (power - cleanChange)
			a.Genocode[k] = v + vChange
			break
		}
		index = index + 1
	}
}
