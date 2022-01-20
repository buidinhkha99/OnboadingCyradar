package pubsub

import (
	"BookShop/connect"
	"context"
	"encoding/json"

	log "github.com/sirupsen/logrus"
)

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

// channel := viper.GetString("redis.channel_name")
func Subscribe(data interface{}, channel string) error {
	ctx := context.Background()
	rdb := connect.ConnectRedis()
	subscriberClient := rdb.Subscribe(ctx, channel)
	msg, err := subscriberClient.ReceiveMessage(ctx)
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
