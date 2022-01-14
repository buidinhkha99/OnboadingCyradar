package model

import (
	"SavingBooks/app/connect"
	"SavingBooks/app/pkg/cache"
	"encoding/json"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Book struct {
	ID          uint64  `json:"id,omitempty" `
	Name        string  `json:"name,omitempty" `
	Quantily    int     `json:"quantily,omitempty" `
	Description string  `json:"description,omitempty" `
	Price       float32 `json:"price,omitempty" `
	Rate        float32 `json:"rate,omitempty" `
	Image       string  `json:"image,omitempty" `
}

type DetailBook struct {
	Book       Book        `json:"book"`
	Category   []Category  `json:"category"`
	GroupBooks []GroupBook `json:"group"`
}
type HandlerMysql interface {
	Createl()
	Update()
	Delete()
	GetAllBook() ([]Book, error)
	Get()
	GetTop()
	GetAllCategory() ([]Category, error)
}
type HandlerMongodb interface {
	Createl()
	Update()
	Delete()
	GetAllBook() ([]Book, error)
	Get()
	GetTop()
	GetAllCategory() ([]Category, error)
}

func (book Book) Create() {
	db := connect.ConnectMysql()
	result := db.Create(&book)
	if result.Error != nil {
		log.Error("Create book have error: %v", result.Error)
	}

}

func (book Book) Update() {
	db := connect.ConnectMysql()
	db.Save(&book)
}

func (book Book) Delete() {
	db := connect.ConnectMysql()
	result := db.Delete(&book)
	if result.Error != nil {
		log.Error("Delete book have error: %v", result.Error)
	}

}

func GetAllBook() ([]Book, error) {
	db := connect.ConnectMysql()
	var books []Book

	result := db.Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}
	return books, nil
}

func (book Book) GetMysql() {
	db := connect.ConnectMysql()
	db.Where("id = ?", book.ID).Find(&book)
}

func GetTop() string {
	key := viper.GetString("redis.key")
	data := cache.ServeJQueryWithRemoteCache(key)
	if data == "" {
		db := connect.ConnectMysql()
		var books []Book
		db.Limit(9).Order("rate desc").Find(&books)
		b, _ := json.Marshal(books)
		err := cache.InsertData(key, string(b))
		if err != nil {
			log.Error("Error when add data in remote cache")
		}
		return string(b)
	}
	return data
}
