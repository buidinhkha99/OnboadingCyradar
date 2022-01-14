package manager

import (
	"BookShop/app/manager/mongodb.go"
	"BookShop/app/manager/mysqldb"
	"BookShop/app/model"
	"BookShop/app/pkg"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type ConnectionDB interface {
	Connect() error
	CreateBook(book model.Book) (err error)
	UpdateBook(book model.Book)
	DeleteBook(book model.Book)
	GetAllBook() ([]model.Book, error)
	GetBook(ID uint64) (book model.Book)
	GetTopBook() string
	GetAllCategory() ([]model.Category, error)
	CreateCategory(category model.Category)
	UpdateCategory(category model.Category)
	DeleteCategory(category model.Category)
	GetCategory(category model.Category)
}

var DB ConnectionDB

func init() {
	err := pkg.LoadConfig()
	if err != nil {
		log.Error("Error loading cofig ")

	}
	if viper.GetBool("mysql.status") {
		DB = &mysqldb.Mysql{}
	}
	if viper.GetBool("mongodb.status") {
		DB = &mongodb.Mongo{}
	}
	err = DB.Connect()
	if err != nil {
		panic(err)
	}
}
