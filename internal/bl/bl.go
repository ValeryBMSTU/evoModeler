package bl

import (
	"errors"
	"github.com/ValeryBMSTU/evoModeler/internal/domain"
	"math/rand"
	"time"
)

type DA interface {
	InsertUser(login, pass string) (userID int, err error)
	InsertSession(userID int) (sessionID int, err error)
	InsertTask(task domain.Task) (taskID int, err error)
	UpdateTaskStatus(taskID int, status string) (err error)
	DeleteSession(sessionID int) (err error)
	SelectUser(login, pass string) (userID int, err error)
	SelectUserByID(userID int) (user domain.User, err error)
	SelectSession(sessionID int) (id int, idUser int, isDeleted bool, err error)
	SelectSolver(solverName string) (solver domain.Solver, err error)
	SelectGenAlg(genAlgName string) (genAlg domain.GenAlg, err error)
	SelectSolvers() (solvers []domain.Solver, err error)
	SelectIssues() (issues []domain.Issue, err error)
}

type Bl struct {
	Da DA
}

func CreateBl(da DA) (bl *Bl, err error) {
	return &Bl{da}, nil
}

func (bl *Bl) CreateUser(login, pass string) (sessionID int, err error) {
	if len(login) < 3 {
		return -1, errors.New("too short login")
	}
	if len(pass) < 6 {
		return -1, errors.New("too short pass")
	}

	userID, err := bl.Da.InsertUser(login, pass)
	if err != nil {
		return -1, err
	}

	sessionID, err = bl.Da.InsertSession(userID)
	if err != nil {
		return -1, err
	}

	return sessionID, nil
}

func (bl *Bl) CreateSession(login, pass string) (sessionID int, err error) {

	userID, err := bl.Da.SelectUser(login, pass)
	if err != nil {
		return -1, err
	}

	sessionID, err = bl.Da.InsertSession(userID)
	if err != nil {
		return -1, err
	}

	return sessionID, nil
}

func (bl *Bl) RemoveSession(sessionID int) (err error) {
	err = bl.Da.DeleteSession(sessionID)
	if err != nil {
		return err
	}

	return nil
}

