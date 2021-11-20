package api

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BL interface {
	CreateUser(login, pass string) (sessionID int, err error)
	CreateSession(login, pass string) (sessionID int, err error)
	RemoveSession(sessionID int) (err error)
	CheckSession(sessionID int) (isExist bool, err error)
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
		if err := next(ctx); err != nil {
			ctx.Error(err)
		}

		path := ctx.Path()
		if path == "/login" || path == "/singup" || path == "/logout" {
			return nil
		}

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

		return nil
	}
}

func DevPrint() {
	fmt.Println("package 'api' has been attach")
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
	taskType := ctx.QueryParam("task_type")

	if taskName == "" || taskType == "" {
		return errors.New("missing param in 'CreateTaskHandler' endpoint")
	}

	// заглушка

	resp := &RespFormat {
		Data: nil,
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
