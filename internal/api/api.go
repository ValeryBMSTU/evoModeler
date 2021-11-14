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

func DevPrint() {
	fmt.Println("package 'api' has been attach")
}

func (api *Api) PingHandler(ctx echo.Context) (err error) {
	fmt.Printf("%s", "Что-то прилетело в PingHandler...")
	ctx.Response().Writer.Write([]byte("pong"))
	return nil
}

func (api *Api) DoNothingHandler(ctx echo.Context) (err error) {
	return nil
}

func (api *Api) SingUpHandler(ctx echo.Context) (err error) {
	login := ctx.QueryParam("login")
	pass := ctx.QueryParam("pass")

	sessionID, err := api.Bl.CreateUser(login, pass)
	if err != nil {
		return err
	}

	data := &struct {
		SessionID int `json:"session_id"`
	}{sessionID}

	return ctx.JSON(http.StatusOK, data)
}

func (api *Api) LogInHandler(ctx echo.Context) (err error) {
	login := ctx.QueryParam("login")
	pass := ctx.QueryParam("pass")

	sessionID, err := api.Bl.CreateSession(login, pass)
	if err != nil {
		return err
	}

	data := &struct {
		SessionID int `json:"session_id"`
	}{sessionID}

	return ctx.JSON(http.StatusOK, data)
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

	data := &struct {
		Meta struct {
			Status string `json:"status"`
		} `json:"meta"`
	}{
		struct {
			Status string `json:"status"`
		}{"OK"},
	}

	return ctx.JSON(http.StatusOK, data)
}

func CreateApi(bl BL) (newApi *Api, err error) {
	return &Api{bl}, nil
}
