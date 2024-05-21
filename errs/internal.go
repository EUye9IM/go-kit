package errs

import (
	"runtime"
	"strings"
)

const (
	colon = ":"
)

func getFuncName() string {
	pc, _, _, ok := runtime.Caller(2)
	if !ok {
		return ""
	}
	fun := runtime.FuncForPC(pc)
	if fun == nil {
		return ""
	}
	nameSplite := strings.Split(fun.Name(), "/")
	name := nameSplite[len(nameSplite)-1]
	return name
}
