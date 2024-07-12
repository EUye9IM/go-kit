package testtool

import (
	"fmt"
	"reflect"
	"runtime/debug"
)

type Case interface {
	Assert(got, expect any)
}

type testItf interface {
	Logf(format string, args ...any)
	Fatalf(format string, args ...any)
}

type caseImpl struct {
	t testItf
}

func (c caseImpl) printStack() {
	c.t.Logf("%v", string(debug.Stack()))
}
func (c caseImpl) Assert(got, expect any) {
	if reflect.TypeOf(got) != reflect.TypeOf(expect) {
		c.printStack()
		c.t.Fatalf("different type\ngot: %v\nexpect: %v\n",
			reflect.TypeOf(got), reflect.TypeOf(expect),
		)
		return
	}
	if fmt.Sprintf("%#v", got) != fmt.Sprintf("%#v", expect) {
		c.printStack()
		c.t.Fatalf("different value\ngot: %v\nexpect: %v\n",
			got, expect,
		)
		return
	}
}

func TestCase(t testItf, c func(Case)) {
	tc := caseImpl{t}
	c(tc)
}
