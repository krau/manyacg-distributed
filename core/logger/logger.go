package logger

import (
	"github.com/gookit/slog"
	"github.com/gookit/slog/handler"
	"github.com/gookit/slog/rotatefile"
	"github.com/krau/Picture-collector/core/config"
)

var L *slog.Logger

func init() {
	slog.DefaultChannelName = "core"
	newLogger := slog.New()
	defer newLogger.Flush()
	logLevel := slog.LevelByName(config.Cfg.App.Log.Level)
	logFilePath := config.Cfg.App.Log.FilePath
	logBackupNum := config.Cfg.App.Log.BackupNum
	var logLevels []slog.Level
	for _, level := range slog.AllLevels {
		if level <= logLevel {
			logLevels = append(logLevels, level)
		}
	}
	consoleH := handler.NewConsoleHandler(logLevels)
	fileH, err := handler.NewTimeRotateFile(
		logFilePath,
		rotatefile.EveryDay,
		handler.WithLogLevels(slog.AllLevels),
		handler.WithBackupNum(logBackupNum),
	)
	if err != nil {
		panic(err)
	}
	newLogger.AddHandlers(consoleH, fileH)
	L = newLogger
}
