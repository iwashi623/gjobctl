package gjobctl

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/glue"
)

type CreateOption struct {
	JobSettingFile *string `name:"job-setting-file" description:"job setting file in json"`
}

func (app *App) Create(opt *CreateOption) error {
	// JSONファイルからGlue Jobの設定を読み込む
	var settingFileName string
	if opt.JobSettingFile == nil {
		settingFileName = app.config.JobName + ".json"
	} else {
		settingFileName = *opt.JobSettingFile
	}
	f, err := os.Open(settingFileName)
	if err != nil {
		return fmt.Errorf("failed to open setting file: %w", err)
	}
	defer f.Close()

	// JSONデータを構造体に変換
	var job glue.Job
	json.NewDecoder(f).Decode(&job)
	if err != nil {
		return fmt.Errorf("failed to decode json: %w", err)
	}

	// AWSセッションを作成
	sess := session.Must(session.NewSession(&aws.Config{
		Region: &app.config.Region},
	))

	// Glueクライアントを作成
	sv := glue.New(sess)

	// Glue Jobを更新
	r, err := sv.CreateJob(&glue.CreateJobInput{
		CodeGenConfigurationNodes: job.CodeGenConfigurationNodes,
		Command:                   job.Command,
		Connections:               job.Connections,
		DefaultArguments:          job.DefaultArguments,
		Description:               job.Description,
		ExecutionClass:            job.ExecutionClass,
		ExecutionProperty:         job.ExecutionProperty,
		GlueVersion:               job.GlueVersion,
		LogUri:                    job.LogUri,
		MaxRetries:                job.MaxRetries,
		Name:                      job.Name,
		NonOverridableArguments:   job.NonOverridableArguments,
		NotificationProperty:      job.NotificationProperty,
		NumberOfWorkers:           job.NumberOfWorkers,
		Role:                      job.Role,
		SecurityConfiguration:     job.SecurityConfiguration,
		SourceControlDetails:      job.SourceControlDetails,
		Timeout:                   job.Timeout,
		WorkerType:                job.WorkerType,
	})
	if err != nil {
		return fmt.Errorf("failed to create Glue Job: %w", err)
	}

	log.Println("Successfully created Glue Job: " + *r.Name)
	log.Println(job)
	return nil
}
