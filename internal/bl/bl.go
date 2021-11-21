package bl

import (
	"errors"
	"github.com/ValeryBMSTU/evoModeler/internal/domain"
)

type DA interface {
	InsertUser(login, pass string) (userID int, err error)
	InsertSession(userID int) (sessionID int, err error)
	DeleteSession(sessionID int) (err error)
	SelectUser(login, pass string) (userID int, err error)
	SelectUserByID(userID int) (user domain.User, err error)
	SelectSession(sessionID int) (id int, idUser int, isDeleted bool, err error)
	SelectSolver(solverName string) (solver domain.Solver, err error)
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

func (bl *Bl) CreateTask(taskName, solverName, genAlgName string, user domain.User) (task domain.Task, err error) {
	_, err = bl.Da.SelectSolver(solverName)
	if err != nil {
		return task, err
	}



	return task, nil
}