package pubsub

import (
	"BookShop/app_book/connect"
	"context"
	"encoding/json"

	"github.com/go-redis/redis/v8"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var SubscriberClient *redis.PubSub
var ctx = context.Background()

func init() {
	channel := viper.GetString("redis.channel_name")
	rdb := connect.ConnectRedis()
	subscriber := rdb.Subscribe(ctx, channel)
	SubscriberClient = subscriber
}

func Publish(data interface{}) error {
	var ctx = context.Background()
	channel := viper.GetString("redis.channel_name")
	rdb := connect.ConnectRedis()
	err := rdb.Publish(ctx, channel, data).Err()
	if err != nil {
		return err
	}
	return nil
}

func Subscribe(data interface{}) {

	msg, err := SubscriberClient.ReceiveMessage(ctx)
	if err != nil {
		log.Error("Can't receive message for redis")
	}

	if err := json.Unmarshal([]byte(msg.Payload), &data); err != nil {
		log.Error("Can't unmarshal message")
	}
}
