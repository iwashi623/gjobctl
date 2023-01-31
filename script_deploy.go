package gjobctl

import (
	"encoding/json"
	"fmt"
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
	fmt.Println(*opt.ScriptLocalPath)
	var settingFileName string
	if opt.JobSettingFile == nil {
		settingFileName = app.config.JobName + ".json"
	} else {
		settingFileName = *opt.JobSettingFile
	}
	fmt.Println(settingFileName)
	f, err := openSttingFile(settingFileName)
	if err != nil {
		return fmt.Errorf("failed to open setting file: %w", err)
	}
	defer f.Close()

	// JSONデータを構造体に変換
	var job glue.Job
	err = json.NewDecoder(f).Decode(&job)
	if err != nil {
		return err
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

	fmt.Println(bucketName)
	fmt.Println(bucketKey)
	// S3にアップロード
	_, err = sv.PutObject(&s3.PutObjectInput{
		Bucket: &bucketName,
		Key:    &bucketKey,
		Body:   sf,
	})
	if err != nil {
		return fmt.Errorf("failed to upload file to S3: %w", err)
	}

	fmt.Println("File successfully uploaded to S3")
	return nil
}

func openSttingFile(fileName string) (*os.File, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	return f, nil
}
