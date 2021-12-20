package config

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strings"
	"sync"
)

type AppConfig struct {
	App    Service     `yaml:"app"`
	Db     PostgresSQL `yaml:"postgres"`
	Logger Log         `yaml:"logger"`
}

type Service struct {
	Port int `yaml:"port"`
}

type PostgresSQL struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	DbName   string `yaml:"dbname"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	SslMode  bool   `yaml:"sslmode"`
}

type Log struct {
	LogFile   string `yaml:"filename"`
	LogFormat string `yaml:"format"`
	LogLevel  int    `yaml:"loglevel"`
}

var (
	once   sync.Once
	config *AppConfig
)

func GetConfig() *AppConfig {
	once.Do(func() {
		confFile, err := ioutil.ReadFile("app-config.yaml")
		if err != nil {
			panic(err)
		}
		err = yaml.Unmarshal(confFile, &config)
		if err != nil {
			panic(err)
		}
	})
	return config
}

func ConfigureLogger() {
	c := GetConfig()
	if strings.Compare(c.Logger.LogFormat, "text") == 0 { //text format
		log.SetFormatter(&log.TextFormatter{
			DisableColors: false,
			FullTimestamp: true,
		})
	} else { //json format
		log.SetFormatter(&log.JSONFormatter{}) // using default configuration
	}
	log.SetLevel(log.AllLevels[c.Logger.LogLevel])
	f, err := os.OpenFile(c.Logger.LogFile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	log.SetOutput(f)
}
