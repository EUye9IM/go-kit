package logs_test

import (
	"bytes"
	"log/slog"
	"strings"
	"testing"

	"github.com/EUye9IM/go-kit/logs"
)

func TestHandlerWithBeforeWithGroup(t *testing.T) {
	var b bytes.Buffer
	logger := slog.New(logs.NewHandler(&b, nil))
	logger = logger.With("key1", "value1")
	logger = logger.WithGroup("group")
	logger.Info("msg", "key2", "value2")
	_, out, _ := strings.Cut(b.String(), "INFO")
	if strings.TrimSpace(out) != "msg key1=value1 group.key2=value2" {
		t.Error(out)
	}
}
