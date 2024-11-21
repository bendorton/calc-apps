package handlers

import (
	"bytes"
	"errors"
	"testing"

	"github.com/bendorton/calc-lib"
)

func assertError(t *testing.T, actual, target error) {
	t.Helper()
	if !errors.Is(actual, target) {
		t.Errorf("expected %v, got %v", target, actual)
	}
}

func TestHandler_WrongNumberOfArgs(t *testing.T) {
	handler := NewHandler(nil, nil)
	err := handler.Handle(nil)
	assertError(t, err, errWrongNumberOfArgs)
}
func TestHandler_InvalidFirstArg(t *testing.T) {
	handler := NewHandler(nil, nil)
	err := handler.Handle([]string{"invalid", "1"})
	assertError(t, err, errInvalidArg)
}
func TestHandler_InvalidSecondArg(t *testing.T) {
	handler := NewHandler(nil, nil)
	err := handler.Handle([]string{"1", "invalid"})
	assertError(t, err, errInvalidArg)
}
func TestHandler_OutputWriterError(t *testing.T) {
	ugh := errors.New("ugh")
	handler := NewHandler(&ErringWriter{err: ugh}, &calc.Addition{})
	err := handler.Handle([]string{"1", "1"})
	assertError(t, err, ugh)
	assertError(t, err, errWriterFailure)
}

func TestHandler_Calculate(t *testing.T) {
	writer := &bytes.Buffer{}
	handler := NewHandler(writer, &calc.Addition{})
	err := handler.Handle([]string{"1", "1"})
	assertError(t, err, nil)
	if writer.String() != "2" {
		t.Errorf("expected 2, got: %s", writer.String())
	}
}

type ErringWriter struct {
	err error
}

func (this *ErringWriter) Write(p []byte) (n int, err error) {
	return 0, this.err
}
