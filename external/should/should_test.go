package should_test

import (
	"errors"
	"testing"
	"time"

	"github.com/bendorton/calc-apps/external/should"
)

type FakeT struct {
	err    error
	helped bool
}

func (this *FakeT) Helper()             { this.helped = true }
func (this *FakeT) Error(values ...any) { this.err = values[0].(error) }

func pass(t *testing.T, actual any, assert should.Assertion, expected ...any) {
	fakeT := &FakeT{}
	should.So(fakeT, actual, assert, expected...)
	if fakeT.err != nil {
		t.Errorf("should not get an error, got [%v]", fakeT.err)
	}
}
func fail(t *testing.T, actual any, assert should.Assertion, expected ...any) {
	fakeT := &FakeT{}
	should.So(fakeT, actual, assert, expected...)
	if !errors.Is(fakeT.err, should.ErrAssertionFailure) {
		t.Errorf("should get assertion error, but got [%v]", fakeT.err)
	} else {
		t.Log(fakeT.err)
	}

	if !fakeT.helped {
		t.Errorf("should have called Helper(), but didn't")
	}
}

func TestShouldEqual(t *testing.T) {
	pass(t, 1, should.Equal, 1)
	pass(t, []int{1, 2, 3}, should.Equal, []int{1, 2, 3})

	fail(t, 1, should.Equal, 2)
	fail(t, []int{1, 2, 3}, should.Equal, []int{1, 2, 3, 4})
}
func TestShouldNotEqual(t *testing.T) {
	pass(t, 1, should.NOT.Equal, 2)
	pass(t, []int{1, 2, 3}, should.NOT.Equal, []int{1, 2, 3, 4})

	fail(t, 1, should.NOT.Equal, 1)
	fail(t, []int{1, 2, 3}, should.NOT.Equal, []int{1, 2, 3})
}
func TestShouldBeTrue(t *testing.T) {
	pass(t, true, should.BeTrue)
	fail(t, false, should.BeTrue)
	fail(t, nil, should.BeTrue)
	fail(t, 0, should.BeTrue)
	fail(t, []int{1}, should.BeTrue)
}
func TestShouldBeFalse(t *testing.T) {
	pass(t, false, should.BeFalse)
	fail(t, true, should.BeFalse)
	fail(t, nil, should.BeFalse)
	fail(t, 1, should.BeFalse)
	fail(t, []int{1}, should.BeFalse)
}
func TestShouldBeNil(t *testing.T) {
	pass(t, nil, should.BeNil)
	fail(t, false, should.BeNil)
	fail(t, true, should.BeNil)
	fail(t, 1, should.BeNil)
	fail(t, 0, should.BeNil)
	fail(t, []int{1}, should.BeNil)
}
func TestShouldNotBeNil(t *testing.T) {
	pass(t, &time.Time{}, should.BeNil)
	fail(t, nil, should.BeNil)
}
