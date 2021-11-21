package domain

type User struct {
	ID    int		`json:"id"`
	Login string	`json:"login"`
	Pass  string	`json:"pass"`
}

type Session struct {
	ID      int		`json:"id"`
	UserID  int		`json:"user_id"`
	Deleted bool 	`json:"deleted"`
}

type Issue struct {
	ID 			int			`json:"id"`
	Name 		string		`json:"name"`
	Description string		`json:"description"`
}

type Solver struct {
	ID 			int			`json:"id"`
	Name 		string		`json:"name"`
	Description string		`json:"description"`
	Model		interface{}	`json:"model"`
	IssueID		int 		`json:"issue_id"`
}

type GenAlg struct {
	ID			int			`json:"id"`
	Config		interface{}	`json:"config"`
}

type Task struct {
	ID			int			`json:"id"`
	SolverID  	int			`json:"task_type_id"`
	GenAlgID	int			`json:"gen_alg_id"`
	Name		string		`json:"name"`
	Config   	interface{}	`json:"config"`
	Solver 		Solver		`json:"solver"`
	GenAlg   	GenAlg		`json:"gen_alg"`
}
