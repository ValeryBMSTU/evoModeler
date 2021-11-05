package da

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host           = "127.0.0.1"
	port           = "5432"
	user           = "postgres"
	password       = "postgres"
	dbname         = "postgres"
	inserUserQuery = `insert into "User" (login, pass) values ($1, $2)`
)

var connStr string = "user=" + user + " password=" + password + " dbname=" + dbname + " sslmode=disable"

func SignUp(login, pass string) error {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("can not connect to database")
		return err
	}
	defer db.Close()

	result, err := db.Exec(inserUserQuery, login, pass)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(result.LastInsertId())

	return nil
}
