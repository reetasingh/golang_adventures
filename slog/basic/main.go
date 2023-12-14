package main

import (
	"log"
	"os"

	"github.com/google/uuid"
	"golang.org/x/exp/slog"
)

var programLevel = new(slog.LevelVar) // Info by default

type Employee struct {
	name string
	id   uuid.UUID
}

func (e Employee) LogValue() slog.Value {
	attr := slog.Group(e.id.String(),
		"name", e.name,
		"id", e.id)
	return slog.AnyValue(attr)
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: programLevel}))
	logger.Info("this is test message", "count", 3, "version", 8)

	// default logger is not one returned by slog
	defaultLogger := log.Default()
	defaultLogger.Println("count", 4)

	// now make slog logger as default logger
	slog.SetDefault(logger)

	// now slog logger is default logger
	defaultLogger = log.Default()
	defaultLogger.Println("count", 5)

	// group key value pairs
	grp := slog.Group("request", "method", "GET", "url", "google.com")
	logger.Info("this is test message", grp)

	// use attrs directly intead of key value pair
	attrs1 := slog.Int("count", 34)
	attrs2 := slog.Int("count", 35)
	logger.Info("example of using attrs", attrs1, attrs2)

	// log struct
	e1 := Employee{name: "abc", id: uuid.New()}
	logger.Info("example of logging struct", "employee", e1)
	// time=2023-12-13T21:52:08.112-08:00 level=INFO msg="example of logging struct" employee="{name:abc id:[195 143 117 123 200 154 75 195 181 246 196 84 160 45 141 211]
	// the uuid field is not being logged properly

	// fixed this by implementing LogValue() method for employee
	// time=2023-12-14T13:21:39.086-08:00 level=INFO msg="example of logging struct" employee="0404737d-68fc-41a0-bd40-666610956a24=[name=abc id=0404737d-68fc-41a0-bd40-666610956a24]"
	logger.Debug("this is debug message1")
	// this not be logged since level is not debug
	programLevel.Set(slog.LevelDebug)
	logger.Debug("this is debug message2")
	// this will now be logged
}
