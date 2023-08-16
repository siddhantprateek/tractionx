package main

import (
	"log"

	api "github.com/siddhantprateek/tractionx/api/routes"
)

func main() {

	if err := api.Init(); err != nil {
		log.Fatal("unable to start application server.")
	}
}
