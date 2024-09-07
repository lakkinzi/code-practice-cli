package main

import (
	"flag"
	"log"
	"practice/cmd"
	"practice/cmd/config"
)

const (
	modeSelect = "select"
	modeCreate = "create"
	modeLast   = "last"
	modeRun    = "run"
)

func main() {
	conf, err := config.Create()
	if err != nil {
		log.Fatal(err)
	}
	mode := flag.String("mode", "select", "")
	flag.Parse()
	switch *mode {
	case modeCreate:
		cmd.Create(conf)
	case modeSelect:
		cmd.Select(conf)
	case modeLast:
		cmd.Last(conf)
	case modeRun:
		cmd.Run(conf)
	}
}
