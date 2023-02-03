package gjobctl

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/glue"
	"github.com/aws/aws-sdk-go/service/s3"
)

type ScriptDeployOption struct {
	JobSettingFile  *string `name:"job-setting-file" short:"f" description:"job setting file in json"`
	ScriptLocalPath *string `arg:"" name:"script-local-path" short:"s" description:"script local path"`
}

func (app *App) ScriptDeploy(opt *ScriptDeployOption) error {
	// JSONファイルからGlue Jobの設定を読み込む
	var settingFileName string
	if opt.JobSettingFile != nil {
		settingFileName = *opt.JobSettingFile
	} else if app.config.JobSettingFile != "" {
		settingFileName = app.config.JobSettingFile
	} else {
		settingFileName = app.config.JobName + ".json"
	}
	f, err := os.Open(settingFileName)
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

	sess, _ := session.NewSession(&aws.Config{
		Region: &app.config.Region},
	)
	sv := s3.New(sess)

	// ファイルを読み込む
	sf, err := os.Open(*opt.ScriptLocalPath)
	if err != nil {
		return fmt.Errorf("failed to open script file: %w", err)
	}
	defer sf.Close()

	// jobの設定ファイルからS3のパスを取得
	s3Path := *job.Command.ScriptLocation
	re := regexp.MustCompile(`s3://([^/]+)/(.*)`)
	match := re.FindStringSubmatch(s3Path)
	bucketName := match[1]
	bucketKey := match[2]

	// S3にアップロード
	_, err = sv.PutObject(&s3.PutObjectInput{
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
