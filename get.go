package gjobctl

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/service/glue"
)

const (
	TimeoutForGet = 5 * time.Second
)

type GetOption struct {
	JobName *string `arg:"" name:"jobname" help:"enter the name of the Glue Job to be getted"`
}

func (app *App) Get(ctx context.Context, opt *GetOption) error {
	ctx, cancel := context.WithTimeout(ctx, TimeoutForGet)
	defer cancel()

	// AWSのセッションを作成
	sess, err := app.createAWSSession()
	if err != nil {
		return err
	}

	// Glueのクライアントを作成
	sv := glue.New(sess)

	// Job名を指定
	jobName := *opt.JobName

	// Glue Jobの情報を取得
	result, err := sv.GetJobWithContext(ctx, &glue.GetJobInput{
		JobName: &jobName,
	})
	if err != nil {
		return fmt.Errorf("failed to get job: %w", err)
	}

	// 取得したJob情報を表示
	job, _ := json.MarshalIndent(result, "", "  ")
	fmt.Println(string(job))
	return nil
}
