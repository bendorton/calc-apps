package gunit

import (
	"reflect"
	"strings"
	"testing"
)

type Setup interface {
	Setup()
}

func Run(t *testing.T, fixture any) {
	// scan fixture type looking for Setup(), Test_()
	fixtureType := reflect.TypeOf(fixture)

	// for each test:
	//  start new subtest
	//    create new instance of fixture
	//    call .Setup()
	//    call .Test_()

	for i := 0; i < fixtureType.NumMethod(); i++ {
		methodName := fixtureType.Method(i).Name
		if strings.HasPrefix(methodName, "SkipTest") {
			t.Run(methodName, func(t *testing.T) {
				t.Skip()
			})
		} else if strings.HasPrefix(methodName, "Test") {
			t.Run(methodName, func(t *testing.T) {
				fixtureValue := reflect.New(fixtureType.Elem())
				fixtureValue.Elem().FieldByName("Fixture").Set(
					reflect.ValueOf(&Fixture{T: t}),
				)
				fixtureWithSetup, ok := fixtureValue.Interface().(Setup)
				if ok {
					fixtureWithSetup.Setup()
				}
				fixtureValue.MethodByName(methodName).Call(nil)
			})
		}
	}
}

type Fixture struct{ *testing.T }

func (this *Fixture) So(actual any, assert assertion, expected ...any) {
	err := assert(actual, expected...)
	if err != nil {
		this.Helper()
		this.Error(err)
	}
}

type assertion func(actual any, expected ...any) error
