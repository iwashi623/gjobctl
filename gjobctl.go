package gjobctl

import (
	"log"
	"os"

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
	f, err := os.Open("gjobctl.yml")
	if err != nil {
		return err
	}
	defer f.Close()

	// Config構造体にデシリアライズ
	var conf AppConfig
	decoder := yaml.NewDecoder(f)
	if err := decoder.Decode(&conf); err != nil {
		return err
	}

	app.Config = &conf
	return nil
}
