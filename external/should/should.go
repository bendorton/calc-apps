package should

import (
	"errors"
	"fmt"
	"reflect"
)

type testingT interface {
	Helper()
	Error(...any)
}

type Assertion func(actual any, expected ...any) error

func So(t testingT, actual any, assert Assertion, expected ...any) bool {
	err := assert(actual, expected...)
	if err != nil {
		t.Helper()
		t.Error(err)
	}
	return err == nil
}

var ErrAssertionFailure = errors.New("assertion failure")

func Equal(actual any, EXPECTED ...any) error {
	if reflect.DeepEqual(actual, EXPECTED[0]) {
		return nil
	}
	return fmt.Errorf("%w: got [%v] want [%v]", ErrAssertionFailure, actual, EXPECTED[0])
}
func BeTrue(actual any, _ ...any) error  { return Equal(actual, true) }
func BeFalse(actual any, _ ...any) error { return Equal(actual, false) }
func BeNil(actual any, _ ...any) error   { return Equal(actual, nil) }
func WrapError(actual any, expected ...any) error {
	actualErr, _ := actual.(error)
	expectedErr, _ := expected[0].(error)
	if !errors.Is(actualErr, expectedErr) {
		return fmt.Errorf("%w: expected [%v] to wrap [%v], but it didn't", ErrAssertionFailure, actual, expected[0])
	}
	return nil
}

type negated struct{}

var NOT negated

func (negated) Equal(actual any, expected ...any) error {
	err := Equal(actual, expected...)
	if err == nil {
		return fmt.Errorf("%w: values not expected to be equal", ErrAssertionFailure)
	}
	return nil
}
func (negated) BeNil(actual any, _ ...any) error {
	err := BeNil(actual)
	if err == nil {
		return fmt.Errorf("%w: value should be nil", ErrAssertionFailure)
	}
	return nil
}
