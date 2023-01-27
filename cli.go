package gjobctl

import (
	"fmt"
	"os"
	"strings"

	"github.com/alecthomas/kong"
)

func cli(sub string, opts *CLIOptions, usage func()) error {
	app := New()
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

type CLIParseFunc func([]string) (string, *CLIOptions, func())

type CLIOptions struct {
	Get          *Get          `cmd:"" help:"Get GlueJob details in Json format."`
	ScriptDeploy *ScriptDeploy `cmd:"" help:"Deploy GlueJob script to S3."`
}

func CLI(parseArgs CLIParseFunc) (int, error) {
	fmt.Println("CLI Start")
	sub, opts, usage := parseArgs(os.Args[1:])

	err := cli(sub, opts, usage)
	if err != nil {
		return 1, err
	}
	return 0, nil
}

func ParseArgs(args []string) (string, *CLIOptions, func()) {
	if len(args) == 0 || len(args) > 0 && args[0] == "help" {
		args = []string{"--help"}
	}

	var opts CLIOptions
	kcxt := kong.Parse(&opts)

	// サブコマンドを取得
	sub := strings.Fields(kcxt.Command())[0]

	return sub, &opts, func() { kcxt.PrintUsage(true) }
}
