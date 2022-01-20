package mysqldb

import (
	"BookShop/connect"
	"BookShop/model"
	"BookShop/pkg"
	"BookShop/pkg/cache"
	"encoding/json"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type Mysql struct {
	conn *gorm.DB // contains the connection to the DB
}

func (db *Mysql) Connect() error {

	mysql, err := connect.ConnectMysql()
	if err != nil {
		return err
	}
	db.conn = mysql
	return nil
}

func (db *Mysql) CreateBook(book *model.Book) error {
	id := pkg.GenID()
	book.ID = id
	result := db.conn.Create(&book)
	if result.Error != nil {
		log.Error("Create book have error: %v", result.Error)
		return result.Error
	}
	return nil

}

func (db *Mysql) UpdateBook(book model.Book) error {
	result := db.conn.Save(&book)
	if result.Error != nil {
		log.Error("Update book have error: %v", result.Error)
	}
	return nil
}

func (db *Mysql) DeleteBook(book model.Book) error {
	result := db.conn.Delete(&book)
	if result.Error != nil {
		log.Error("Delete book have error: %v", result.Error)
		return result.Error
	}
	return nil

}

func (db *Mysql) GetAllBook() ([]model.Book, error) {
	var books []model.Book
	result := db.conn.Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}
	return books, nil
}

func (db *Mysql) GetBook(ID string) (book model.Book, err error) {
	db.conn.Where("id = ?", ID).Find(&book)
	return book, nil
}

func (db *Mysql) GetTopBook() string {
	key := viper.GetString("redis.key")
	data := cache.ServeJQueryWithRemoteCache(key)
	if data == "" {
		var books []model.Book
		db.conn.Limit(9).Order("rate desc").Find(&books)
		b, _ := json.Marshal(books)
		err := cache.InsertData(key, string(b))
		if err != nil {
			log.Error("Error when add data in remote cache")
		}
		return string(b)
	}
	return data
}
func (db *Mysql) GetGroupBookByID(id string) (groupBook []model.GroupBook, err error) {
	db.conn.Where("book_id = ?", id).Find(&groupBook)
	return groupBook, nil
}

// handler Catergory

func (db *Mysql) CreateCategory(category model.Category) error {
	id := pkg.GenID()
	category.ID = id
	result := db.conn.Create(&category)
	if result.Error != nil {
		log.Error("Create category have error: %v", result.Error)
		return result.Error
	}
	return nil
}

func (db *Mysql) UpdateCategory(category model.Category) error {
	result := db.conn.Save(&category)
	if result.Error != nil {
		log.Error("Create book have error: %v", result.Error)
		return result.Error
	}
	return nil
}

func (db *Mysql) DeleteCategory(category model.Category) error {
	result := db.conn.Delete(&category)
	if result.Error != nil {
		log.Error("Delete category have error: %v", result.Error)
		return result.Error
	}
	return nil
}

func (db *Mysql) GetAllCategory() ([]model.Category, error) {
	var category []model.Category
	result := db.conn.Find(&category)
	if result.Error != nil {
		return nil, result.Error
	}
	return category, nil
}

func (db *Mysql) GetCategory(ID string) (category model.Category) {
	db.conn.Where("id = ?", ID).Find(&category)
	return category
}

func (db *Mysql) UpdateGroupBook(groupBook model.GroupBook) error {
	result := db.conn.Save(&groupBook)
	if result.Error != nil {
		log.Error("Create book have error: %v", result.Error)
	}
	return nil
}

func (db *Mysql) CreateGoupBook(groupBook model.GroupBook) error {
	id := pkg.GenID()
	groupBook.ID = id
	result := db.conn.Create(&groupBook)
	if result.Error != nil {
		log.Error("Create book have error: %v", result.Error)
	}
	return nil
}
