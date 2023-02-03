package gjobctl

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/glue"
)

type ListOption struct {
}

func (app *App) List(ctx context.Context, opt *ListOption) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// AWSのセッションを作成
	sess, err := session.NewSession(&aws.Config{
		Region: &app.config.Region},
	)
	if err != nil {
		return fmt.Errorf("failed to create session: %w", err)
	}
	sv := glue.New(sess)

	// Glue Jobの一覧を取得
	result, err := sv.GetJobsWithContext(ctx, &glue.GetJobsInput{
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
