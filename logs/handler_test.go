package logs_test

import (
	"bytes"
	"log/slog"
	"strings"
	"testing"

	"github.com/EUye9IM/go-kit/logs"
	"github.com/EUye9IM/go-kit/testtool"
)

func TestHandlerWithBeforeWithGroup(t *testing.T) {
	testtool.TestCase(t, func(c testtool.Case) {
		var b bytes.Buffer
		logger := slog.New(logs.NewHandler(&b, nil))
		logger = logger.With("key1", "value1")
		logger = logger.WithGroup("group")
		logger.Info("msg", "key2", "value2")
		_, out, _ := strings.Cut(b.String(), "INFO")
		c.Assert(strings.TrimSpace(out), "msg key1=value1 group.key2=value2")
	})
}
