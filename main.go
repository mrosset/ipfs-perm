package main

import (
	"errors"
	"flag"
	"github.com/str1ngs/util/console/command"
	"log"
	"os"
)

var (
	elog = log.New(os.Stderr, "", log.Lshortfile)
)

func main() {
	flag.Parse()
	command.Add("add", add, "add path to ipfs with saving permissions")
	command.Add("get", get, "get hash from ipfs with permissions")
	err := command.Run()
	if err != nil {
		elog.Fatal(err)
	}
}

func get() error {
	if len(command.Args()) < 1 {
		return errors.New("no hash specified")
	}
	hash := command.Args()[0]
	return Get(hash)
}

func add() error {
	if len(command.Args()) < 1 {
		return errors.New("no directory specified")
	}
	path := command.Args()[0]
	return Add(path)
}
