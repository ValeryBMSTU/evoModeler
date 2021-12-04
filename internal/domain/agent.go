package domain

type Agent struct {
	ID           int                `json:"id"`
	FitnessValue float32            `json:"fitness_value"`
	ParentID     int                `json:"parent_id"`
	GenerationID int                `json:"generation_id"`
	CodeID       int                `json:"code_id"`
	Genocode     map[string]float64 `json:"genocode"`
}
