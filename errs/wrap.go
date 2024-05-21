package errs

import (
	"fmt"
)

// Wrap 用于包装一个错误。
func Wrap(err error, format string, a ...any) error {
	if err == nil {
		return nil
	}
	name := getFuncName()
	if format != "" {
		format = format + colon + " "
	}
	if name != "" {
		format = format + name + " fail" + colon + " "
	}
	format = format + "%w"
	return fmt.Errorf(format, append(a, err)...)
}
