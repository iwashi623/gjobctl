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

type BaseOption interface {
}

type CreateOption struct {
	JobSettingFile *string `name:"job-setting-file" short:"f" description:"job setting file in json"`
}

func (app *App) Create(ctx context.Context, opt *CreateOption) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

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

	sess, err := app.createAWSSession()
	if err != nil {
		return err
	}

	// Glueクライアントを作成
	sv := glue.New(sess)

	// Glue Jobを更新
	r, err := sv.CreateJobWithContext(ctx, &glue.CreateJobInput{
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
