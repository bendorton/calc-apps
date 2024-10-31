package main

import (
	"log"
	"os"

	"bendorton/calc-apps/handlers"

	"github.com/bendorton/calc-lib"
)

func main() {
	handler := handlers.NewHandler(os.Stdout, &calc.Addition{})

	err := handler.Handle(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
}
