package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Gentleelephant/pzlog/pzlog"
	"github.com/Gentleelephant/vue-project-backend/model"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
	"gopkg.in/natefinch/lumberjack.v2"
	"gopkg.in/yaml.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	ProjectConfig *Config
	Plog          *zap.Logger
	Filepath      string
	DB            *gorm.DB
	Rdb           *redis.Client
)

type LogConfig struct {
	lumberjack.Logger `yaml:",inline"`

	TimeFormat string `json:"timeformat" yaml:"timeformat"`
	//
	LogLevel string `json:"loglevel" yaml:"loglevel"`
	// whether to print to the console
	PrintConsole bool `json:"printconsole" yaml:"printconsole"`
	// json or console
	Encoder string `json:"encoder" yaml:"encoder"`
}

type Redis struct {
	Host string `json:"host" yaml:"host"`
	Port string `json:"port" yaml:"port"`
	Db   int    `json:"db" yaml:"db"`
}

type Config struct {
	// 数据库配置
	Database Database `json:"database" yaml:"database"`
	// Server配置
	Server Server `json:"server" yaml:"server"`
	// 日志配置
	LogConfig LogConfig `json:"log" yaml:"log"`
	// Redis配置
	Redis Redis `json:"redis" yaml:"redis"`
}

type Database struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	DBname   string `json:"dbname"`
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
	fmt.Println("file: ", string(file))
	fmt.Println("config: ", config)
	return config
}

// config DB
func initDB() *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second,   // 慢 SQL 阈值
			LogLevel:                  logger.Silent, // 日志级别
			IgnoreRecordNotFoundError: true,          // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,         // 禁用彩色打印
		},
	)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		ProjectConfig.Database.Username,
		ProjectConfig.Database.Password,
		ProjectConfig.Database.Host,
		ProjectConfig.Database.Port,
		ProjectConfig.Database.DBname)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}
	return db
}

func initRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", ProjectConfig.Redis.Host, ProjectConfig.Redis.Port),
		Password: "", // no password set
		DB:       ProjectConfig.Redis.Db,
	})
	return rdb
}

func createTable(db *gorm.DB) {
	err := db.AutoMigrate(&model.Account{})
	if err != nil {
		Plog.Error("create table failed", zap.String("table", "accounts"), zap.Error(err))
		return
	}
	err = db.AutoMigrate(&model.Menu{})
	if err != nil {
		return
	}
	err = db.AutoMigrate(&model.Role{})
	if err != nil {
		return
	}
	err = db.AutoMigrate(&model.AccountInfo{})
	if err != nil {
		return
	}
	if err != nil {
		Plog.Error("create table failed", zap.String("table", "menus"), zap.Error(err))
		return
	}

}

func Initial() {
	ProjectConfig = loadConfig(Filepath)
	Plog = pzlog.GetLogger(&pzlog.PzlogConfig{
		Logger: lumberjack.Logger{
			Filename: ProjectConfig.LogConfig.Filename,
			MaxSize:  ProjectConfig.LogConfig.MaxSize,
		},
		PrintConsole: true,
		Encoder:      ProjectConfig.LogConfig.Encoder,
	})
	DB = initDB()
	createTable(DB)
	Rdb = initRedis()
}
