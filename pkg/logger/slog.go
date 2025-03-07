package logger

import (
	"io"
	"log/slog"
	"os"
)

func newSlog(env string) *slogLogger {
	var (
		l *slog.Logger
		f io.WriteCloser
	)

	switch env {
	case "DEV":
		l = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))
	default:
		var err error

		f, err = os.OpenFile("logs.json", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0777)
		if err != nil {
			panic("failed to open logs file: " + err.Error())
		}

		l = slog.New(slog.NewJSONHandler(io.MultiWriter(os.Stdout, f), &slog.HandlerOptions{
			Level: slog.LevelInfo,
		}))
	}

	return &slogLogger{
		l: l,
	}
}

type slogLogger struct {
	f io.WriteCloser

	l *slog.Logger
}

func (s *slogLogger) Close() error {
	if s.f == nil {
		return nil
	}

	return s.f.Close()
}

func (s *slogLogger) Info(msg string, args ...Arg) {
	if len(args) == 0 {
		s.l.Info(msg)
	} else {
		s.l.Info(msg, s.argsToAttrs(args)...)
	}
}

func (s *slogLogger) Warn(msg string, args ...Arg) {
	if len(args) == 0 {
		s.l.Warn(msg)
	} else {
		s.l.Warn(msg, s.argsToAttrs(args)...)
	}
}

func (s *slogLogger) Error(msg string, args ...Arg) {
	if len(args) == 0 {
		s.l.Error(msg)
	} else {
		s.l.Error(msg, s.argsToAttrs(args)...)
	}
}

func (s *slogLogger) Debug(msg string, args ...Arg) {
	if len(args) == 0 {
		s.l.Debug(msg)
	} else {
		s.l.Debug(msg, s.argsToAttrs(args)...)
	}
}

func (s *slogLogger) argsToAttrs(args []Arg) []any {
	vals := make([]any, 0, len(args))
	for _, arg := range args {
		vals = append(vals, slog.Any(arg.key, arg.val))
	}

	return vals
}
