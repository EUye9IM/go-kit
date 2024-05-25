package errs

import (
	"fmt"
)

// New 用于创建一个错误。
func New(format string, a ...any) error {
	return &errInnerType{
		pc:  getPc(),
		msg: fmt.Sprintf(format, a...),
	}
}
