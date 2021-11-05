package api

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func DevPrint() {
	fmt.Println("package 'api' has been attach")
}

func PingHandler(c echo.Context) error {
	fmt.Printf("%s", "Что-то прилетело в PingHandler...")
	c.Response().Writer.Write([]byte("pong"))
	return nil
}

func DoNothingHandler(c echo.Context) error {
	return nil
}
