package gjobctl

import "fmt"

type Get struct {
	JobName *string `arg:"" name:"path" help:"Paths to remove." type:"path"`
}

func get() error {
	fmt.Println("Start Get")
	return nil
}
