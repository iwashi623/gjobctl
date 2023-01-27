package gjobctl

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type ScriptDeployOption struct {
}

func (app *App) ScriptDeploy(opt *ScriptDeployOption) error {
	sess, _ := session.NewSession(&aws.Config{
		Region: &(*app.Config).Region},
	)
	client := s3.New(sess)

	// ファイルを読み込む
	file, err := os.Open((*app.Config).ScriptDIR + "/" + (*app.Config).ScriptName)
	if err != nil {
		return err
	}
	defer file.Close()

	s3Path := (*app.Config).BucketPath + "/" + (*app.Config).ScriptName

	// S3にアップロード
	_, err = client.PutObject(&s3.PutObjectInput{
		Bucket: &(*app.Config).BucketName,
		Key:    &s3Path,
		Body:   file,
	})
	if err != nil {
		return err
	}

	fmt.Println("File successfully uploaded to S3")
	return nil
}
