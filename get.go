package gjobctl

import "fmt"

type Get struct {
	JobName *string `arg:"" name:"jobname" help:"enter the name of the Glue Job to be getted"`
}

func (*App) Get(opt *Get) error {
	fmt.Println("Start Get")
	return nil
}
