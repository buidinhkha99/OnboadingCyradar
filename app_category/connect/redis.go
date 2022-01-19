package connect

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

func ConnectRedis() *redis.Client {
	address := viper.GetString("redis.address")
	password := viper.GetString("redis.password")
	db := viper.GetInt("redis.db")
	rdb := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       db,
	})
	return rdb
}
