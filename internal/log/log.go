package log

import "log/slog"

var Log *slog.Logger

func InitLogger() *slog.Logger {
	if Log == nil {
		Log = slog.Default()
	}
	return Log
}
