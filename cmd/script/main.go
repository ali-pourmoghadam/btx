package main

import (
	"log"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		log.Println("No argument provided.")
		return
	}
}
