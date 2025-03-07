package logger

import "context"

type Logger interface {
	Info(msg string, args ...Arg)
	Warn(msg string, args ...Arg)
	Error(msg string, args ...Arg)
	Debug(msg string, args ...Arg)

	Close() error
}

const (
	Slog = iota
)

func NewLogger(t uint, env string) Logger {
	switch t {
	case Slog:
		return newSlog(env)
	default:
		panic("invalid logger type")
	}
}

func WithArg(key, val string) Arg {
	return Arg{
		key: key,
		val: val,
	}
}

type Arg struct {
	key string
	val string
}

type loggerCtxKeyType string

const (
	loggerCtxKey loggerCtxKeyType = "ctx_logger"
)

func WrapLogger(ctx context.Context, log Logger) context.Context {
	return context.WithValue(ctx, loggerCtxKey, log)
}

func UnwrapLogger(ctx context.Context) Logger {
	return ctx.Value(loggerCtxKey).(Logger)
}
