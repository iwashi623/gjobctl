package gjobctl

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/service/glue"
)

const (
	TimeoutForUpdate = 5 * time.Second
)

type UpdateOption struct {
	JobSettingFile *string `name:"job-setting-file" short:"f" description:"job setting file in json"`
}

func (app *App) Update(ctx context.Context, opt *UpdateOption) error {
	ctx, cancel := context.WithTimeout(ctx, TimeoutForUpdate)
	defer cancel()

	// JSONファイルからGlue Jobの設定を読み込む
	var fn string
	if opt.JobSettingFile != nil {
		fn = *opt.JobSettingFile
	} else {
		fn = app.getJobSettingFileName()
	}
	f, err := os.Open(fn)
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
	sess, err := app.createAWSSession()
	if err != nil {
		return err
	}

	// Glueクライアントを作成
	sv := glue.New(sess)

	// Glue Jobを更新
	r, err := sv.UpdateJobWithContext(ctx, &glue.UpdateJobInput{
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
