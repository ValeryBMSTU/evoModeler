package domain

import (
	"fmt"
	"github.com/ValeryBMSTU/evoModeler/pkg"
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

func (ga *GenAlg) CalculateFitness(generation Generation, solver Solver) (newGeneration Generation,
	bestScore float64, bestParams map[string]float64, avgScore float64, avgParams map[string]float64, err error) {

	bestScore = 999999.0
	scores := make([]float64, 0, 0)
	bestParams = make(map[string]float64)
	avgParams = make(map[string]float64)
	params := make(map[string][]float64)

	for index, agent := range generation.Agents {
		agent.FitnessValue, err = solver.Solve(agent.Genocode)
		if err != nil {
			return generation, bestScore, bestParams, avgScore, avgParams, err
		}

		scores = append(scores, agent.FitnessValue)
		for k, v := range agent.Genocode {
			params[k] = append(params[k], v)
		}

		if agent.FitnessValue < bestScore {
			bestScore = agent.FitnessValue
			for k, v := range agent.Genocode {
				bestParams[k] = v
			}
		}

		fmt.Println("Age: ", generation.OrderNumber, "agent ", index, ": ", agent.FitnessValue)
		generation.Agents[index] = agent
	}

	avgScore = pkg.ArrayFloat64Avg(scores)
	for k, _ := range generation.Agents[0].Genocode {
		avgParams[k] = pkg.ArrayFloat64Avg(params[k])
	}

	return generation, bestScore, bestParams, avgScore, avgParams, nil
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

func (ga *GenAlg) Reproduction(generation Generation, age int) (newGeneration Generation, err error) {
	newGeneration = generation
	newGeneration.OrderNumber = age
	for len(newGeneration.Agents) < ga.PopSize {

		oldAgentPos := rand.Int() % (len(generation.Agents) / 2)
		newGenocode := ga.CopyGenocode(generation.Agents[oldAgentPos].Genocode)
		newAgent := Agent{
			ID:           rand.Int(),
			FitnessValue: 0,
			ParentID:     generation.Agents[oldAgentPos].ID,
			GenerationID: generation.ID,
			CodeID:       -1,
			Genocode:     newGenocode,
		}

		if r := rand.Float64(); r < ga.MutationChance {
			newAgent.Mutate(ga.MutationPower)
		}

		newGeneration.Agents = append(newGeneration.Agents, newAgent)
	}

	ga.ShuffleAgents(newGeneration.Agents)

	return newGeneration, nil
}

func (ga *GenAlg) CopyGenocode(oldCode map[string]float64) (newCode map[string]float64) {
	newCode = make(map[string]float64)
	for k, v := range oldCode {
		newCode[k] = v
	}
	return newCode
}

func (ga *GenAlg) ShuffleAgents(agents []Agent) {
	for i := range agents {
		j := rand.Intn(i + 1)
		agents[i], agents[j] = agents[j], agents[i]
	}
}
