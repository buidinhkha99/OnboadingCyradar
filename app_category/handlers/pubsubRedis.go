package handlers

import (
	pubsub "BookShop/app_category/manager/pub_sub"
	"BookShop/app_category/model"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type BookSub struct {
	ID          string `json:"id"`
	Description string `json:"description"`
}

func HandlerSubscribe() error {
	var bookSub BookSub
	bookDetail := model.DetailBook{}
	err := pubsub.Subscribe(bookSub)
	if err != nil {
		return err
	}
	if bookSub.Description == "GetGroup&Category" {
		bookDetail.GroupBooks, _ = db.GetGroupBookByID(bookSub.ID)
		for _, group := range bookDetail.GroupBooks {
			bookDetail.Category = append(bookDetail.Category, db.GetCategory(group.CategoryID))
		}
		channel := viper.GetString("redis.channel_category")
		err = pubsub.Publish(bookDetail, channel)
		if err != nil {
			log.Error("Can't publisher data, err: %v", err)
			return err
		}
	}
	return nil
}
