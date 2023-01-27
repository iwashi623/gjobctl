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
		log.Printf("ERROR: %s", err)
	}
	os.Exit(exitCode)
}
