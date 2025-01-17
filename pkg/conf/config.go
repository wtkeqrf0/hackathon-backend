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

	TemplatePath string `yaml:"template_path" env:"TEMPLATE_PATH" env-required:"true"`

	Session struct {
		CookieName string        `yaml:"cookie_name" env:"COOKIE_NAME" env-default:"session_id"`
		CookiePath string        `yaml:"cookie_path" env:"COOKIE_PATH" env-default:"/api"`
		Duration   time.Duration `yaml:"duration" env:"COOKIE_DURATION" env-default:"720h"`
	} `yaml:"session"`

	Listen struct {
		MainPath   string `yaml:"main_path" env:"MAIN_PATH" env-default:"/api"`
		DomainName string `yaml:"domain_name" env:"DOMAIN_NAME" env-default:"localhost"`
		Port       int    `yaml:"port" env:"PORT" env-default:"3000"`
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

		Redis struct {
			DbId int `yaml:"db_id" env:"REDIS_DB" env-default:"0"`
			// if prod=1, host will always be "redis" (docker constant)
			Host string `yaml:"host" env:"REDIS_HOST" env-default:"127.0.0.1"`
			Port int    `yaml:"port" env:"REDIS_POST" env-default:"6379"`
		} `yaml:"redis"`
	} `yaml:"db"`

	Email struct {
		From     string `yaml:"from" env:"EMAIL_FROM" env-default:"you-together@gmail.com"`
		User     string `yaml:"user" env:"EMAIL_USER"`
		Password string `yaml:"password" env:"EMAIL_PASSWORD"`
		Host     string `yaml:"host" env:"EMAIL_STMP_HOST"`
		Port     int    `yaml:"port" env:"EMAIL_PORT"`
	} `yaml:"email"`

	Insurance struct {
		Contribution int `yaml:"contribution" env-required:"true"`
		Medical      int `yaml:"medical" env-required:"true"`
	} `yaml:"insurance"`
}

var (
	inst Config
	once sync.Once
)

// GetConfig builds the golang type by environment variables
// or (if not specified) configuration file and returns it
func GetConfig() *Config {
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
			inst.DB.Redis.Host = "redis"
			//gin.SetMode(gin.ReleaseMode)

		} else {
			logrus.SetLevel(logrus.DebugLevel)
		}
	})

	return &inst
}
