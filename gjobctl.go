package gjobctl

import (
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
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

func (app *App) getJobSettingFileName() string {
	if app.config.JobSettingFile != "" {
		return app.config.JobSettingFile
	} else {
		return app.config.JobName + ".json"
	}
}

func (app *App) createAWSSession() (*session.Session, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: &app.config.Region},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create session: %w", err)
	}
	return sess, nil
}

type AppConfig struct {
	Region         string `yaml:"region"`
	JobName        string `yaml:"job_name"`
	JobSettingFile string `yaml:"job_setting_file"`
}
