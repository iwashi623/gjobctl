package gjobctl

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/service/glue"
)

type RunOption struct {
	JobName *string `arg:"" name:"jobname" help:"enter the name of the Glue Job to run"`
}

func (app *App) Run(ctx context.Context, opt *RunOption) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// AWSのセッションを作成
	sess, err := app.createAWSSession()
	if err != nil {
		return err
	}

	// Glueのクライアントを作成
	sv := glue.New(sess)

	input := &glue.StartJobRunInput{
		JobName: opt.JobName,
	}

	result, err := sv.StartJobRunWithContext(ctx, input)
	if err != nil {
		return fmt.Errorf("failed to start job: %w", err)
	}

	log.Printf("Successfully Started Glue Job. job-name: %s job-run-id:%s", *opt.JobName, *result.JobRunId)

	return nil
}
