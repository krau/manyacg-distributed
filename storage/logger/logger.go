package logger

import (
	"github.com/gookit/slog"
	"github.com/gookit/slog/handler"
	"github.com/gookit/slog/rotatefile"
	"github.com/krau/manyacg/storage/config"
)

var L *slog.Logger

func init() {
	slog.DefaultChannelName = "storage"
	newLogger := slog.New()
	defer newLogger.Flush()
	logLevel := slog.LevelByName(config.Cfg.Log.Level)
	logFilePath := config.Cfg.Log.FilePath
	logBackupNum := config.Cfg.Log.BackupNum
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
