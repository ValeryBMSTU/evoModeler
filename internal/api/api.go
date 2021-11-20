package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type BL interface {
	CreateUser(login, pass string) (sessionID int, err error)
	CreateSession(login, pass string) (sessionID int, err error)
	RemoveSession(sessionID int) (err error)
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

func CreateApi(bl BL) (newApi *Api, err error) {
	return &Api{bl}, nil
}
