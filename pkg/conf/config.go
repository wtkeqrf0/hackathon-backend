package conf

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"sync"
	"time"
)

type Config struct {
	// in docker image is always 1
	// can be specified by environment variable
	Prod int `yaml:"prod" env:"PROD" env-default:"0"`

	// specify the path where will the logs output be stored
	LogsPath string `yaml:"logs_path" env:"LOGS_PATH" env-default:"cons"`

	Token struct {
		Secret         string        `yaml:"secret" env:"SECRET" env-default:"1234"`
		AccessDuration time.Duration `yaml:"access_duration" env:"ACCESS_DURATION" env-default:"12h"`
		AccessName     string        `yaml:"access_name" env:"ACCESS_NAME" env-default:"Authorization"`
	} `yaml:"token"`

	Listen struct {
		MainPath string `yaml:"main_path" env:"MAIN_PATH" env-default:"/api"`
		Host     string `yaml:"host" env:"HOST" env-default:"localhost"`
		Port     int    `yaml:"port" env:"PORT" env-default:"3000"`
	} `yaml:"listen"`

	DB struct {
		Postgres struct {
			Username string `yaml:"username" env:"POSTGRES_USERNAME" env-default:"postgres"`
			DBName   string `yaml:"db_name" env:"POSTGRES_DB" env-default:"while.act"`
			Password string `yaml:"password" env:"POSTGRES_PASSWORD" env-default:"postgres"`
			// if prod=1, host will always be "postgres" (docker constant)
			Host string `yaml:"host" env:"POSTGRES_HOST" env-default:"127.0.0.1"`
			Port int    `yaml:"port" env:"POSTGRES_PORT" env-default:"5432"`
		} `yaml:"postgres"`
	} `yaml:"db"`
}

var (
	inst Config
	once sync.Once
)

// GetConfig builds the golang type by environment variables
// or (if not specified) configuration file and returns it
func GetConfig() Config {
	once.Do(func() {
		godotenv.Load()

		if err := cleanenv.ReadConfig("configs/config.yml", &inst); err != nil {
			logrus.WithError(err).Error("error occurred while reading config file")
			help, _ := cleanenv.GetDescription(&inst, nil)
			logrus.Info(&help)
			logrus.Exit(0)
		}

		if inst.Prod == 1 {
			inst.DB.Postgres.Host = "postgres"
			//gin.SetMode(gin.ReleaseMode)

		} else {
			logrus.SetLevel(logrus.DebugLevel)
		}

		if inst.Token.Secret == "1234" {
			logrus.Warn("secret word not set. Set it in env (\"SECRET\") or in config file. Default secret - \"1234\"")
		}
	})

	return inst
}
