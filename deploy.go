package gjobctl

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/glue"
)

type DeployOption struct {
}

func (app *App) Deploy(opt *DeployOption) error {
	// JSONファイルからGlue Jobの設定を読み込む
	jn := app.config.JobName
	f, err := os.Open(jn + ".json")
	if err != nil {
		err = fmt.Errorf("%w:  The Json file name should be job_name.json specified in gjobctl.yml", err)
		return err
	}
	defer f.Close()

	// JSONデータを構造体に変換
	var job glue.Job
	json.NewDecoder(f).Decode(&job)

	if err != nil {
		return err
	}

	fmt.Println(job)

	// AWSセッションを作成
	sess := session.Must(session.NewSession(&aws.Config{
		Region: &app.config.Region},
	))

	// Glueクライアントを作成
	client := glue.New(sess)

	// Glue Jobを更新
	_, err = client.UpdateJob(&glue.UpdateJobInput{
		JobName: &jn,
		JobUpdate: &glue.JobUpdate{
			Command:           job.Command,
			Connections:       job.Connections,
			DefaultArguments:  job.DefaultArguments,
			Description:       job.Description,
			ExecutionProperty: job.ExecutionProperty,
			LogUri:            job.LogUri,
			MaxRetries:        job.MaxRetries,
			Role:              job.Role,
		},
	})
	if err != nil {
		return err
	}

	fmt.Println("Successfully updated Glue Job: " + jn)
	return nil
}
