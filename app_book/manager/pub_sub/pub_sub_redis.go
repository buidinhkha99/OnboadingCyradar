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

func Publish(data interface{}, channel string) error {
	var ctx = context.Background()
	rdb := connect.ConnectRedis()
	dataJson, err := json.Marshal(data)
	if err != nil {
		return err
	}
	err = rdb.Publish(ctx, channel, dataJson).Err()
	if err != nil {
		return err
	}
	return nil
}

func Subscribe(data interface{}) error {
	msg, err := SubscriberClient.ReceiveMessage(ctx)
	if err != nil {
		log.Error("Can't receive message for redis")
		return err
	}

	if err := json.Unmarshal([]byte(msg.Payload), &data); err != nil {
		log.Error("Can't unmarshal message")
		return err
	}
	return nil
}
