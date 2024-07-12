/*
logs 实现 slog 的 handler

没有很好的性能，只是在控制台上有人类友好的输出。对于日志性能高的场合请使用其他方案。
*/
package logs

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/fatih/color"
)

// PrintFunc 用于单条对日志生成字符串
type PrintFunc func(t time.Time, Level slog.Level, source *runtime.Frame, msg string, attr []slog.Attr) string

func defaultPrintFunc(t time.Time, level slog.Level, source *runtime.Frame, msg string, attr []slog.Attr) string {
	timeStr := color.BlackString("%v", t.Local().Format("2006-01-02 15:04:05.000"))

	sourceStr := ""
	if source != nil {
		sourceStr = color.BlackString("%v:%v ", filepath.Base(source.File), source.Line)
	}

	attrStr := strings.Builder{}
	for _, a := range attr {
		attrStr.WriteByte(' ')
		attrStr.WriteString(color.BlackString("%v=", a.Key))
		attrStr.WriteString(color.MagentaString(a.Value.String()))
	}

	levelStr := fmt.Sprintf("%6v", level)
	switch {
	case level < slog.LevelDebug,
		level < slog.LevelInfo:
		levelStr = color.BlueString(levelStr)
	case level < slog.LevelWarn:
		levelStr = color.GreenString(levelStr)
	case level < slog.LevelError:
		levelStr = color.YellowString(levelStr)
	default:
		levelStr = color.RedString(levelStr)
	}

	return fmt.Sprint(
		timeStr, " ",
		levelStr, " ", sourceStr,
		msg, attrStr.String())
}

// 选项
type Options struct {
	AddSource bool
	Level     slog.Leveler
	PrintFunc PrintFunc
}

func (o *Options) copy() *Options {
	var o2 Options
	if o != nil {
		o2 = *o
	}
	return &o2
}

type Handler struct {
	opt               Options
	preformattedAttrs []slog.Attr
	groups            string // all groups started from WithGroup
	mu                *sync.Mutex
	w                 io.Writer
}

func NewHandler(w io.Writer, opt *Options) *Handler {
	opt = opt.copy()
	if opt.PrintFunc == nil {
		opt.PrintFunc = defaultPrintFunc
	}
	return &Handler{
		opt: *opt,
		mu:  &sync.Mutex{},
		w:   w,
	}
}
func (h *Handler) clone() *Handler {
	// We can't use assignment because we can't copy the mutex.
	return &Handler{
		opt:               h.opt,
		preformattedAttrs: append([]slog.Attr(nil), h.preformattedAttrs...),
		groups:            h.groups,
		mu:                h.mu,
		w:                 h.w,
	}
}

// Enabled reports whether the handler handles records at the given level.
// The handler ignores records whose level is lower.
// It is called early, before any arguments are processed,
// to save effort if the log event should be discarded.
// If called from a Logger method, the first argument is the context
// passed to that method, or context.Background() if nil was passed
// or the method does not take a context.
// The context is passed so Enabled can use its values
// to make a decision.
func (h *Handler) Enabled(_ context.Context, l slog.Level) bool {
	minLevel := slog.LevelInfo
	if h.opt.Level != nil {
		minLevel = h.opt.Level.Level()
	}
	return l >= minLevel
}

// Handle handles the Record.
// It will only be called when Enabled returns true.
// The Context argument is as for Enabled.
// It is present solely to provide Handlers access to the context's values.
// Canceling the context should not affect record processing.
// (Among other things, log messages may be necessary to debug a
// cancellation-related problem.)
//
// Handle methods that produce output should observe the following rules:
//   - If r.Time is the zero time, ignore the time.
//   - If r.PC is zero, ignore it.
//   - Attr's values should be resolved.
//   - If an Attr's key and value are both the zero value, ignore the Attr.
//     This can be tested with attr.Equal(Attr{}).
//   - If a group's key is empty, inline the group's Attrs.
//   - If a group has no Attrs (even if it has a non-empty key),
//     ignore it.
func (h *Handler) Handle(_ context.Context, r slog.Record) error {
	var a []slog.Attr
	a = append(a, h.preformattedAttrs...)
	r.Attrs(func(attr slog.Attr) bool {
		a = append(a, slog.Attr{
			Key:   h.groups + attr.Key,
			Value: attr.Value,
		})
		return true
	})
	var str string
	if h.opt.AddSource {
		source, _ := runtime.CallersFrames([]uintptr{r.PC}).Next()
		str = h.opt.PrintFunc(r.Time, r.Level, &source, r.Message, a)
	} else {
		str = h.opt.PrintFunc(r.Time, r.Level, nil, r.Message, a)
	}
	h.mu.Lock()
	fmt.Fprintln(h.w, str)
	h.mu.Unlock()
	return nil
}

// WithAttrs returns a new Handler whose attributes consist of
// both the receiver's attributes and the arguments.
// The Handler owns the slice: it may retain, modify or discard it.
func (h *Handler) WithAttrs(attrs []slog.Attr) slog.Handler {
	h2 := h.clone()
	for i := range attrs {
		h2.preformattedAttrs = append(h2.preformattedAttrs,
			slog.Attr{Key: h.groups + attrs[i].Key, Value: attrs[i].Value},
		)
	}
	return h2
}

// WithGroup returns a new Handler with the given group appended to
// the receiver's existing groups.
// The keys of all subsequent attributes, whether added by With or in a
// Record, should be qualified by the sequence of group names.
//
// How this qualification happens is up to the Handler, so long as
// this Handler's attribute keys differ from those of another Handler
// with a different sequence of group names.
//
// A Handler should treat WithGroup as starting a Group of Attrs that ends
// at the end of the log event. That is,
//
//	logger.WithGroup("s").LogAttrs(level, msg, slog.Int("a", 1), slog.Int("b", 2))
//
// should behave like
//
//	logger.LogAttrs(level, msg, slog.Group("s", slog.Int("a", 1), slog.Int("b", 2)))
//
// If the name is empty, WithGroup returns the receiver.
func (h *Handler) WithGroup(name string) slog.Handler {
	h2 := h.clone()
	if name != "" {
		h2.groups += name + "."
	}
	return h2
}
