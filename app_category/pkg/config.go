package pkg

import (
	"fmt"

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
	fmt.Println(id)
	return id.String()
}
