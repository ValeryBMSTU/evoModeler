package domain

type GenAlg struct {
	ID             int         `json:"id"`
	Name           string      `json:"name"`
	Description    string      `json:"description"`
	Config         interface{} `json:"config"`
	PopSize        int         `json:"pop_size"`
	MutationChance float64     `json:"mutation_chance"`
	MutationPower  float64     `json:"mutation_power"`
}

func (ga *GenAlg) InitGeneration(taskID int, params map[string]float64) (generation Generation, err error) {

	agents := make([]Agent, 0, ga.PopSize)
	for i := 0; i < ga.PopSize; i++ {
		agent, err := ga.CreateAgent(0, params)
		if err != nil {
			return generation, nil
		}
		agents = append(agents, agent)
	}

	generation = Generation{
		ID:          -1,
		OrderNumber: 0,
		TaskID:      taskID,
		Agents:      agents,
	}

	return generation, nil
}

func (ga *GenAlg) CreateAgent(generationID int, genocode map[string]float64) (agent Agent, err error) {
	agent = Agent{
		ID:           -1,
		FitnessValue: -1,
		ParentID:     -1,
		GenerationID: generationID,
		CodeID:       0,
		Genocode:     genocode,
	}

	return agent, nil
}
