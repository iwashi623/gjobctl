package gjobctl

import (
	"fmt"
	"os"
)

func cli(sub string) error {
	switch sub {
	case "script-deploy":
		ScriptDeploy()
	case "get":
		fmt.Println("Start Get")
	default:
		return fmt.Errorf("unknown subcommand: %s", sub)
	}
	return nil
}

type CLIParseFunc func([]string) (string, *CLIOptions, func(), error)

type CLIOptions struct {
	Get *Get `cmd:"" help:"Get GlueJob details in Json format."`
}

func CLI(parseArgs CLIParseFunc) (int, error) {
	fmt.Println("CLI Start")
	sub, _, _, _ := parseArgs(os.Args[1:])

	err := cli(sub)
	if err != nil {
		return 1, err
	}
	return 0, nil
}

func ParseArgs(args []string) (string, *CLIOptions, func(), error) {
	return "", nil, nil, nil
}
