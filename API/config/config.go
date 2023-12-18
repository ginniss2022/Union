package initializer

import (
	// "errors"
	"os"

	"github.com/spf13/viper"
)

type EnvVars struct {
	PORT string `mapstructure:"PORT"`
	PG_HOST string `mapstructure:"PG_HOST"`
	PG_PORT string `mapstructure:"PG_PORT"`
	PG_USER string `mapstructure:"PG_USER"`
	PG_PASS string `mapstructure:"PG_PASS"`
	PG_DB   string `mapstructure:"PG_DB"`
}

func LoadConfig() (config EnvVars, err error) {
	env := os.Getenv("GO_ENV")
	if env == "production" {
		return EnvVars{
			PORT:  os.Getenv("PORT"),
			PG_HOST:  os.Getenv("PG_HOST"),
			PG_PORT: os.Getenv("PG_PORT"),
			PG_USER:         os.Getenv("PG_USER"),
			PG_PASS:         os.Getenv("PG_PASS"),
			PG_DB:         os.Getenv("PG_DB"),
		}, nil
	}

	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	// validate config here
	// `if config.MONGODB_URI == "" {
	// 	err = errors.New("MONGODB_URI is required")
	// 	return
	// }

	// if config.MONGODB_NAME == "" {
	// 	err = errors.New("MONGODB_NAME is required")
	// 	return
	// }
	return
}