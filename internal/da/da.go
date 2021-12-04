package da

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/ValeryBMSTU/evoModeler/internal/domain"

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
	insertTaskQuery    = `INSERT INTO "Task" (name, create_date, description, status, id_user, id_GA, id_solver)
						  VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`

	updateTaskStatusQuery = `UPDATE "Task" SET status = $1 WHERE id = $2`

	selectUserQuery               = `SELECT id, login, pass FROM "User" WHERE login=$1 and pass=$2`
	selectUserByIDQuery           = `SELECT id, login, pass FROM "User" WHERE id=$1`
	selectSessionByIDQuery        = `SELECT id, id_user, is_deleted FROM "Session" WHERE id=$1`
	selectSolversQuery            = `SELECT id, name, description, id_issue FROM "Solver"`
	selectSolverBySolverNameQuery = `SELECT id, name, description, id_issue FROM "Solver" WHERE name=$1`
	selectGenAlgByGenAlgNameQuery = `SELECT id, name, description, config FROM "GeneticAlgorithm WHERE name=$1"`
	selectIssuesQuery             = `SELECT id, name, description FROM "Issue"`

	deleteSessionQuery = `UPDATE "Session" SET is_deleted=true WHERE id = $1`
)

var connStr string = "user=" + user + " password=" + password + " dbname=" + dbname + " sslmode=disable"

type Da struct {
	connStr string
}

func CreateDa() (da *Da, err error) {
	return &Da{connStr: connStr}, nil
}

func (da *Da) InsertUser(login, pass string) (userID int, err error) {
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

	return lastInserID, err
}

func (da *Da) InsertSession(userID int) (sessionID int, err error) {
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

	return lastInserID, nil
}

func (da *Da) InsertTask(task domain.Task) (taskID int, err error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("can not connect to database")
		return -1, err
	}
	defer db.Close()

	fmt.Println(taskID)
	var lastInsertID int
	err = db.QueryRow(insertTaskQuery,
		task.Name,
		task.CreateDate,
		task.Description,
		task.Status,
		task.UserID,
		task.GenAlgID,
		task.SolverID,
		false).Scan(&lastInsertID)
	if err != nil {
		fmt.Println(err)
		return -1, err
	}

	return lastInsertID, err
}

func (da *Da) UpdateTaskStatus(taskID int, status string) (err error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("can not connect to database")
		return err
	}
	defer db.Close()

	fmt.Println(taskID)
	var lastInsertID int
	err = db.QueryRow(updateTaskStatusQuery,
		status,
		taskID,
		false).Scan(&lastInsertID)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return err
}

func (da *Da) DeleteSession(sessionID int) (err error) {
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

func (da *Da) SelectUser(login, pass string) (userID int, err error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("can not connect to database")
		return -1, err
	}
	defer db.Close()

	var uLogin, uPass string
	err = db.QueryRow(selectUserQuery, login, pass).Scan(&userID, &uLogin, &uPass)
	if err != nil {
		fmt.Println(err)
		return -1, err
	}

	return userID, nil
}

func (da *Da) SelectUserByID(userID int) (user domain.User, err error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("can not connect to database")
		return user, err
	}
	defer db.Close()

	err = db.QueryRow(selectUserByIDQuery, userID).Scan(&user.ID, &user.Login, &user.Pass)
	if err != nil {
		fmt.Println(err)
		return user, err
	}

	return user, nil
}

func (da *Da) SelectSession(sessionID int) (id int, idUser int, isDeleted bool, err error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("can not connect to database")
		return -1, -1, false, err
	}
	defer db.Close()

	err = db.QueryRow(selectSessionByIDQuery, sessionID).Scan(&id, &idUser, &isDeleted)
	if err != nil {
		fmt.Println(err)
		return -1, -1, false, err
	}

	return id, idUser, isDeleted, nil
}

func (da *Da) SelectSolver(solverName string) (solver domain.Solver, err error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("can not connect to database")
		return solver, err
	}
	defer db.Close()

	var ID, issueID int
	var name, desc string
	err = db.QueryRow(selectSolverBySolverNameQuery, solverName).Scan(
		&ID,
		&name,
		&desc,
		&issueID)
	if err != nil {
		return solver, err
	}

	if name == "ant" {
		solver = &domain.AntSolver{
			ID:          ID,
			Name:        name,
			Description: desc,
			Model:       nil,
			IssueID:     issueID,
			Alpha:       0,
			Beta:        0,
			Rho:         0,
			Quantity:    0,
		}
	} else {
		return solver, errors.New("unknown solver")
	}

	return solver, nil
}

func (da *Da) SelectGenAlg(genAlgName string) (genAlg domain.GenAlg, err error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("can not connect to database")
		return genAlg, err
	}
	defer db.Close()

	err = db.QueryRow(selectGenAlgByGenAlgNameQuery, genAlgName).Scan(
		&genAlg.ID,
		&genAlg.Name,
		&genAlg.Description,
		&genAlg.Config)
	if err != nil {
		return genAlg, err
	}

	genAlg.PopSize = 10
	genAlg.MutationChance = 0.1
	genAlg.MutationPower = 0.1

	return genAlg, nil
}

func (da *Da) SelectIssues() (issues []domain.Issue, err error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("can not connect to database")
		return issues, err
	}
	defer db.Close()

	rows, err := db.Query(selectIssuesQuery)
	if err != nil {
		return issues, err
	}
	defer rows.Close()

	for rows.Next() {
		var issue domain.Issue
		if err := rows.Scan(
			&issue.ID,
			&issue.Name,
			&issue.Description); err != nil {
			return issues, err
		}
		issues = append(issues, issue)
	}
	if err = rows.Err(); err != nil {
		return issues, err
	}

	return issues, nil
}

func (da *Da) SelectSolvers() (solvers []domain.Solver, err error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("can not connect to database")
		return solvers, err
	}
	defer db.Close()

	rows, err := db.Query(selectSolversQuery)
	if err != nil {
		return solvers, err
	}
	defer rows.Close()

	for rows.Next() {
		var solver domain.Solver
		var ID, issueID int
		var name, desc string
		if err := rows.Scan(
			&ID,
			&name,
			&desc,
			&issueID); err != nil {
			return solvers, err
		}
		if name == "Ant" {
			solver = &domain.AntSolver{
				ID:          ID,
				Name:        name,
				Description: desc,
				Model:       nil,
				IssueID:     issueID,
				Alpha:       0,
				Beta:        0,
				Rho:         0,
				Quantity:    0,
			}
		} else {
			return solvers, errors.New("unknown solver")
		}
		solvers = append(solvers, solver)
	}
	if err = rows.Err(); err != nil {
		return solvers, err
	}

	return solvers, nil
}
