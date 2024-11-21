package handlers

import (
	"errors"
	"fmt"
	"io"
	"strconv"
)

type Calculator interface {
	Calculate(a, b int) int
}

type Handler struct {
	stdout     io.Writer
	calculator Calculator
}

func NewHandler(stdout io.Writer, calculator Calculator) *Handler {
	return &Handler{
		stdout:     stdout,
		calculator: calculator,
	}
}

func (this *Handler) Handle(args []string) error {
	if len(args) != 2 {
		return errWrongNumberOfArgs
	}

	a, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("%w: [%s] %w", errInvalidArg, args[0], err)
	}
	b, err := strconv.Atoi(args[1])
	if err != nil {
		return fmt.Errorf("%w: [%s] %w", errInvalidArg, args[1], err)
	}

	result := this.calculator.Calculate(a, b)

	_, err = fmt.Fprint(this.stdout, result)
	if err != nil {
		return fmt.Errorf("%w: %w", err, errWriterFailure)
	}
	return nil
}

var (
	errWrongNumberOfArgs = errors.New("usage: calc <a> <b>")
	errInvalidArg        = errors.New("invalid argument")
	errWriterFailure     = errors.New("writer failure")
)
