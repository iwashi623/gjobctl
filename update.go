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

type UpdateOption struct {
	JobSettingFile *string `name:"job-setting-file" short:"f" description:"job setting file in json"`
}

func (app *App) Update(opt *UpdateOption) error {
	// JSONファイルからGlue Jobの設定を読み込む
	var settingFileName string
	if opt.JobSettingFile != nil {
		settingFileName = *opt.JobSettingFile
	} else if app.config.JobSettingFile != "" {
		settingFileName = app.config.JobSettingFile
	} else {
		settingFileName = app.config.JobName + ".json"
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
	r, err := sv.UpdateJob(&glue.UpdateJobInput{
		JobName: job.Name,
		JobUpdate: &glue.JobUpdate{
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
			NonOverridableArguments:   job.NonOverridableArguments,
			NotificationProperty:      job.NotificationProperty,
			NumberOfWorkers:           job.NumberOfWorkers,
			Role:                      job.Role,
			SecurityConfiguration:     job.SecurityConfiguration,
			SourceControlDetails:      job.SourceControlDetails,
			Timeout:                   job.Timeout,
			WorkerType:                job.WorkerType,
		},
	})
	if err != nil {
		return fmt.Errorf("failed to updates Glue Job: %w", err)
	}

	log.Println("Successfully updatesd Glue Job: " + *r.JobName)
	log.Println(job)
	return nil
}
