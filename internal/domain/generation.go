package domain

type Generation struct {
	ID          int     `json:"id"`
	OrderNumber int     `json:"order_number"`
	TaskID      int     `json:"task_id"`
	Agents      []Agent `json:"agents"`
}

func (g *Generation) Len() int {
	return len(g.Agents)
}

func (g *Generation) Swap(i, j int) {
	g.Agents[i], g.Agents[j] = g.Agents[j], g.Agents[i]
}

func (g *Generation) Less(i, j int) bool {
	return g.Agents[i].FitnessValue < g.Agents[j].FitnessValue
}
