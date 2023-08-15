package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"sync"
)

type TgConfig struct {
	DB    DBConfig    `json:"db"`
	Redis RedisConfig `json:"redis"`
	Log   LogConfig   `json:"log"`
	Set   SetConfig   `json:"set"`
}

type DBConfig struct {
	Host            string `json:"host"`
	Port            int    `json:"port"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	DBName          string `json:"dbname"`
	MaxOpenConn     int    `json:"maxopenconn"`
	MaxIdleConn     int    `json:"maxidleconn"`
	ConnMaxLifetime int    `json:"connmaxlifetime"`
}

type RedisConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Db       int    `json:"db"`
	Password string `json:"password"`
}

type LogConfig struct {
	Level          string `json:"level"`
	LogFormat      string `json:"format"`
	Path           string `json:"path"`
	Filename       string `json:"filename"`
	LogFileMaxSize int    `json:"logFileMaxSize"`
	LogStdout      bool   `json:"logStdout"`
	LogMaxAge      int    `json:"logMaxAge"`
}

type SetConfig struct {
	Url string `json:"url"`
}

var tgGameConfig *TgConfig
var once sync.Once

func InitConfig() {
	once.Do(func() {
		path, err := os.Getwd()
		if err != nil {
			panic(err)
		}
		fmt.Println(path)
		conf := viper.New()

		conf.SetConfigName(fmt.Sprintf("config"))
		conf.AddConfigPath(fmt.Sprintf("%s/", path))
		conf.SetConfigType("yaml")

		if err = conf.ReadInConfig(); err != nil {
			panic(err)
		}

		tgConfig := &TgConfig{}

		dbHost := conf.GetString("database.host")
		port := conf.GetInt("database.port")
		username := conf.GetString("database.username")
		password := conf.GetString("database.password")
		dbname := conf.GetString("database.dbname")
		maxOpenConn := conf.GetInt("database.maxopenconn")
		maxIdleConn := conf.GetInt("database.maxidleconn")
		connMaxLifetime := conf.GetInt("database.connmaxlifetime")

		dbConfig := DBConfig{
			Host:            dbHost,
			Port:            port,
			Username:        username,
			Password:        password,
			DBName:          dbname,
			MaxOpenConn:     maxOpenConn,
			MaxIdleConn:     maxIdleConn,
			ConnMaxLifetime: connMaxLifetime,
		}

		tgConfig.DB = dbConfig

		redisHost := conf.GetString("redis.host")
		redisPort := conf.GetInt("redis.port")
		redisPassword := conf.GetString("redis.password")
		redisDB := conf.GetInt("redis.db")

		redisConfig := RedisConfig{
			Host:     redisHost,
			Port:     redisPort,
			Password: redisPassword,
			Db:       redisDB,
		}

		tgConfig.Redis = redisConfig

		loglevel := conf.GetString("log.loglevel")
		logFormat := conf.GetString("log.logFormat")
		logPath := conf.GetString("log.logPath")
		logFileName := conf.GetString("log.logFileName")
		logFileMaxSize := conf.GetInt("log.logFileMaxSize")
		logMaxAge := conf.GetInt("log.logMaxAge")
		logStdout := conf.GetBool("log.logStdout")

		logConfig := LogConfig{
			Level:          loglevel,
			LogFormat:      logFormat,
			Path:           logPath,
			Filename:       logFileName,
			LogFileMaxSize: logFileMaxSize,
			LogMaxAge:      logMaxAge,
			LogStdout:      logStdout,
		}

		setUrl := conf.GetString("set.url")
		setConfig := SetConfig{
			Url: setUrl,
		}

		tgConfig.Set = setConfig

		tgConfig.Log = logConfig
		tgGameConfig = tgConfig

	})
}

func Config() *TgConfig {
	return tgGameConfig
}
