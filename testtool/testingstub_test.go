package testtool_test

import (
	"fmt"
	"testing"
)

type testingStub struct {
	callingLog [][]string
}

func (t *testingStub) Clear() {
	t.callingLog = nil
}
func (t *testingStub) Logf(format string, args ...any) {}

func (t *testingStub) Fatalf(format string, args ...any) {
	arglist := []string{format}
	for _, a := range args {
		arglist = append(arglist, fmt.Sprint(a))
	}
	t.callingLog = append(t.callingLog, arglist)
}
func (t *testingStub) CheckFatalf(stdTest *testing.T, expect [][]string) {
	if !t.checkFatalf(expect) {
		stdTest.Fatalf("expect: %v, got: %v", expect, t.callingLog)
	}
}
func (t *testingStub) checkFatalf(expect [][]string) bool {
	if len(t.callingLog) != len(expect) {
		return false
	}
	for i, real := range t.callingLog {
		if len(real) != len(expect[i]) {
			return false
		}
		for j, real := range real {
			if real != expect[i][j] {
				return false
			}
		}
	}
	return true
}
