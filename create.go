package gjobctl

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/glue"
)

type CreateOption struct {
	JobScriptFile *string `short:"f" long:"job-script-file" description:"Job script file path" required:"true"`
}

func (app *App) Create(opt *CreateOption) error {
	// JSONファイルからGlue Jobの設定を読み込む
	f, err := os.Open(*opt.JobScriptFile)
	if err != nil {
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
	sv := glue.New(sess)

	// Glue Jobを更新
	r, err := sv.CreateJob(&glue.CreateJobInput{
		Name:              job.Name,
		Command:           job.Command,
		Connections:       job.Connections,
		DefaultArguments:  job.DefaultArguments,
		Description:       job.Description,
		ExecutionProperty: job.ExecutionProperty,
		LogUri:            job.LogUri,
		MaxRetries:        job.MaxRetries,
		Role:              job.Role,
	})
	if err != nil {
		return fmt.Errorf("failed to create Glue Job: %w", err)
	}

	fmt.Println("Successfully created Glue Job: " + *r.Name)
	return nil
}
