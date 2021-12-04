package domain

type Task struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	CreateDate  string `json:"create_date"`
	Description string `json:"description"`
	Status      string `json:"status"`
	UserID      int    `json:"user_id"`
	GenAlgID    int    `json:"gen_alg_id"`
	SolverID    int    `json:"solver_id"`
	Solver      Solver `json:"solver"`
	GenAlg      GenAlg `json:"gen_alg"`
}
