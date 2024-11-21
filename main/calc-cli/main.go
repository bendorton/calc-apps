package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/bendorton/calc-apps/handlers"
	"github.com/bendorton/calc-lib"
)

func main() {
	var operation string
	flag.StringVar(&operation, "op", "+", "The mathematical operation")
	flag.Parse()

	calculator, err := loadCalculator(operation)
	if err != nil {
		log.Fatal(err)
	}

	handler := handlers.NewHandler(os.Stdout, calculator)
	err = handler.Handle(flag.Args())
	if err != nil {
		log.Fatal(err)
	}
}

func loadCalculator(operation string) (handlers.Calculator, error) {
	switch operation {
	case "+":
		return &calc.Addition{}, nil
	case "-":
		return &calc.Subtraction{}, nil
	case "*":
		return &calc.Multiplication{}, nil
	case "/":
		return &calc.Division{}, nil
	}
	return nil, fmt.Errorf("invalid operation %q", operation)
}
