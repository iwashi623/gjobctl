package gjobctl

import (
	"fmt"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type App struct {
	Config *AppConfig
}

func New() (*App, error) {
	var app App
	err := app.setup()
	if err != nil {
		log.Println(err)
	}
	return &app, nil
}

func (app *App) setup() error {
	yamlData, _ := ioutil.ReadFile("gjobctl.yml")

	// Config構造体にデシリアライズ
	var conf AppConfig
	if err := yaml.Unmarshal(yamlData, &conf); err != nil {
		fmt.Println(err)
		return err
	}

	app.Config = &conf
	return nil
}
