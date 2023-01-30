package gjobctl

type RunOption struct {
	JobName *string `arg:"" name:"jobname" help:"enter the name of the Glue Job to be getted"`
}

func (app *App) Run(opt *RunOption) error {
	// // AWSのセッションを作成
	// sess, err := session.NewSession(&aws.Config{
	// 	Region: &app.config.Region},
	// )
	// if err != nil {
	// 	return err
	// }

	// // Glueのクライアントを作成
	// sv := glue.New(sess)

	// // Job名を指定
	// jobName := *opt.JobName
}
