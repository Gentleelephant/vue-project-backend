package config

import (
	"fmt"
	"github.com/Gentleelephant/pzlog/pzlog"
	"go.uber.org/zap"
	"gopkg.in/yaml.v2"
	"os"
)

var (
	ProjectConfig *Config
	Plog          *zap.Logger
	Filepath      string
)

type LogConfig struct {
	Filename string `json:"filename" yaml:"filename"`

	TimeFormat string `json:"timeformat" yaml:"timeformat"`

	LogLevel string `json:"loglevel" yaml:"loglevel"`

	PrintConsole bool `json:"printconsole" yaml:"printconsole"`

	Encoder string `json:"encoder" yaml:"encoder"`

	// 日志文件最大大小
	MaxSize int `json:"maxsize" yaml:"maxsize"`

	MaxBackups int `json:"maxbackups" yaml:"maxbackups"`

	MaxAge int `json:"maxage" yaml:"maxage"`

	LocalTime bool `json:"localtime" yaml:"localtime"`

	Compress bool `json:"compress" yaml:"compress"`
}
type Config struct {
	// 数据库配置
	Database Database `json:"database" yaml:"database"`
	// Server配置
	Server Server `json:"server" yaml:"server"`
	// 日志配置
	LogConfig LogConfig `json:"log" yaml:"log"`
}

type Database struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Server struct {
	Port string `json:"port"`
}

// LoadConfig 加载配置
func loadConfig(filePath string) *Config {
	file, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	var config = &Config{}
	err = yaml.Unmarshal(file, config)
	if err != nil {
		panic(err)
	}
	return config
}

func Initial() {
	ProjectConfig = loadConfig(Filepath)
	fmt.Printf("config: %+v", ProjectConfig)
	Plog = pzlog.GetLogger(&pzlog.PzlogConfig{
		Filename:     ProjectConfig.LogConfig.Filename,
		PrintConsole: true,
		Encoder:      ProjectConfig.LogConfig.Encoder,
	})
}
