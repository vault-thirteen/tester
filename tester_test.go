package tester

import (
	"testing"
)

func Test_NewTest(t *testing.T) {

	var aTestingT *testing.T
	var result *Test

	aTestingT = new(testing.T)
	result = NewTest(aTestingT)
	if result.t != aTestingT {
		t.FailNow()
	}
}
