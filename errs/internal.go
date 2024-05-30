package errs

import (
	"runtime"
	"strings"
)

const (
	colon = ": "
)

type errInnerType struct {
	pc            uintptr
	funcNameCache *string
	msg           string
	err           error
}

func (e *errInnerType) Error() string {
	if e == nil {
		return ""
	}
	funcName := e.getFuncName()
	var ret strings.Builder
	ret.WriteString(funcName)
	ret.WriteString(" fail")
	if e.msg != "" {
		ret.WriteString(colon)
		ret.WriteString(e.msg)
	}
	if e.err != nil {
		ret.WriteString(colon)
		ret.WriteString(e.err.Error())
	}
	return ret.String()
}
func (e *errInnerType) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.err
}
func (e *errInnerType) GetPc() uintptr {
	if e == nil {
		return 0
	}
	return e.pc
}

func getPc() uintptr {
	pc, _, _, ok := runtime.Caller(2)
	if !ok {
		return 0
	}
	return pc
}

type pcGetter interface {
	GetPc() uintptr
}

func getFuncName(p pcGetter) string {
	if p == nil || p.GetPc() == 0 {
		return ""
	}
	fun := runtime.FuncForPC(p.GetPc())
	if fun == nil {
		return ""
	}
	nameSplite := strings.Split(fun.Name(), "/")
	if len(nameSplite) == 0 {
		return ""
	}
	return nameSplite[len(nameSplite)-1]
}

func (e *errInnerType) getFuncName() string {
	if e.funcNameCache != nil {
		return *e.funcNameCache
	}
	e.funcNameCache = new(string)
	*e.funcNameCache = getFuncName(e)
	return *e.funcNameCache
}
