/*
errs 用于错误处理相关的处理，
和 fmt.Errorf 相比，能记录一些调用栈

当需要返回一个 error ：

	func foo() error {
		// ...
		return errs.New("some error message") // main.foo fail: some error message
	}

当收到了一个 error 需要继续向上传递：

	func foo2() error {
		// ...
		err := foo()
		if err != nil {
			return errs.Wrap(err, "bad foo") // bad foo: main.foo2 fail: main.foo fail: some error message
		}
		return nil
	}
*/
package errs
