package main

import (
	"flag"
	"fmt"
	"github.com/ValeryBMSTU/evoModeler/internal/auth"
	"github.com/ValeryBMSTU/evoModeler/internal/bl"
	"github.com/ValeryBMSTU/evoModeler/internal/da"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
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
	CreateTaskHandler(ctx echo.Context) (err error)
	GetIssuesHandler(ctx echo.Context) (err error)
	GetSolversHandler(ctx echo.Context) (err error)
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
	e.POST("/task", api.CreateTaskHandler)
	e.GET("/issue/list", api.GetIssuesHandler)
	e.GET("/solver/list", api.GetSolversHandler)
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

	serverDa, err := da.CreateDa()
	if err != nil {
		log.Fatal(err)
	}
	serverBl, err := bl.CreateBl(serverDa)
	if err != nil {
		log.Fatal(err)
	}

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	conn, err := grpc.Dial("127.0.0.1:5301", opts...)

	if err != nil {
		grpclog.Fatalf("fail to dial: %v", err)
	}

	defer conn.Close()

	authClient := auth.NewAuthClient(conn)
	//request := &auth.Request{
	//	Message: args[1],
	//}
	//response, err := client.Do(context.Background(), request)
	//
	//if err != nil {
	//	grpclog.Fatalf("fail to dial: %v", err)
	//}
	//
	//fmt.Println(response.Message)

	serverApi, err := api.CreateApi(serverBl, authClient, true)
	if err != nil {
		log.Fatal(err)
	}
	server, err := CreateEchoServer(serverApi)
	if err != nil {
		log.Fatal(err)
	}

	mw, err := api.CreateCustomMiddlewares(serverBl)
	if err != nil {
		log.Fatal(err)
	}

	server.Use(mw.ErrorMiddleware)
	server.Use(mw.AuthMiddleware)

	err = server.Start(":" + cfg.Port)
	if err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
