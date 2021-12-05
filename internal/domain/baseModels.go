package domain

type User struct {
	ID    int    `json:"id"`
	Login string `json:"login"`
	Pass  string `json:"pass"`
}

type Session struct {
	ID      int  `json:"id"`
	UserID  int  `json:"user_id"`
	Deleted bool `json:"deleted"`
}

type Issue struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Result struct {
	BestScores []float64            `json:"best_scores"`
	BestParams map[string][]float64 `json:"best_params"`
	AvgScores  []float64            `json:"avg_scores"`
	AvgParams  map[string][]float64 `json:"avg_params"`
}
