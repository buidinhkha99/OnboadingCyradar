package handlers

import (
	"BookShop/model"
	"BookShop/pkg"
	"fmt"
	"sync"

	pubsub "BookShop/pub_sub"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func SubscribeBook(channel string, ch chan model.DetailBook, wg *sync.WaitGroup) {
	var bookSub model.DetailBook
	err := pubsub.Subscribe(&bookSub, channel)
	fmt.Println(2)
	if err != nil {
		log.Errorf("Can't subscriber data, err: %v", err)
		return
	}
	ch <- bookSub
	defer wg.Done()
}

func PublishBook(bookPub model.BookPublish, channel string) error {
	channelDefault := viper.GetString("redis.channel_book")
	err := pubsub.Publish(bookPub, channelDefault)
	if err != nil {
		return err
	}
	return nil
}

func ListenPubSub(bookPub model.BookPublish) (data model.DetailBook, err error) {
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

func CheckPubSub() {
	ch := make(chan model.CatPublish, 5)
	go CheckMessBook(ch)
	for data := range ch {
		var bookDetail model.DetailBook
		switch data.Description {
		case "GetBooks":

			// get data in database
			book, _ := db.GetBookWithCatergory(data.IdCatergory)
			// subscribe data
			bookDetail.Books = book
			err := pubsub.Publish(bookDetail, data.Channel)
			if err != nil {
				log.Error("Can't publisher data, err: %v", err)
				continue
			}
		}

	}
}

func CheckMessBook(ch chan model.CatPublish) {
	var catSub model.CatPublish
	channel := viper.GetString("redis.channel_category")
	for {
		err := pubsub.Subscribe(&catSub, channel)
		if err != nil {
			log.Errorf("Can't not get data from book %v", err.Error())
		}
		ch <- catSub
	}

}
