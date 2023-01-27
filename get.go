package gjobctl

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/glue"
)

type Get struct {
	JobName *string `arg:"" name:"jobname" help:"enter the name of the Glue Job to be getted"`
}

func (*App) Get(opt *Get) error {
	// AWSのセッションを作成
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Glueのクライアントを作成
	svc := glue.New(sess)

	// Job名を指定
	jobName := *opt.JobName

	// Glue Jobの情報を取得
	result, err := svc.GetJob(&glue.GetJobInput{
		JobName: aws.String(jobName),
	})
	if err != nil {
		fmt.Println(err)
		return err
	}

	// 取得したJob情報を表示
	fmt.Println(result)
	return nil
}
