package bl

import (
	"errors"

	"github.com/ValeryBMSTU/evoModeler/internal/da"
)

type id int

type sessionToken int

type user struct {
	id    id
	login string
	pass  string
}

type session struct {
	id      id
	userID  id
	deleted bool
}

func CreateUser(login, pass string) (sessionID int, err error) {
	if len(login) < 3 {
		return -1, errors.New("too short login")
	}
	if len(pass) < 6 {
		return -1, errors.New("too short pass")
	}

	userID, err := da.InsertUser(login, pass)
	if err != nil {
		return -1, err
	}

	sessionID, err = da.InsertSession(userID)
	if err != nil {
		return -1, err
	}

	return sessionID, nil
}

func CreateSession(login, pass string) (sessionID int, err error) {

	userID, err := da.SelectUser(login, pass)
	if err != nil {
		return -1, err
	}

	sessionID, err = da.InsertSession(userID)
	if err != nil {
		return -1, err
	}

	return sessionID, nil
}

func RemoveSession(sessionID int) (err error) {
	err = da.DeleteSession(sessionID)
	if err != nil {
		return err
	}

	return nil
}
