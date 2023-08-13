package platform

type WarpedLogger interface {
	Debugf(format string, args ...any)
	Infof(format string, args ...any)
	Warnf(format string, args ...any)
	Errorf(format string, args ...any)
}

// LogMiddleware to match the interface tls_client.Logger interface.
type LogMiddleware struct {
	w WarpedLogger
}

func NewLoggerWarp(logger WarpedLogger) *LogMiddleware {
	return &LogMiddleware{w: logger}
}

func (l *LogMiddleware) Debug(format string, args ...any) {
	l.w.Debugf(format, args...)
}
func (l *LogMiddleware) Info(format string, args ...any) {
	l.w.Infof(format, args...)
}
func (l *LogMiddleware) Warn(format string, args ...any) {
	l.w.Warnf(format, args...)
}
func (l *LogMiddleware) Error(format string, args ...any) {
	l.w.Errorf(format, args...)
}
