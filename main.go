package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("requires an argument")
		return
	}
	switch os.Args[1] {
	case "add":
		err := Add(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
	case "get":
		err := Get(os.Args[2])
		if err != nil {
			log.Fatal(err)
		}
	}
	return
}
