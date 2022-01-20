package pkg

import (
	"github.com/google/uuid"
	"github.com/spf13/viper"
)

func LoadConfig() error {
	viper.AddConfigPath(".")
	viper.SetConfigName("conf")
	viper.SetConfigType("toml")

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

func GenID() string {
	id := uuid.New()
	return id.String()
}
