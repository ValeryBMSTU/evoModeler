package api

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ValeryBMSTU/evoModeler/internal/domain"
	"github.com/labstack/echo/v4"
)

type BL interface {
	CreateUser(login, pass string) (sessionID int, err error)
	CreateSession(login, pass string) (sessionID int, err error)
	RemoveSession(sessionID int) (err error)
	CheckSession(sessionID int) (isExist bool, err error)
	TakeUser(sessionID int) (user domain.User, err error)
	TakeSession(userID int) (session domain.Session, err error)
	CreateTask(taskName, solverName, genAlgName string, user domain.User) (task domain.Task, err error)
	TakeSolver(solverName string) (solver domain.Solver, err error)
	TakeSolvers() (solver []domain.Solver, err error)
	TakeIssues() (issues []domain.Issue, err error)
}

type Api struct {
	Bl BL
}

type RespFormat struct {
	Data interface{} `json:"data"`
	Meta Meta`json:"meta"`
}

type Meta struct {
	Info string `json:"info"`
	Err string `json:"err"`
}

type CustomMiddlewares struct {
	BL BL
}

func (m *CustomMiddlewares) AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		path := ctx.Path()
		if path != "/login" && path != "/singup" && path != "/logout" {
			strSessionID, err := ctx.Cookie("session_id")
			if err != nil {
				return err
			}
			sessionID, err := strconv.Atoi(strSessionID.Value)

			isExist, err := m.BL.CheckSession(sessionID)
			if !isExist {
				return errors.New("session does not exist")
			}
			if err != nil {
				return err
			}

			session, err := m.BL.TakeSession(sessionID)
			if err != nil {
				return err
			}

			user, err := m.BL.TakeUser(session.UserID)
			if err != nil {
				return err
			}

			ctx.Set("session", session)
			ctx.Set("user", user)
		}

		if err := next(ctx); err != nil {
			return err
		}

		return nil
	}
}

func (m *CustomMiddlewares) ErrorMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {

		if err := next(ctx); err != nil {
			ctx.Logger().Error(err)

			resp := &RespFormat {
				Data: nil,
				Meta: Meta{"OK", err.Error()},
			}

			return ctx.JSON(http.StatusInternalServerError, resp)
		}

		return nil
	}
}

func DevPrint() {
	fmt.Println("package 'api' has been attach")
}

func (api *Api) CheckAuth(ctx echo.Context) (err error) {
	_, ok := ctx.Get("user").(domain.User)
	if !ok {
		return errors.New("can not get user info")
	}

	_, ok = ctx.Get("session").(domain.User)
	if !ok {
		return errors.New("can not get session info")
	}

	return nil
}

func (api *Api) PingHandler(ctx echo.Context) (err error) {
	fmt.Printf("%s", "Что-то прилетело в PingHandler...")

	resp := &RespFormat {
		Data: struct {Pong string `json:"pong"`}{"pong"},
		Meta: Meta{"OK", ""},
	}

	return ctx.JSON(http.StatusOK, resp)
}

func (api *Api) DoNothingHandler(ctx echo.Context) (err error) {
	resp := &RespFormat {
		Data: nil,
		Meta: Meta{"OK", ""},
	}

	return ctx.JSON(http.StatusOK, resp)
}

func (api *Api) SingUpHandler(ctx echo.Context) (err error) {
	login := ctx.QueryParam("login")
	pass := ctx.QueryParam("pass")

	sessionID, err := api.Bl.CreateUser(login, pass)
	if err != nil {
		return err
	}

	resp := &RespFormat {
		Data: struct {SessionID int `json:"session_id"`}{sessionID},
		Meta: Meta{"OK", ""},
	}

	return ctx.JSON(http.StatusOK, resp)
}

func (api *Api) LogInHandler(ctx echo.Context) (err error) {
	login := ctx.QueryParam("login")
	pass := ctx.QueryParam("pass")

	sessionID, err := api.Bl.CreateSession(login, pass)
	if err != nil {
		return err
	}

	resp := &RespFormat {
		Data: struct {SessionID int `json:"session_id"`}{sessionID},
		Meta: Meta{"OK", ""},
	}

	return ctx.JSON(http.StatusOK, resp)
}

func (api *Api) LogOutHandler(ctx echo.Context) (err error) {
	sessionID, err := strconv.Atoi(ctx.QueryParam("session_id"))
	if err != nil {
		return err
	}

	err = api.Bl.RemoveSession(sessionID)
	if err != nil {
		return err
	}

	resp := &RespFormat {
		Data: nil,
		Meta: Meta{"OK", ""},
	}

	return ctx.JSON(http.StatusOK, resp)
}

func (api *Api) CreateTaskHandler(ctx echo.Context) (err error) {
	taskName := ctx.QueryParam("task_name")
	solverName := ctx.QueryParam("solver_name")
	genAlgName := ctx.QueryParam("gen_alg_name")

	if taskName == "" || solverName == "" || genAlgName == "" {
		return errors.New("missing param in 'CreateTaskHandler' endpoint")
	}

	user, ok := ctx.Get("user").(domain.User)
	if !ok {
		return errors.New("can not get user info")
	}
	_, ok = ctx.Get("session").(domain.Session)
	if !ok {
		return errors.New("can not get session info")
	}

	//taskType, err := api.Bl.TakeTaskType(taskTypeName)
	//if err != nil {
	//	return err
	//}

	task, err := api.Bl.CreateTask(taskName, solverName, genAlgName, user)
	if err != nil {
		return err
	}

	resp := &RespFormat {
		Data: task,
		Meta: Meta{"OK", ""},
	}

	return ctx.JSON(http.StatusOK, resp)
}

func (api *Api) GetSolversHandler(ctx echo.Context) (err error) {
	solvers, err := api.Bl.TakeSolvers()
	if err != nil {
		return err
	}

	resp := &RespFormat {
		Data: solvers,
		Meta: Meta{"OK", ""},
	}

	return ctx.JSON(http.StatusOK, resp)
}

func (api *Api) GetIssuesHandler(ctx echo.Context) (err error) {
	issues, err := api.Bl.TakeIssues()
	if err != nil {
		return err
	}

	resp := &RespFormat {
		Data: issues,
		Meta: Meta{"OK", ""},
	}

	return ctx.JSON(http.StatusOK, resp)
}

func CreateApi(bl BL) (newApi *Api, err error) {
	return &Api{bl}, nil
}

func CreateCustomMiddlewares(bl BL) (newMiddlewares *CustomMiddlewares, err error) {
	return &CustomMiddlewares{bl}, nil
}
