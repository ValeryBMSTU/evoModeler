package domain

import (
	"errors"
	"github.com/ValeryBMSTU/evoModeler/pkg"
	"math/rand"
)

type Solver interface {
	GetID() (id int)
	Set(model interface{}) (err error)
	Solve(params map[string]float64) (score float64, err error)
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

func (s *AntSolver) Solve(params map[string]float64) (score float64, err error) {
	L := s.Model
	cities := len(L[0])
	ages := 10
	ants := 5

	alpha := params["alpha"]
	beta := params["beta"]
	rho := params["rho"]
	quantity := params["quantity"]
	e := 1
	ph := float64(quantity) / float64(cities)

	revMatrix := s.CalcRevMatrix(L)
	tao := s.CalcTaoMatrix(L, ph)

	bestDist := 999999
	bestRoute := make([]int, 0, 0)
	antRoute := pkg.GetZeroIntMatrix(ants, cities)
	antDist := make([]int, ants, ants)
	antBestDist := make([]int, ages, ages)
	antAvgDist := make([]int, ages, ages)
	//bestAge := 999999

	for age := 0; age < ages; age++ {
		antRoute = pkg.GetZeroIntMatrix(ants, cities)
		antDist = make([]int, ants, ants)

		for ant := 0; ant < ants; ant++ {
			antRoute[ant][0] = ant % cities

			for city := 1; city < cities; city++ {
				fromCity := antRoute[ant][city-1]

				P := pkg.MultArrays(pkg.PowArray(tao[fromCity], alpha), pkg.PowArray(revMatrix[fromCity], beta))
				for i := 0; i < city; i++ {
					P[antRoute[ant][i]] = 0
				}
				P = pkg.DivideArray(P, pkg.ArraySum(P))

				isNotChosen := true
				for isNotChosen {
					r := rand.Float64()
					for c := 0; c < cities; c++ {
						if P[c] >= r {
							antRoute[ant][city] = c
							isNotChosen = false
							break
						}
					}
				}
			}

			for city := 0; city < cities; city++ {
				cityFrom := -1
				if city == 0 {
					cityFrom = antRoute[ant][cities-1]
				} else {
					cityFrom = antRoute[ant][city-1]
				}
				cityTo := antRoute[ant][city]
				antDist[ant] += L[cityFrom][cityTo]
			}

			for antDist[ant] < bestDist {
				bestDist = antDist[ant]
				bestRoute = antRoute[ant]
				//bestAge = age
			}

		}

		tao = pkg.MultMatrix(tao, 1-rho)

		for ant := 0; ant < ants; ant++ {
			for city := 0; city < cities; city++ {
				cityFrom := -1
				if city == 0 {
					cityFrom = antRoute[ant][cities-1]
				} else {
					cityFrom = antRoute[ant][city-1]
				}
				cityTo := antRoute[ant][city]
				tao[cityFrom][cityTo] = tao[cityFrom][cityTo] + (float64(quantity) / float64(antDist[ant]))
				tao[cityTo][cityFrom] = tao[cityFrom][cityTo]
			}
		}

		for city := 0; city < cities; city++ {
			cityFrom := -1
			if city == 0 {
				cityFrom = bestRoute[cities-1]
			} else {
				cityFrom = bestRoute[city-1]
			}
			cityTo := bestRoute[city]
			tao[cityFrom][cityTo] = tao[cityFrom][cityTo] + (float64(e) * float64(quantity) / float64(bestDist))
			tao[cityTo][cityFrom] = tao[cityFrom][cityTo]
		}

		antBestDist[age] = bestDist
		antAvgDist[age] = pkg.ArrayAvg(antDist)
	}

	//fmt.Println(bestDist)

	return float64(bestDist), nil
}

func (s *AntSolver) CalcRevMatrix(matrix [][]int) (revMatrix [][]float64) {
	revMatrix = make([][]float64, len(matrix))
	for i := 0; i < len(matrix); i++ {
		revMatrix[i] = make([]float64, len(matrix))
	}

	for i, arr := range matrix {
		for j, v := range arr {
			if i != j {
				revMatrix[i][j] = 1.0 / float64(v)
			}
		}
	}

	return revMatrix
}

func (s *AntSolver) CalcTaoMatrix(matrix [][]int, ph float64) (newMatrix [][]float64) {
	taoMatrix := make([][]float64, len(matrix))
	for i := 0; i < len(matrix); i++ {
		taoMatrix[i] = make([]float64, len(matrix))
	}

	for i, arr := range matrix {
		for j, _ := range arr {
			taoMatrix[i][j] = 1.0 * ph
		}
	}

	return taoMatrix
}
