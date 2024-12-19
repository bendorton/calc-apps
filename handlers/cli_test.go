package handlers

import (
	"bytes"
	"errors"
	"testing"

	"github.com/bendorton/calc-apps/external/should"
	"github.com/bendorton/calc-lib"
)

func TestHandler_WrongNumberOfArgs(t *testing.T) {
	handler := NewHandler(nil, nil)
	err := handler.Handle(nil)
	should.So(t, err, should.WrapError, errWrongNumberOfArgs)
}
func TestHandler_InvalidFirstArg(t *testing.T) {
	handler := NewHandler(nil, nil)
	err := handler.Handle([]string{"invalid", "1"})
	should.So(t, err, should.WrapError, errInvalidArg)
}
func TestHandler_InvalidSecondArg(t *testing.T) {
	handler := NewHandler(nil, nil)
	err := handler.Handle([]string{"1", "invalid"})
	should.So(t, err, should.WrapError, errInvalidArg)
}
func TestHandler_OutputWriterError(t *testing.T) {
	ugh := errors.New("ugh")
	handler := NewHandler(&ErringWriter{err: ugh}, &calc.Addition{})
	err := handler.Handle([]string{"1", "1"})
	should.So(t, err, should.WrapError, ugh)
	should.So(t, err, should.WrapError, errWriterFailure)
}

func TestHandler_Calculate(t *testing.T) {
	writer := &bytes.Buffer{}
	handler := NewHandler(writer, &calc.Addition{})
	err := handler.Handle([]string{"1", "1"})
	should.So(t, err, should.BeNil)
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
