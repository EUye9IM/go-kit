package main

import (
	"log/slog"
	"os"

	"github.com/EUye9IM/go-kit/logs"
)

func main() {
	slog.SetDefault(slog.New(logs.NewHandler(os.Stdout, &logs.Options{
		WithSource: true,
		Level:      slog.LevelDebug,
	})))

	data := map[int]string{
		1: "one",
		2: "two",
		3: "three",
	}

	slog.Debug("this is debug level", "data", data)
	slog.Info("this is info level", "data", data)
	slog.Warn("this is warn level", "data", data)
	slog.Error("this is error level", "data", data)
}
