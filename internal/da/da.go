package da

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "127.0.0.1"
	port     = "5432"
	user     = "postgres"
	password = "postgres"
	dbname   = "postgres"

	insertUserQuery    = `insert into "User" (login, pass) values ($1, $2) returning id`
	insertSessionQuery = `insert into "Session" (id_user, is_deleted) values ($1, $2) returning id`

	selectUserQuery = `SELECT id, login, pass FROM "User" WHERE login=$1, pass=$2`

	deleteSessionQuery = `UPDATE "Session" SET is_deleted=true WHERE id = $1`
)

var connStr string = "user=" + user + " password=" + password + " dbname=" + dbname + " sslmode=disable"

func InsertUser(login, pass string) (userID int, err error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("can not connect to database")
		return -1, err
	}
	defer db.Close()

	var lastInserID int
	err = db.QueryRow(insertUserQuery, login, pass).Scan(&lastInserID)
	if err != nil {
		fmt.Println(err)
		return -1, err
	}
	fmt.Println(lastInserID)

	return lastInserID, err
}

func InsertSession(userID int) (seesionID int, err error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("can not connect to database")
		return -1, err
	}
	defer db.Close()

	fmt.Println(userID)
	var lastInserID int
	err = db.QueryRow(insertSessionQuery, userID, false).Scan(&lastInserID)
	if err != nil {
		fmt.Println(err)
		return -1, err
	}
	fmt.Println(lastInserID)

	return lastInserID, nil
}

func DeleteSession(sessionID int) error {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("can not connect to database")
		return err
	}
	defer db.Close()

	_, err = db.Exec(deleteSessionQuery, sessionID)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func SelectUser(login, pass string) (userID int, err error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("can not connect to database")
		return -1, err
	}
	defer db.Close()

	err = db.QueryRow(selectUserQuery, userID, false).Scan(&userID)
	if err != nil {
		fmt.Println(err)
		return -1, err
	}
	fmt.Println(userID)

	return userID, nil
}
