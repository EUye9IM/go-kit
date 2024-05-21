package errs

import (
	"fmt"
)

// New 用于创建一个错误。
func New(format string, a ...any) error {
	name := getFuncName()
	str := fmt.Sprintf(format, a...)
	return fmt.Errorf("%v fail%v %v", name, colon, str)
}
