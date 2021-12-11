package domain

import "math/rand"

type Agent struct {
	ID           int                `json:"id"`
	FitnessValue float64            `json:"fitness_value"`
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
			change := rand.Float64() * 2 * power
			a.Genocode[k] = v * change
			if k == "rho" && a.Genocode[k] > 0.9999 {
				a.Genocode[k] = 0.9999
			}
			if k == "quantity" && a.Genocode[k] < 1.0 {
				a.Genocode[k] = 1.0
			}
			break
		}
		index = index + 1
	}
}
