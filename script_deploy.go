package gjobctl

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
	"time"

	"github.com/aws/aws-sdk-go/service/glue"
	"github.com/aws/aws-sdk-go/service/s3"
)

type ScriptDeployOption struct {
	JobSettingFile  *string `name:"job-setting-file" short:"f" description:"job setting file in json"`
	ScriptLocalPath *string `arg:"" name:"script-local-path" short:"s" description:"script local path"`
}

func (app *App) ScriptDeploy(ctx context.Context, opt *ScriptDeployOption) error {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// JSONファイルからGlue Jobの設定を読み込む
	var fn string
	if opt.JobSettingFile != nil {
		fn = *opt.JobSettingFile
	} else {
		fn = app.getJobSettingFileName()
	}
	f, err := os.Open(fn)
	if err != nil {
		return fmt.Errorf("failed to open setting file: %w", err)
	}
	defer f.Close()

	// JSONデータを構造体に変換
	var job glue.Job
	err = json.NewDecoder(f).Decode(&job)
	if err != nil {
		return fmt.Errorf("failed to decode json: %w", err)
	}

	sess, err := app.createAWSSession()
	if err != nil {
		return err
	}

	sv := s3.New(sess)

	// ファイルを読み込む
	sf, err := os.Open(*opt.ScriptLocalPath)
	if err != nil {
		return fmt.Errorf("failed to open script file: %w", err)
	}
	defer sf.Close()

	// jobの設定ファイルからS3のパスを取得
	bucketName, bucketKey, err := getBucketNameAndKey(*job.Command.ScriptLocation)
	if err != nil {
		return err
	}

	// S3にアップロード
	_, err = sv.PutObjectWithContext(ctx, &s3.PutObjectInput{
		Bucket: &bucketName,
		Key:    &bucketKey,
		Body:   sf,
	})
	if err != nil {
		return fmt.Errorf("failed to upload file to S3: %w", err)
	}

	log.Println("File successfully uploaded to S3")
	return nil
}

func getBucketNameAndKey(s3Path string) (string, string, error) {
	re := regexp.MustCompile(`s3://([^/]+)/(.*)`)
	match := re.FindStringSubmatch(s3Path)
	if match[1] == "" || match[2] == "" {
		return "", "", fmt.Errorf("failed to parse s3 path: %s", s3Path)
	}
	return match[1], match[2], nil
}
