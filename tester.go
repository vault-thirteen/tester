package tester

import (
	"testing"
)

// Tester.

// This Package provides various useful Functions which help in writing Tests.

type Test struct {
	t *testing.T
}

// NewTest Function creates a new Test Object.
func NewTest(
	t *testing.T,
) *Test {
	return &Test{t: t}
}
