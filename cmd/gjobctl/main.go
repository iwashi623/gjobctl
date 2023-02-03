package main

import (
	"context"
	"log"
	"os"

	"github.com/iwashi623/gjobctl"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	ctx := context.Background()
	exitCode, err := gjobctl.CLI(ctx, gjobctl.ParseArgs)
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}
	os.Exit(exitCode)
}
