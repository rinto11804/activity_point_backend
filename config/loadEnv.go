package config

import (
	"time"

	"github.com/spf13/viper"
)

type Config struct {
	DBHost         string `mapstructure:"POSTGRES_HOST"`
	DBUserName     string `mapstructure:"POSTGRES_USER"`
	DBUserPassword string `mapstructure:"POSTGRES_PASSWORD"`
	DBName         string `mapstructure:"POSTGRES_DB"`
	DBPort         string `mapstructure:"POSTGRES_PORT"`
	DBUrl          string `mapstructure:"DATABASE_URL"`

	Port               string `mapstructure:"PORT"`
	AWSRegion          string `mapstructure:"AWS_REGION"`
	AWSAccessKeyID     string `mapstructure:"AWS_ACCESS_KEY_ID"`
	AWSSecretAccessKey string `mapstructure:"AWS_SECRET_ACCESS_KEY"`

	JwtSecret    string        `mapstructure:"JWT_SECRET"`
	JwtExpiresIn time.Duration `mapstructure:"JWT_EXPIRED_IN"`
	JwtMaxAge    int           `mapstructure:"JWT_MAXAGE"`

	ClientOrigin string `mapstructure:"CLIENT_ORIGIN"`
}

func LoadConfig() (config Config, err error) {
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
	viper.AutomaticEnv()

	err = viper.Unmarshal(&config)
	return
}

// func LoadConfig() (config Config, err error) {
// 	return Config{
// 		DBUserName:         os.Getenv("POSTGRES_USER"),
// 		DBHost:             os.Getenv("POSTGRES_HOST"),
// 		DBUserPassword:     os.Getenv("POSTGRES_PASSWORD"),
// 		DBName:             os.Getenv("POSTGRES_DB"),
// 		DBPort:             os.Getenv("POSTGRES_PORT"),
// 		DBUrl:              os.Getenv("DATABASE_URL"),
// 		JwtSecret:          os.Getenv("JWT_SECRET"),
// 		Port:               os.Getenv("PORT"),
// 		AWSRegion:          os.Getenv("AWS_REGION"),
// 		AWSAccessKeyID:     os.Getenv("AWS_ACCESS_KEY_ID"),
// 		AWSSecretAccessKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
// 	}, nil
// }
