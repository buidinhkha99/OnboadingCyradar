package mongodb

import (
	"BookShop/app/connect"
	"BookShop/app/model"
	"BookShop/app/pkg/cache"
	"context"
	"encoding/json"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
)

type Mongo struct {
	client *mongo.Client
	ctx    context.Context
	cancel context.CancelFunc
}

func (mg *Mongo) Connect() error {
	client, ctx, cancel, err := connect.Connect()
	if err != nil {
		log.Errorf("Cannot connect to mongoDB: %v", err)
		return err
	}
	mg.client = client
	mg.cancel = cancel
	mg.ctx = ctx
	return nil
}
func (mg *Mongo) CreateBook(book *model.Book) error {
	// result := mg.client.Create(&book)
	// if result.Error != nil {
	// 	log.Error("Create book have error: %v", result.Error)
	// }
	// defer connect.Close(mg.client, mg.ctx, mg.cancel)
	return nil

}

func (mg *Mongo) UpdateBook(book model.Book) error {
	// mg.client.Save(&book)
	// defer connect.Close(mg.client, mg.ctx, mg.cancel)
	return nil
}

func (mg *Mongo) DeleteBook(book model.Book) error {
	// result := mg.client.Delete(&book)
	// if result.Error != nil {
	// 	log.Error("Delete book have error: %v", result.Error)
	// }
	// defer connect.Close(mg.client, mg.ctx, mg.cancel)
	return nil

}

func (mg *Mongo) GetAllBook() ([]model.Book, error) {
	var books []model.Book
	// result := mg.client.Find(&books)
	// if result.Error != nil {
	// 	return nil, result.Error
	// }
	// return books, nil
	// defer connect.Close(mg.client, mg.ctx, mg.cancel)
	return books, nil
}

func (mg *Mongo) GetBook(ID int64) (book model.Book) {
	// mg.client.Where("id = ?", ID).Find(&book)
	return book
}

func (mg *Mongo) GetTopBook() string {
	key := viper.GetString("redis.key")
	data := cache.ServeJQueryWithRemoteCache(key)
	if data == "" {
		mg := connect.ConnectMysql()
		var books []model.Book
		mg.Limit(9).Order("rate desc").Find(&books)
		b, _ := json.Marshal(books)
		err := cache.InsertData(key, string(b))
		if err != nil {
			log.Error("Error when add data in remote cache")
		}
		return string(b)
	}
	return data
}

func (db *Mongo) GetGroupBookByID(id int64) (groupBook []model.GroupBook) {
	// db.conn.Where("book_id = ?", id).Find(&groupBook)
	return groupBook
}

// handler Catergory

func (mg *Mongo) CreateCategory(category model.Category) error {
	// result := mg.client.Create(&category)
	// if result.Error != nil {
	// 	log.Error("Create category have error: %v", result.Error)
	// }
	return nil
}

func (mg *Mongo) UpdateCategory(category model.Category) error {
	// mg.client.Save(&category)
	return nil
}

func (mg *Mongo) DeleteCategory(category model.Category) error {
	// result := mg.client.Delete(&category)
	// if result.Error != nil {
	// 	log.Error("Delete category have error: %v", result.Error)
	// }
	return nil
}

func (mg *Mongo) GetAllCategory() ([]model.Category, error) {
	var category []model.Category
	// result := mg.client.Find(&category)
	// if result.Error != nil {
	// 	return nil, result.Error
	// }
	return category, nil
}
func (db *Mongo) GetCategory(ID int64) (category model.Category) {
	// db.conn.Where("id = ?", ID).Find(&category)
	return category
}

func (db *Mongo) UpdateGroupBook(groupBook model.GroupBook) error {
	// db.conn.Save(&groupBook)
	return nil
}

func (db *Mongo) CreateGoupBook(groupBook model.GroupBook) error {
	// result := db.conn.Create(&groupBook)
	// if result.Error != nil {
	// 	log.Error("Create book have error: %v", result.Error)
	// }
	return nil
}
