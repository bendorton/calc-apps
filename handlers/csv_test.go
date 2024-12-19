package handlers

import (
	"bytes"
	"errors"
	"io"
	"log"
	"strings"
	"testing"

	"github.com/bendorton/calc-apps/external/should"
	"github.com/bendorton/calc-lib"
)

var csvInput = strings.Join([]string{
	"1,+,2",
	"2,-,1",
	"NaN,+,2",
	"1,+,NaN",
	"1,nop,2",
	"3,+,4",
	"3,*,4",
	"20,/,10",
}, "\n")

var csvOutput = strings.Join([]string{
	"1,+,2,3",
	"3,+,4,7",
	"",
}, "\n")

func TestCSVHandler(t *testing.T) {
	var logBuffer, outputBuffer bytes.Buffer
	logger := log.New(&logBuffer, "", log.LstdFlags)
	reader := strings.NewReader(csvInput)
	calculators := map[string]Calculator{"+": &calc.Addition{}}
	handler := NewCSVHandler(logger, reader, &outputBuffer, calculators)

	err := handler.Handle()

	should.So(t, err, should.BeNil)
	if outputBuffer.String() != csvOutput {
		t.Errorf("got %q want %q", outputBuffer.String(), csvOutput)
	}

	t.Log(logBuffer.String())
}

func TestCSVHandler_WriteError(t *testing.T) {
	logger := log.New(io.Discard, "", log.LstdFlags)
	reader := strings.NewReader(csvInput)
	ugh := errors.New("ugh")
	outputBuffer := ErringWriter{err: ugh}
	calculators := map[string]Calculator{"+": &calc.Addition{}}
	handler := NewCSVHandler(logger, reader, &outputBuffer, calculators)

	err := handler.Handle()

	should.So(t, err, should.WrapError, ugh)
}

func TestCSVHandler_ReadError(t *testing.T) {
	logger := log.New(io.Discard, "", log.LstdFlags)
	ugh := errors.New("ugh")
	reader := ErringReader{err: ugh}
	var outputBuffer bytes.Buffer
	calculators := map[string]Calculator{"+": &calc.Addition{}}
	handler := NewCSVHandler(logger, reader, &outputBuffer, calculators)

	err := handler.Handle()

	should.So(t, err, should.WrapError, ugh)
}

type ErringReader struct {
	err error
}

func (this ErringReader) Read(_ []byte) (int, error) {
	return 0, this.err
}
