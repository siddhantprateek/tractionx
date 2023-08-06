package main

import (
	"log"
	api "tractionx/routes"
)

func main() {

	if err := api.Init(); err != nil {
		log.Fatal("unable to start application server.")
	}
}
