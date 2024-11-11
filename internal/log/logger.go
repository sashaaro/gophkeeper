// Package logger initializes logger.
package log

import (
	"os"

	"github.com/rs/zerolog"
)

var loggerInstance = zerolog.New(os.Stderr).With().Timestamp().Logger()

type Config struct {
	Level string
}

func InitLogger(levelName string) error {
	level, err := zerolog.ParseLevel(levelName)
	if err != nil {
		return err
	}
	loggerInstance = loggerInstance.Level(level)
	return nil
}

func Debug(msg string, opts ...Opt) {
	log(loggerInstance.Debug(), msg, opts...)
}

func Warn(msg string, opts ...Opt) {
	log(loggerInstance.Warn(), msg, opts...)
}

func Info(msg string, opts ...Opt) {
	log(loggerInstance.Info(), msg, opts...)
}

func Error(msg string, opts ...Opt) {
	log(loggerInstance.Error(), msg, opts...)
}

func Fatal(msg string, opts ...Opt) {
	log(loggerInstance.Fatal(), msg, opts...)
}

func Panic(msg string, opts ...Opt) {
	log(loggerInstance.Panic(), msg, opts...)
}

func log(event *zerolog.Event, msg string, opts ...Opt) {
	for _, o := range opts {
		o(event)
	}
	event.Msg(msg)
}

type Opt func(event *zerolog.Event)

func Err(err error) Opt {
	return func(event *zerolog.Event) {
		event.Err(err)
	}
}

func Str(k, v string) Opt {
	return func(event *zerolog.Event) {
		event.Str(k, v)
	}
}
