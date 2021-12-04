package domain

import "errors"

type Solver interface {
	GetID() (id int)
	Set(model interface{}) (err error)
	Solve(params ...map[string]float64) (score float64, err error)
	GetBaseParams() (params map[string]float64)
}

type AntSolver struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Model       [][]int `json:"model"`
	IssueID     int     `json:"issue_id"`
	Alpha       float64 `json:"alpha"`
	Beta        float64 `json:"beta"`
	Rho         float64 `json:"rho"`
	Quantity    float64 `json:"quantity"`
}

func (s *AntSolver) GetID() (id int) {
	return s.ID
}

func (s *AntSolver) GetBaseParams() (params map[string]float64) {
	return map[string]float64{
		"alpha":    s.Alpha,
		"beta":     s.Beta,
		"rho":      s.Rho,
		"quantity": s.Quantity,
	}
}

func (s *AntSolver) Set(model interface{}) (err error) {
	matrix, ok := model.([][]int)
	if !ok {
		return errors.New("fail casting model to matrix")

	}

	s.Model = matrix

	return nil
}

func (s *AntSolver) Solve(params ...map[string]float64) (score float64, err error) {

	return 0, nil
}
