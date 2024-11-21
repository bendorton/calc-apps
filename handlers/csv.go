package handlers

import (
	"encoding/csv"
	"io"
	"log"
	"strconv"
)

type CSVHandler struct {
	logger      *log.Logger
	input       *csv.Reader
	output      *csv.Writer
	calculators map[string]Calculator
}

func NewCSVHandler(
	logger *log.Logger,
	input io.Reader,
	output io.Writer,
	calculators map[string]Calculator,
) *CSVHandler {
	return &CSVHandler{
		logger:      logger,
		input:       csv.NewReader(input),
		output:      csv.NewWriter(output),
		calculators: calculators,
	}
}

func (this *CSVHandler) Handle() error {
	for {
		record, err := this.input.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		a, err := strconv.Atoi(record[0])
		if err != nil {
			this.logger.Println("cannot parse operand:", record[0], err)
			continue
		}
		b, err := strconv.Atoi(record[2])
		if err != nil {
			this.logger.Println("cannot parse operand:", record[2], err)
			continue
		}
		calculator, ok := this.calculators[record[1]]
		if !ok {
			this.logger.Println("unsupported operation:", record[1])
			continue
		}
		c := calculator.Calculate(a, b)
		err = this.output.Write(append(record, strconv.Itoa(c)))
		if err != nil {
			break
		}
	}
	this.output.Flush()
	return this.output.Error()
}
