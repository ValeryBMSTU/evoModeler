package domain

type Generation struct {
	ID          int     `json:"id"`
	OrderNumber int     `json:"order_number"`
	TaskID      int     `json:"task_id"`
	Agents      []Agent `json:"agents"`
}
