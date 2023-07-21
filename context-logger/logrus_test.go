package golang_context

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLoggingLevel(t *testing.T) {
	logger := logrus.New()
	logger.SetLevel(logrus.TraceLevel)

	logger.Trace("This is Trace")
	logger.Debug("This is Debug")
	logger.Info("This is Info")
	logger.Warn("This is Warn")
	logger.Error("This is Error")
}

func TestOutput(t *testing.T) {
	log := logrus.New()
	file, _ := os.OpenFile("application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	log.SetOutput(file)

	log.Trace("Trace Logging")
	log.Info("Hello Logging")
	log.Warn("Hello Logging")
	log.Error("Hello Logging")
	// log.Fatal("Fatal log")
}

func TestFormatter(t *testing.T) {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})

	log.Info("Hello Logging")
	log.Warn("Hello Logging")
	log.Error("Hello Logging")
}

func TestField(t *testing.T) {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})

	logg := log.WithField("method", "test")

	logg.Info("Hello Logging")
	logg.Warn("Hello Logging")
	logg.Error("Hello Logging")

	logg.WithFields(logrus.Fields{
		"client": "name",
		"role":   "admin",
	}).Error("Hello Logging")
}

func TestEntry(t *testing.T) {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})

	newEntry := logrus.NewEntry(logger)
	newEntry.Info("Hello newEntry")

	log := newEntry.WithField("method", "test")

	log.Info("Hello Logging")
	log.Warn("Hello Logging")
	log.Error("Hello Logging")
}
