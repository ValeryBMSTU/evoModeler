package main

import (
	"flag"
	"fmt"
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

func ParseFlags() (string, error) {
	var configPath string
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

func main() {
	api.DevPrint()

	cfgPath, err := ParseFlags()
	if err != nil {
		log.Fatal(err)
	}

	cfg, err := NewConfig(cfgPath)
	if err != nil {
		log.Fatal()
	}

	e := echo.New()
	e.GET("/ping", api.PingHandler)
	e.GET("/", api.DoNothingHandler)

	err = e.Start(cfg.Host + ":" + cfg.Port)
	if err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
