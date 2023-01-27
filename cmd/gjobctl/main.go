package main

import (
	"log"
	"os"

	"github.com/iwashi623/gjobctl"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	exitCode, err := gjobctl.CLI(gjobctl.ParseArgs)
	if err != nil {
		log.Fatalf("ERROR: %s", err)
	}
	os.Exit(exitCode)
}
