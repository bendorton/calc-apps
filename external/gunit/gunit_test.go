package gunit_test

import (
	"testing"

	"github.com/bendorton/calc-apps/external/gunit"
	"github.com/bendorton/calc-apps/external/should"
)

func TestMySuperCoolFixture(t *testing.T) {
	gunit.Run(t, new(MySuperCoolFixture))
}

type MySuperCoolFixture struct {
	*gunit.Fixture
}

func (this *MySuperCoolFixture) Test_FindMethod() {
	this.So(1, should.Equal, 1)
}
