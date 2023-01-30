package gjobctl

import (
	"fmt"
	"os"
	"strings"

	"github.com/alecthomas/kong"
)

func cli(sub string, opts *CLIOptions, usage func()) error {
	app, err := New()
	if err != nil {
		return err
	}
	switch sub {
	case "list":
		return app.List(opts.List)
	case "get":
		return app.Get(opts.Get)
	case "deploy":
		return app.Deploy(opts.Deploy)
	case "script-deploy":
		return app.ScriptDeploy(opts.ScriptDeploy)
	case "run":
		return app.Run(opts.Run)
	default:
		usage()
	}
	return nil
}

type CLIParseFunc func([]string) (string, *CLIOptions, func(), error)

type CLIOptions struct {
	List         *ListOption         `cmd:"" help:"Get GlueJob List"`
	Get          *GetOption          `cmd:"" help:"Get GlueJob details in Json format."`
	Deploy       *DeployOption       `cmd:"" help:"Deploy GlueJob to Glue."`
	ScriptDeploy *ScriptDeployOption `cmd:"" help:"Deploy GlueJob script to S3."`
	Run          *RunOption          `cmd:"" help:"Run GlueJob."`
}

func CLI(parseArgs CLIParseFunc) (int, error) {
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
