package errs_test

import (
	"errors"
	"testing"

	"github.com/EUye9IM/go-kit/errs"
	"github.com/EUye9IM/go-kit/testtool"
)

func makeError() error {
	return errs.New("err")
}

func withFmt() error {
	return errs.Wrap(makeError(), "%v a error", "make")
}
func withOutFmt() error {
	return errs.Wrap(makeError(), "")
}

func TestWrap(t *testing.T) {
	testtool.TestCase(t, func(c testtool.Case) {
		c.Assert(withFmt().Error(), "errs_test.withFmt fail: make a error: errs_test.makeError fail: err")
		c.Assert(withOutFmt().Error(), "errs_test.withOutFmt fail: errs_test.makeError fail: err")
	})
}
func TestUnwrap(t *testing.T) {
	testtool.TestCase(t, func(c testtool.Case) {
		err1 := withFmt()
		err2 := errs.Wrap(err1, "e2")
		err3 := errors.Unwrap(err2)
		if !errors.Is(err1, err3) {
			t.Log(err1)
			t.Log(err3)
			t.Fail()
		}
	})
}
