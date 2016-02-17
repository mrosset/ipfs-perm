package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("OK")
	err := Add("testdata")
	if err != nil {
		log.Fatal(err)
	}
}
