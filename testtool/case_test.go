package testtool_test

import (
	"testing"

	"github.com/EUye9IM/go-kit/testtool"
)

func TestCaseAssert(t *testing.T) {
	stub := testingStub{}

	// success
	testtool.TestCase(&stub, func(c testtool.Case) {
		// compareable
		c.Assert(1, 1)
		c.Assert("123", "123")
		// slice
		c.Assert([]int{1, 2, 3}, []int{1, 2, 3})
	})
	stub.CheckFatalf(t, nil)
	stub.Clear()

	// different type
	testtool.TestCase(&stub, func(c testtool.Case) {
		c.Assert(1, "1")
	})
	stub.CheckFatalf(t, [][]string{
		{"different type\ngot: %v\nexpect: %v\n", "int", "string"},
	})
	stub.Clear()

	// different value
	testtool.TestCase(&stub, func(c testtool.Case) {
		c.Assert(1, 2)
	})
	stub.CheckFatalf(t, [][]string{
		{"different value\ngot: %v\nexpect: %v\n", "1", "2"},
	})
	stub.Clear()
}
