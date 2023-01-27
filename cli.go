package gjobctl

import (
	"fmt"
	"os"
)

func cli(sub string) error {
	switch sub {
	case "script-deploy":
		ScriptDeploy()
	default:
		return fmt.Errorf("unknown subcommand: %s", sub)
	}
	return nil
}

func CLI() (int, error) {
	fmt.Println("CLI Start")
	sub := os.Args[1]

	err := cli(sub)
	if err != nil {
		return 1, err
	}
	return 0, nil
}
