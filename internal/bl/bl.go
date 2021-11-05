package bl

import (
	"errors"
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

var (
	users    []user
	sessions []session
)

func SingUp(login, pass string) (sessionID int, err error) {
	if len(login) < 3 {
		return 0, errors.New("too short login")
	}
	if len(pass) < 6 {
		return 0, errors.New("too short pass")
	}

	newUser := user{
		id(len(users) + 1),
		login,
		pass,
	}
	users = append(users, newUser)

	newSession := session{
		id(len(sessions) + 1),
		newUser.id,
		false,
	}
	sessions = append(sessions, newSession)

	return int(newSession.id), nil
}

func LogIn(login, pass string) (sessionID int, err error) {
	userID := id(-1)
	for _, user := range users {
		if user.login == login && user.pass == pass {
			userID = user.id
			break
		}
	}
	if userID == id(-1) {
		return -1, errors.New("undefined user")
	}

	newSession := session{
		id(len(sessions) + 1),
		userID,
		false,
	}
	sessions = append(sessions, newSession)

	return int(newSession.id), nil
}

func LogOut(sessionID int) (err error) {
	for idx, session := range sessions {
		if session.id == id(sessionID) {
			sessions[idx].deleted = true
			return nil
		}
	}

	return errors.New("undefined session")
}
