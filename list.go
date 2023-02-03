package gjobctl

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/service/glue"
)

const (
	TimeoutForList       = 5 * time.Second
	MaxListCount   int64 = 1000
)

type ListOption struct {
}

func (app *App) List(ctx context.Context, opt *ListOption) error {
	ctx, cancel := context.WithTimeout(ctx, TimeoutForList)
	defer cancel()

	// AWSのセッションを作成
	sess, err := app.createAWSSession()
	if err != nil {
		return err
	}
	sv := glue.New(sess)

	max := MaxListCount
	// Glue Jobの一覧を取得
	result, err := sv.GetJobsWithContext(ctx, &glue.GetJobsInput{
		MaxResults: &max,
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
