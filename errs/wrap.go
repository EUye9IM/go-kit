package errs

import (
	"fmt"
)

// Wrap 用于包装一个错误。
func Wrap(err error, format string, a ...any) error {
	return &errInnerType{
		pc:  getPc(),
		msg: fmt.Sprintf(format, a...),
		err: err,
	}
}
