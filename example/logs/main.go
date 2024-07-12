package main

import (
	"fmt"
	"log/slog"
	"os"
	"runtime"
	"time"

	"github.com/EUye9IM/go-kit/logs"
)

func main() {
	opt := &logs.Options{
		AddSource: true,
		Level:     slog.LevelDebug,
	}
	slog.SetDefault(slog.New(logs.NewHandler(os.Stdout, opt)))

	slog.Info("handler with debug level", "opt", opt)
	slog.Debug("this is debug level")
	slog.Info("this is info level")
	slog.Warn("this is warn level")
	slog.Error("this is error level")
	fmt.Println()

	opt.AddSource = true
	slog.SetDefault(slog.New(logs.NewHandler(os.Stdout, opt)))
	slog.Info("handler WithSource", "opt", opt)
	fmt.Println()

	opt.PrintFunc = func(t time.Time, Level slog.Level, s *runtime.Frame, msg string, attr []slog.Attr) string {
		return fmt.Sprintf("%v %v %v %v %v", t, s, Level, msg, attr)
	}
	slog.SetDefault(slog.New(logs.NewHandler(os.Stdout, opt)))
	slog.Info("handler with your own print function", "opt", opt)
}
