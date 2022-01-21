package handlers

import (
	"BookShop/model"
	"BookShop/pkg"
	pubsub "BookShop/pub_sub"
	"sync"

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

func SubscribeBook(channel string, ch chan model.DetailBook, wg *sync.WaitGroup) {
	var bookSub model.DetailBook
	err := pubsub.Subscribe(&bookSub, channel)
	if err != nil {
		log.Errorf("Can't subscriber data, err: %v", err)
		return
	}
	ch <- bookSub
	defer wg.Done()
}

func PublishBook(bookPub model.CatPublish, channel string) error {
	channelDefault := viper.GetString("redis.channel_category")
	err := pubsub.Publish(bookPub, channelDefault)
	if err != nil {
		return err
	}
	return nil
}

func ListenPubSub(bookPub model.CatPublish) (data model.DetailBook, err error) {
	// subscribe result from redis
	channel := pkg.GenID()
	bookPub.Channel = channel
	var wg sync.WaitGroup
	wg.Add(1)
	ch := make(chan model.DetailBook)
	go SubscribeBook(channel, ch, &wg)

	// publish message to redis
	err = PublishBook(bookPub, channel)

	if err != nil {
		log.Errorf("Can't publisher data, err: %v", err)
		return data, err
	}
	select {
	case data = <-ch:
		close(ch)
		wg.Wait()
		return data, nil
	}
}
