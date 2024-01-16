package log

import "log/slog"

var Log *slog.Logger

func InitLogger() {
	Log = slog.Default()
}
