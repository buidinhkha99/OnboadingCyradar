package postgesdb

import (
	"BookShop/app_book/connect"
	"BookShop/app_book/model"
	"BookShop/app_book/pkg"
	"BookShop/app_book/pkg/cache"
	"encoding/json"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type Postgres struct {
	pg *gorm.DB
}

func (db *Postgres) Connect() error {
	postgres, err := connect.ConnectPostgres()
	if err != nil {
		return err
	}
	db.pg = postgres
	return nil
}

func (db *Postgres) CreateBook(book *model.Book) error {
	id := pkg.GenID()
	book.ID = id
	result := db.pg.Create(&book)
	if result.Error != nil {
		log.Error("Create book have error: %v", result.Error)
		return result.Error
	}
	return nil

}

func (db *Postgres) UpdateBook(book model.Book) error {
	result := db.pg.Save(&book)
	if result.Error != nil {
		log.Error("Update book have error: %v", result.Error)
	}
	return nil
}

func (db *Postgres) DeleteBook(book model.Book) error {
	result := db.pg.Delete(&book)
	if result.Error != nil {
		log.Error("Delete book have error: %v", result.Error)
		return result.Error
	}
	return nil

}

func (db *Postgres) GetAllBook() ([]model.Book, error) {
	var books []model.Book
	result := db.pg.Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}
	return books, nil
}

func (db *Postgres) GetBook(ID string) (book model.Book, err error) {
	db.pg.Where("id = ?", ID).Find(&book)
	return book, nil
}

func (db *Postgres) GetTopBook() string {
	key := viper.GetString("redis.key")
	data := cache.ServeJQueryWithRemoteCache(key)
	if data == "" {
		var books []model.Book
		db.pg.Limit(9).Order("rate desc").Find(&books)
		b, _ := json.Marshal(books)
		err := cache.InsertData(key, string(b))
		if err != nil {
			log.Error("Error when add data in remote cache")
		}
		return string(b)
	}
	return data
}
func (db *Postgres) GetGroupBookByID(id string) (groupBook []model.GroupBook, err error) {
	db.pg.Where("book_id = ?", id).Find(&groupBook)
	return groupBook, nil
}

// handler Catergory

func (db *Postgres) CreateCategory(category model.Category) error {
	id := pkg.GenID()
	category.ID = id
	result := db.pg.Create(&category)
	if result.Error != nil {
		log.Error("Create category have error: %v", result.Error)
		return result.Error
	}
	return nil
}

func (db *Postgres) UpdateCategory(category model.Category) error {
	result := db.pg.Save(&category)
	if result.Error != nil {
		log.Error("Create book have error: %v", result.Error)
		return result.Error
	}
	return nil
}

func (db *Postgres) DeleteCategory(category model.Category) error {
	result := db.pg.Delete(&category)
	if result.Error != nil {
		log.Error("Delete category have error: %v", result.Error)
		return result.Error
	}
	return nil
}

func (db *Postgres) GetAllCategory() ([]model.Category, error) {
	var category []model.Category
	result := db.pg.Find(&category)
	if result.Error != nil {
		return nil, result.Error
	}
	return category, nil
}

func (db *Postgres) GetCategory(ID string) (category model.Category) {
	db.pg.Where("id = ?", ID).Find(&category)
	return category
}

func (db *Postgres) UpdateGroupBook(groupBook model.GroupBook) error {
	result := db.pg.Save(&groupBook)
	if result.Error != nil {
		log.Error("Create book have error: %v", result.Error)
	}
	return nil
}

func (db *Postgres) CreateGoupBook(groupBook model.GroupBook) error {
	id := pkg.GenID()
	groupBook.ID = id
	result := db.pg.Create(&groupBook)
	if result.Error != nil {
		log.Error("Create book have error: %v", result.Error)
	}
	return nil
}
