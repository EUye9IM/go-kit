package errs_test

import (
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
	err := withFmt().Error()
	if err != "make a error: errs_test.withFmt fail: errs_test.makeError fail: err" {
		t.Log(err)
		t.Fail()
	}
	err = withOutFmt().Error()
	if err != "errs_test.withOutFmt fail: errs_test.makeError fail: err" {
		t.Log(err)
		t.Fail()
	}
}
