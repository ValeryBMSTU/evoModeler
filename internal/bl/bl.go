package bl

import (
	"errors"
)

type id int

type sessionToken int

type DA interface {
	InsertUser(login, pass string) (userID int, err error)
	InsertSession(userID int) (sessionID int, err error)
	DeleteSession(sessionID int) (err error)
	SelectUser(login, pass string) (userID int, err error)
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