func (bl *Bl) CheckSession(sessionID int) (isExist bool, err error) {
	_, _, _, err = bl.Da.SelectSession(sessionID)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (bl *Bl) TakeSession(sessionID int) (session domain.Session, err error) {
	_, userID, isDeleted, err := bl.Da.SelectSession(sessionID)
	if err != nil {
		return session, err
	}

	session.ID = sessionID
	session.UserID = userID
	session.Deleted = isDeleted

	return session, nil
}

func (bl *Bl) TakeUser(userID int) (user domain.User, err error) {
	user, err = bl.Da.SelectUserByID(userID)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (bl *Bl) TakeSolver(solverName string) (solver domain.Solver, err error) {
	solver, err = bl.Da.SelectSolver(solverName)
	if err != nil {
		return solver, err
	}

	return solver, err
}

func (bl *Bl) TakeSolvers() (solvers []domain.Solver, err error) {
	solvers, err = bl.Da.SelectSolvers()
	if err != nil {
		return solvers, err
	}

	return solvers, err
}

func (bl *Bl) TakeIssues() (issues []domain.Issue, err error) {
	issues, err = bl.Da.SelectIssues()
	if err != nil {
		return issues, err
	}

	return issues, nil
}

func (bl *Bl) CreateTask(taskName, solverName, genAlgName string, user domain.User) (result domain.Result, err error) {
	solver, err := bl.Da.SelectSolver(solverName)
	if err != nil {
		return result, err
	}

	genAlg, err := bl.Da.SelectGenAlg(genAlgName)
	if err != nil {
		return result, err
	}

	task := domain.Task{
		ID:          -1,
		Name:        taskName,
		CreateDate:  time.Now().String(),
		Description: "default desc",
		Status:      "init",
		UserID:      user.ID,
		GenAlgID:    genAlg.ID,
		SolverID:    solver.GetID(),
		Solver:      solver,
		GenAlg:      genAlg,
	}

	taskID, err := bl.Da.InsertTask(task)
	if err != nil {
		return result, err
	}

	task.ID = taskID

	result, err = bl.RunTask(task)
	if err != nil {
		return result, nil
	}

	return result, nil
}

func (bl *Bl) RunTask(task domain.Task) (result domain.Result, err error) {
	err = bl.Da.UpdateTaskStatus(task.ID, "running")
	if err != nil {
		return result, err
	}

	//taskModel := [][]int{{0, 2, 30, 9, 1},
	//	{4, 0, 47, 7, 7},
	//	{31, 33, 0, 33, 36},
	//	{20, 13, 16, 0, 28},
	//	{9, 36, 22, 22, 0}}

	//taskModel := [][]int{{0, 48, 16, 11, 28, 42, 49, 1, 24, 40, 14, 43, 12, 32, 39, 6, 42, 11, 39, 9},
	//	{29, 0, 42, 19, 29, 41, 34, 29, 24, 2, 4, 15, 10, 17, 20, 37, 21, 15, 1, 19},
	//	{29, 40, 0, 23, 29, 6, 18, 37, 36, 27, 2, 40, 42, 46, 22, 48, 41, 44, 15, 48},
	//	{8, 23, 36, 0, 45, 8, 42, 49, 8, 25, 48, 2, 35, 9, 32, 46, 26, 12, 27, 31},
	//	{32, 5, 26, 9, 0, 13, 16, 11, 34, 2, 10, 6, 6, 39, 45, 4, 19, 19, 26, 41},
	//	{43, 6, 12, 18, 37, 0, 35, 43, 10, 5, 31, 22, 45, 49, 13, 38, 49, 35, 11, 14},
	//	{39, 13, 29, 10, 19, 5, 0, 38, 44, 48, 18, 48, 49, 34, 26, 45, 11, 31, 33, 12},
	//	{30, 29, 48, 30, 1, 12, 2, 0, 19, 28, 13, 38, 13, 4, 25, 5, 49, 38, 22, 42},
	//	{5, 48, 9, 44, 34, 13, 9, 9, 0, 46, 1, 36, 4, 19, 20, 26, 25, 38, 48, 47},
	//	{49, 21, 8, 48, 9, 34, 2, 20, 9, 0, 33, 38, 12, 9, 7, 46, 34, 16, 41, 1},
	//	{40, 2, 20, 14, 35, 8, 6, 20, 33, 34, 0, 23, 8, 15, 14, 31, 28, 7, 8, 46},
	//	{2, 8, 48, 40, 31, 31, 36, 24, 11, 45, 10, 0, 49, 40, 2, 27, 49, 42, 2, 4},
	//	{40, 40, 12, 1, 11, 6, 46, 36, 36, 45, 47, 23, 0, 49, 4, 31, 9, 7, 23, 11},
	//	{35, 18, 18, 7, 16, 37, 46, 36, 34, 45, 4, 12, 6, 0, 15, 30, 39, 42, 5, 47},
	//	{46, 46, 38, 43, 40, 49, 37, 31, 16, 18, 15, 9, 27, 49, 0, 13, 27, 22, 19, 27},
	//	{13, 44, 18, 44, 40, 23, 25, 9, 23, 13, 32, 22, 7, 28, 28, 0, 21, 26, 42, 2},
	//	{3, 10, 46, 6, 23, 35, 47, 6, 40, 7, 25, 37, 26, 6, 6, 15, 0, 34, 15, 28},
	//	{25, 9, 28, 12, 31, 46, 47, 45, 13, 46, 35, 35, 36, 20, 28, 13, 37, 0, 16, 24},
	//	{31, 29, 32, 6, 29, 2, 6, 23, 14, 45, 43, 33, 32, 7, 22, 31, 35, 28, 0, 48},
	//	{19, 42, 46, 15, 19, 19, 14, 32, 32, 19, 35, 39, 44, 7, 48, 20, 14, 14, 28, 0}}

	nn := 42
	taskModel := make([][]int, nn, nn)
	for i, _ := range taskModel {
		taskModel[i] = make([]int, nn, nn)
		for j, _ := range taskModel[i] {
			if i == j {
				taskModel[i][j] = 0
			} else {
				taskModel[i][j] = rand.Intn(99) + 1
			}
		}
	}

	task.Solver.Set(taskModel)

	result, err = bl.RunGenAlg(task.GenAlg, task.Solver, task.ID)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (bl *Bl) RunGenAlg(genAlg domain.GenAlg, solver domain.Solver, taskID int) (result domain.Result, err error) {
	generation, err := genAlg.InitGeneration(taskID, solver.GetBaseParams())
	if err != nil {
		return result, err
	}

	result = domain.Result{
		BestScores: make([]float64, 0, 0),
		BestParams: make(map[string][]float64),
		AvgScores:  make([]float64, 0, 0),
		AvgParams:  make(map[string][]float64),
	}

	agesCount := 30
	var bestScore, avgScore float64
	var bestParams, avgParams map[string]float64

	for i := 0; i < agesCount; i++ {
		generation, err = genAlg.Selection(generation)
		if err != nil {
			return result, err
		}
		generation, err = genAlg.Reproduction(generation, i)
		if err != nil {
			return result, err
		}
		generation, bestScore, bestParams, avgScore, avgParams, err = genAlg.CalculateFitness(generation, solver)
		if err != nil {
			return result, err
		}

		result.BestScores = append(result.BestScores, bestScore)
		for k, v := range bestParams {
			result.BestParams[k] = append(result.BestParams[k], v)
		}
		result.AvgScores = append(result.AvgScores, avgScore)
		for k, v := range avgParams {
			result.AvgParams[k] = append(result.AvgParams[k], v)
		}
	}

	return result, nil
}
