package elogs

import (
	"log/slog"
	"os"
)

type Logger struct {
	ServiceName string `json:"service_name,omitempty"`
	PathToWrite string `json:"path_to_write"`
	TerminalMsg bool   `json:"terminal_msg"`
	LogLevel    int    `json:"log_level"`
}

var term = slog.New(slog.NewJSONHandler(os.Stdout, nil))

func LogToFile(path string) *slog.Logger {
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		term.Error("error opening file", path, err)
	}
	return slog.New(slog.NewJSONHandler(f, nil))
}

func (l *Logger) Info(msg string, args ...any) {
	if l.LogLevel == 0 || l.LogLevel == 1 || l.LogLevel == 2 {
		args = append(args, "service_name", l.ServiceName)
		if l.TerminalMsg {
			term.Info(msg, args...)
		}
		if l.PathToWrite != "" {
			LogToFile(l.PathToWrite).Info(msg, args...)
		}
	}

}

func (l *Logger) Error(msg string, args ...any) {
	if l.LogLevel == 1 || l.LogLevel == 2 {
		args = append(args, "service_name", l.ServiceName)
		if l.TerminalMsg {
			term.Error(msg, args...)
		}
		if l.PathToWrite != "" {
			LogToFile(l.PathToWrite).Error(msg, args...)
		}
	}
}

func (l *Logger) Warn(msg string, args ...any) {
	if l.LogLevel == 2 {
		args = append(args, "service_name", l.ServiceName)
		if l.TerminalMsg {
			term.Warn(msg, args...)
		}
		if l.PathToWrite != "" {
			LogToFile(l.PathToWrite).Warn(msg, args...)
		}
	}
}
