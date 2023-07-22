package logger

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func InitLogger() {
	// set default logger to print to stdout and prettify
	log.Logger = zerolog.New(os.Stdout).With().Timestamp().Caller().Logger().Output(zerolog.ConsoleWriter{Out: os.Stderr})
	// udp://127.0.0.1:12201 is log udp url
	hook, err := NewGraylogHook("udp://127.0.0.1:12201")
	if err != nil {
		panic(err)
	}
	//Set global logger with graylog hook
	log.Logger = log.Hook(hook)

	// Print and send sample log
	for range time.Tick(time.Millisecond * 200) {
		for range time.Tick(time.Millisecond * 200) {
			log.Info().Msg("info log")
			log.Warn().Msg("warning log")
			log.Debug().Msg("debug log")
		}
	}
}
