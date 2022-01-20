package cache

import (
	"BookShop/connect"
	"context"

	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
)

var ctx = context.Background()

func InsertData(key, data string) error {
	rdb := connect.ConnectRedis()
	err := rdb.Set(ctx, key, data, 0).Err()
	if err != nil {
		return nil
	} else {
		return err
	}

}
func ServeJQueryWithRemoteCache(key string) string {
	rdb := connect.ConnectRedis()
	val2, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		log.Info("No data in remote cache")
		return ""
	}
	if err != nil {
		log.Info("Data has not been created in remote cache")
		return ""
	}
	log.Info("There are data in remote cache")
	return val2
}
