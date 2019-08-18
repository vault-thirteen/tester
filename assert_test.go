package tester

import (
	"errors"
	"testing"
)

func Test_errorIsSet(t *testing.T) {

	var result bool

	// Test #1. Negative.
	result = errorIsSet(nil)
	if result != false {
		t.FailNow()
	}

	// Test #2. Positive.
	result = errorIsSet(errors.New("Some Error"))
	if result != true {
		t.FailNow()
	}
}

func Test_errorIsEmpty(t *testing.T) {

	var result bool

	// Test #1. Negative.
	result = errorIsEmpty(errors.New("Some Error"))
	if result != false {
		t.FailNow()
	}

	// Test #2. Positive.
	result = errorIsEmpty(nil)
	if result != true {
		t.FailNow()
	}
}

func Test_interfacesAreEqual(t *testing.T) {

	type TestTypeX struct {
		Age  int
		Name string
	}

	var result bool

	// Test #1. Negative.
	result = interfacesAreEqual(
		TestTypeX{Age: 10, Name: "John"},
		TestTypeX{Age: 11, Name: "Jack"},
	)
	if result != false {
		t.FailNow()
	}

	// Test #2. Positive.
	result = interfacesAreEqual(
		TestTypeX{Age: 12, Name: "Alpha"},
		TestTypeX{Age: 12, Name: "Alpha"},
	)
	if result != true {
		t.FailNow()
	}
}

func Test_interfacesAreDifferent(t *testing.T) {

	type TestTypeX struct {
		Age  int
		Name string
	}

	var result bool

	// Test #1. Negative.
	result = interfacesAreDifferent(
		TestTypeX{Age: 12, Name: "Alpha"},
		TestTypeX{Age: 12, Name: "Alpha"},
	)
	if result != false {
		t.FailNow()
	}

	// Test #2. Positive.
	result = interfacesAreDifferent(
		TestTypeX{Age: 10, Name: "John"},
		TestTypeX{Age: 11, Name: "Jack"},
	)
	if result != true {
		t.FailNow()
	}
}
