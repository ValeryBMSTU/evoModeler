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
