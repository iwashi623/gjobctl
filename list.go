package gjobctl

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/glue"
)

type ListOption struct {
}

func (app *App) List(opt *ListOption) error {
	// AWSのセッションを作成
	sess, err := session.NewSession(&aws.Config{
		Region: &app.config.Region},
	)
	if err != nil {
		return fmt.Errorf("failed to create session: %w", err)
	}
	sv := glue.New(sess)

	// Glue Jobの一覧を取得
	result, err := sv.GetJobs(&glue.GetJobsInput{
		MaxResults: aws.Int64(1000),
	})
	if err != nil {
		return fmt.Errorf("failed to get job list: %w", err)
	}

	// JobNameを改行しながら出力
	for _, job := range result.Jobs {
		fmt.Println(*job.Name)
	}
	return nil
}
