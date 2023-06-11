package util

import "github.com/spf13/viper"

type Config struct {
	PGHost     string `mapstructure:"PG_HOST"`
	PGPORT     string `mapstructure:"PG_PORT"`
	PGUsername string `mapstructure:"PG_USERNAME"`
	PGPassword string `mapstructure:"PG_PASSWORD"`
	PGDatabase string `mapstructure:"PG_Database"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("gin")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	viper.Unmarshal(&config)
	return
}
