package gjobctl

import (
	"fmt"
	"os"
	"strings"

	"github.com/alecthomas/kong"
)

func cli(sub string, opts *CLIOptions, usage func()) error {
	app, err := New()
	fmt.Println((*app.Config).JobName)
	if err != nil {
		return err
	}
	switch sub {
	case "script-deploy":
		return app.ScriptDeploy(opts.ScriptDeploy)
	case "get":
		return app.Get(opts.Get)
	default:
		usage()
	}
	return nil
}

type CLIParseFunc func([]string) (string, *CLIOptions, func(), error)

type CLIOptions struct {
	Get          *Get          `cmd:"" help:"Get GlueJob details in Json format."`
	ScriptDeploy *ScriptDeploy `cmd:"" help:"Deploy GlueJob script to S3."`
}

func CLI(parseArgs CLIParseFunc) (int, error) {
	fmt.Println("CLI Start")
	sub, opts, usage, err := parseArgs(os.Args[1:])
	if err != nil {
		return 1, err
	}

	err = cli(sub, opts, usage)
	if err != nil {
		return 1, err
	}
	return 0, nil
}

func ParseArgs(args []string) (string, *CLIOptions, func(), error) {
	if len(args) == 0 || len(args) > 0 && args[0] == "help" {
		args = []string{"--help"}
	}

	var opts CLIOptions
	parser, err := kong.New(&opts)
	if err != nil {
		return "", nil, nil, fmt.Errorf("failed to new kong: %w", err)
	}
	kcxt, err := parser.Parse(args)
	if err != nil {
		return "", nil, nil, fmt.Errorf("failed to parse args: %w", err)
	}

	// サブコマンドを取得
	sub := strings.Fields(kcxt.Command())[0]

	return sub, &opts, func() { kcxt.PrintUsage(true) }, nil
}
