package gjobctl

import (
	"fmt"
)

// func uploadToS3(filePath string, s3Bucket string, s3ObjectKey string) error {
// 	sess, err := session.NewSession()
// 	if err != nil {
// 		return err
// 	}

// 	uploader := s3manager.NewUploader(sess)

// 	file, err := os.Open(filePath)
// 	if err != nil {
// 		return err
// 	}
// 	defer file.Close()

// 	_, err = uploader.Upload(&s3manager.UploadInput{
// 		Bucket: aws.String(s3Bucket),
// 		Key:    aws.String(s3ObjectKey),
// 		Body:   file,
// 	})
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

type ScriptDeploy struct {
}

func (*App) ScriptDeploy(opt *ScriptDeploy) error {
	fmt.Println("Start ScriptDeploy")
	// error := uploadToS3()
	// if error != nil {
	// 	return error
	// }
	return nil
}
