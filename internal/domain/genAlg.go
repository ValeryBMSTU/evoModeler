package domain

import (
	"math/rand"
	"sort"
)

type GenAlg struct {
	ID             int         `json:"id"`
	Name           string      `json:"name"`
	Description    string      `json:"description"`
	Config         interface{} `json:"config"`
	PopSize        int         `json:"pop_size"`
	MutationChance float64     `json:"mutation_chance"`
	MutationPower  float64     `json:"mutation_power"`
	DropPart       float64     `json:"drop_part"`
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

func (ga *GenAlg) CalculateFitness(generation Generation, solver Solver) (newGeneration Generation, err error) {
	for index, agent := range generation.Agents {
		agent.FitnessValue, err = solver.Solve(agent.Genocode)
		if err != nil {
			return generation, err
		}
		generation.Agents[index] = agent
	}
	return generation, nil
}

func (ga *GenAlg) Selection(generation Generation) (newGeneration Generation, err error) {
	generation, err = ga.SortGeneration(generation)
	if err != nil {
		return generation, err
	}

	generation.Agents = generation.Agents[:int(float64(ga.PopSize)*(1.0-ga.DropPart))]

	return generation, nil
}

func (ga *GenAlg) SortGeneration(generation Generation) (newGeneration Generation, err error) {
	sort.Sort(&generation)
	return generation, nil
}

func (ga *GenAlg) Reproduction(generation Generation) (newGeneration Generation, err error) {
	newGeneration = generation
	for len(newGeneration.Agents) < ga.PopSize {
		newAgent := generation.Agents[rand.Int()%len(generation.Agents)]
		if r := rand.Float64(); r < ga.MutationChance {
			newAgent.Mutate(ga.MutationPower)
		}
		newGeneration.Agents = append(newGeneration.Agents, newAgent)
	}

	return newGeneration, nil
}
