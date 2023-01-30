package gjobctl

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/glue"
)

type RunOption struct {
	JobName *string `arg:"" name:"jobname" help:"enter the name of the Glue Job to run"`
}

func (app *App) Run(opt *RunOption) error {
	// AWSのセッションを作成
	sess, err := session.NewSession(&aws.Config{
		Region: &app.config.Region},
	)
	if err != nil {
		return err
	}

	// Glueのクライアントを作成
	sv := glue.New(sess)

	input := &glue.StartJobRunInput{
		JobName: opt.JobName,
	}

	result, err := sv.StartJobRun(input)
	if err != nil {
		return err
	}

	log.Printf("Successfully Started Glue Job. job-name: %s job-run-id:%s", *opt.JobName, *result.JobRunId)

	return nil
}
