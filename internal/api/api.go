package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/ValeryBMSTU/evoModeler/internal/bl"
	"github.com/labstack/echo/v4"
)

func DevPrint() {
	fmt.Println("package 'api' has been attach")
}

func PingHandler(ctx echo.Context) error {
	fmt.Printf("%s", "Что-то прилетело в PingHandler...")
	ctx.Response().Writer.Write([]byte("pong"))
	return nil
}

func DoNothingHandler(ctx echo.Context) error {
	return nil
}

func SingUpHandler(ctx echo.Context) error {
	login := ctx.QueryParam("login")
	pass := ctx.QueryParam("pass")

	sessionID, err := bl.CreateUser(login, pass)
	if err != nil {
		return err
	}

	data := &struct {
		SessionID int `json:"session_id"`
	}{sessionID}

	return ctx.JSON(http.StatusOK, data)
}

func LogInHandler(ctx echo.Context) error {
	login := ctx.QueryParam("login")
	pass := ctx.QueryParam("pass")

	sessionID, err := bl.CreateSession(login, pass)
	if err != nil {
		return err
	}

	data := &struct {
		SessionID int `json:"session_id"`
	}{sessionID}

	return ctx.JSON(http.StatusOK, data)
}

func LogOutHandler(ctx echo.Context) error {
	sessionID, err := strconv.Atoi(ctx.QueryParam("session_id"))
	if err != nil {
		return err
	}

	err = bl.RemoveSession(sessionID)
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
