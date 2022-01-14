package mysqldb

import (
	"BookShop/app/connect"
	"BookShop/app/model"
	"BookShop/app/pkg/cache"
	"encoding/json"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

// var database string

// func init() {
// 	err := pkg.LoadConfig()
// 	if err != nil {
// 		log.Error("Error loading cofig")
// 	}

// 	if viper.GetBool("mysql.status") {
// 		database = "mysql"
// 	}

// 	if viper.GetBool("mongodb.status") {
// 		database = "mongodb"
// 	}
// }

type Mysql struct {
	conn *gorm.DB // contains the connection to the DB
}

func (db *Mysql) Connect() error {

	db.conn = connect.ConnectMysql()
	return nil
}
func (db *Mysql) CreateBook(book model.Book) (err error) {
	result := db.conn.Create(&book)
	if result.Error != nil {
		log.Error("Create book have error: %v", result.Error)
	}
	return nil

}

func (db *Mysql) UpdateBook(book model.Book) {
	db.conn.Save(&book)
}

func (db *Mysql) DeleteBook(book model.Book) {
	result := db.conn.Delete(&book)
	if result.Error != nil {
		log.Error("Delete book have error: %v", result.Error)
	}

}

func (db *Mysql) GetAllBook() ([]model.Book, error) {
	var books []model.Book
	result := db.conn.Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}
	return books, nil
}

func (db *Mysql) GetBook(ID uint64) (book model.Book) {
	db.conn.Where("id = ?", ID).Find(&book)
	return book
}

func (db *Mysql) GetTopBook() string {
	key := viper.GetString("redis.key")
	data := cache.ServeJQueryWithRemoteCache(key)
	if data == "" {
		db := connect.ConnectMysql()
		var books []model.Book
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

// handler Catergory

func (db *Mysql) CreateCategory(category model.Category) {
	result := db.conn.Create(&category)
	if result.Error != nil {
		log.Error("Create category have error: %v", result.Error)
	}
}

func (db *Mysql) UpdateCategory(category model.Category) {
	db.conn.Save(&category)
}

func (db *Mysql) DeleteCategory(category model.Category) {
	result := db.conn.Delete(&category)
	if result.Error != nil {
		log.Error("Delete category have error: %v", result.Error)
	}
}

func (db *Mysql) GetAllCategory() ([]model.Category, error) {
	var category []model.Category
	result := db.conn.Find(&category)
	if result.Error != nil {
		return nil, result.Error
	}
	return category, nil
}
func (db *Mysql) GetCategory(category model.Category) {
	db.conn.Where("id = ?", category.ID).Find(&category)
}
