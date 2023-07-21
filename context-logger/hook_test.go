package golang_context_test

import (
	"fmt"
	"testing"

	"github.com/sirupsen/logrus"
)

type SimpleHook struct {
}

func (s *SimpleHook) Levels() []logrus.Level {
	return []logrus.Level{logrus.ErrorLevel, logrus.PanicLevel}
}

func (s *SimpleHook) Fire(entry *logrus.Entry) error {
	fmt.Println("Hook called")
	return nil
}

func TestHook(t *testing.T) {
	log := logrus.New()
	log.AddHook(&SimpleHook{})

	log.Info("Hello Logging")
	log.Warn("Hello Logging")
	log.Error("Hello Logging")
}

func TestSingleton(t *testing.T) {
	logrus.Info("Hello World")
	logrus.Warn("Hello World")

	logrus.SetFormatter(&logrus.JSONFormatter{})

	logrus.Info("Hello World")
	logrus.Warn("Hello World")
}
