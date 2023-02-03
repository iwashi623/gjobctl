package gjobctl

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type App struct {
	config *AppConfig
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
		return fmt.Errorf("failed to open config file: %w", err)
	}
	defer f.Close()

	// Config構造体にデシリアライズ
	var conf AppConfig
	decoder := yaml.NewDecoder(f)
	if err := decoder.Decode(&conf); err != nil {
		return fmt.Errorf("failed to decode config file: %w", err)
	}

	app.config = &conf
	return nil
}

type AppConfig struct {
	Region         string `yaml:"region"`
	JobName        string `yaml:"job_name"`
	JobSettingFile string `yaml:"job_setting_file"`
}
