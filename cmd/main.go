package main

import (
	"flag"
	"fmt"
	"github.com/ValeryBMSTU/evoModeler/internal/bl"
	"github.com/ValeryBMSTU/evoModeler/internal/da"
	"log"
	"net/http"
	"os"

	"github.com/ValeryBMSTU/evoModeler/internal/api"
	"github.com/labstack/echo/v4"
	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Host string `yaml1:"host"`
	Port string `yaml1:"port"`
}

type API interface {
	PingHandler(ctx echo.Context) (err error)
	DoNothingHandler(ctx echo.Context) (err error)
	SingUpHandler(ctx echo.Context) (err error)
	LogInHandler(ctx echo.Context) (err error)
	LogOutHandler(ctx echo.Context) (err error)
}

type App struct {
	Api *echo.Echo
}

func ParseFlags() (configPath string, err error) {
	flag.StringVar(&configPath, "config", "./config.yaml", "path to config file")
	flag.Parse()
	fmt.Println(configPath)
	return configPath, nil
}

func NewConfig(configPath string) (*Config, error) {
	config := &Config{}

	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	d := yaml.NewDecoder(file)
	if err := d.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}

func CreateEchoServer(api API) (server *echo.Echo, err error) {
	e := echo.New()
	e.GET("/ping", api.PingHandler)
	e.GET("/", api.DoNothingHandler)
	e.POST("/singup", api.SingUpHandler)
	e.POST("/login", api.LogInHandler)
	e.DELETE("/logout", api.LogOutHandler)
	return e, nil
}

func main() {
	api.DevPrint()

	cfgPath, err := ParseFlags()
	if err != nil {
		log.Fatal(err)
	}

	cfg, err := NewConfig(cfgPath)
	if err != nil {
		log.Fatal(err)
	}

	da, err := da.CreateDa()
	if err != nil {
		log.Fatal(err)
	}
	bl, err := bl.CreateBl(da)
	if err != nil {
		log.Fatal(err)
	}
	api, err := api.CreateApi(bl)
	if err != nil {
		log.Fatal(err)
	}
	server, err := CreateEchoServer(api)
	if err != nil {
		log.Fatal(err)
	}

	err = server.Start(":" + cfg.Port)
	if err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
