package errs_test

import (
	"errors"
	"testing"

	"github.com/EUye9IM/go-kit/errs"
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
	err1 := withFmt().Error()
	if err1 != "errs_test.withFmt fail: make a error: errs_test.makeError fail: err" {
		t.Log(err1)
		t.Fail()
	}
	err2 := withOutFmt().Error()
	if err2 != "errs_test.withOutFmt fail: errs_test.makeError fail: err" {
		t.Log(err2)
		t.Fail()
	}
}
func TestUnwrap(t *testing.T) {
	err1 := withFmt()

	err2 := errs.Wrap(err1, "e2")

	err3 := errors.Unwrap(err2)
	if !errors.Is(err1, err3) {
		t.Log(err1)
		t.Log(err3)
		t.Fail()
	}
}
