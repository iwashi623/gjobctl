package gjobctl

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/glue"
)

type GetOption struct {
	JobName *string `arg:"" name:"jobname" help:"enter the name of the Glue Job to be getted"`
}

func (app *App) Get(opt *GetOption) error {
	// AWSのセッションを作成
	sess, err := session.NewSession(&aws.Config{
		Region: &app.config.Region},
	)
	if err != nil {
		return fmt.Errorf("failed to create session: %w", err)
	}

	// Glueのクライアントを作成
	sv := glue.New(sess)

	// Job名を指定
	jobName := *opt.JobName

	// Glue Jobの情報を取得
	result, err := sv.GetJob(&glue.GetJobInput{
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
