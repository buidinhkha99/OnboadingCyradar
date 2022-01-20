package handlers

import (
	"BookShop/model"
	pubsub "BookShop/pub_sub"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func CheckPubSub() {
	ch := make(chan model.BookPublish, 5)
	go CheckMessBook(ch)
	for data := range ch {
		var bookDetail model.DetailBook
		switch data.Description {
		case "GetGroupCategory":

			// get data in database
			bookDetail.GroupBooks, _ = db.GetGroupBookByID(data.IdBook)
			for _, group := range bookDetail.GroupBooks {
				bookDetail.Category = append(bookDetail.Category, db.GetCategory(group.CategoryID))
			}
			// subscribe data
			err := pubsub.Publish(bookDetail, data.Channel)
			if err != nil {
				log.Error("Can't publisher data, err: %v", err)
				continue
			}

		case "CreateGroupBook":

			err := db.CreateGoupBook(data.GroupBook)
			if err != nil {
				log.Error("Can't Create Group Book data, err: %v", err)
				bookDetail.Status = false
				err = pubsub.Publish(bookDetail, data.Channel)
				if err != nil {
					log.Error("Can't publisher data, err: %v", err)
					continue
				}
				return
			}

			// subscribe data
			bookDetail.Status = true
			err = pubsub.Publish(bookDetail, data.Channel)
			if err != nil {
				log.Error("Can't publisher data, err: %v", err)
				continue
			}
		case "UpdateGroupBook":
			err := db.UpdateGroupBook(data.GroupBook)
			if err != nil {
				log.Error("Can't Update Group Book data, err: %v", err)
				bookDetail.Status = false
				err = pubsub.Publish(bookDetail, data.Channel)
				if err != nil {
					log.Error("Can't publisher data, err: %v", err)
					continue
				}
				return
			}

			// subscribe data
			bookDetail.Status = true
			err = pubsub.Publish(bookDetail, data.Channel)
			if err != nil {
				log.Error("Can't publisher data, err: %v", err)
				continue
			}
		}

	}
}

func CheckMessBook(ch chan model.BookPublish) {
	var bookSub model.BookPublish
	channel := viper.GetString("redis.channel_book")
	for {
		err := pubsub.Subscribe(&bookSub, channel)
		if err != nil {
			log.Errorf("Can't not get data from book %v", err.Error())
		}
		ch <- bookSub
	}

}
