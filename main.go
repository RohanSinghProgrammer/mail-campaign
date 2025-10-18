package main

import (
	"log"
)

func main() {
	err := producer()
	if err != nil {
		log.Fatalf("%s", err.Error())
	}
}
