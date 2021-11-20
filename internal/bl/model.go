package bl

type user struct {
	id    id
	login string
	pass  string
}
// 123
type session struct {
	id      id
	userID  id
	deleted bool `json:"deleted"`
}
