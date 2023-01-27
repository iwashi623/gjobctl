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
	exitCode, err := gjobctl.CLI()
	if err != nil {
		log.Println(err)
	}
	os.Exit(exitCode)
}
