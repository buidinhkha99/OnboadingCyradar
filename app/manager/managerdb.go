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
	CreateBook(book *model.Book) error
	UpdateBook(book model.Book) error
	DeleteBook(book model.Book) error
	GetAllBook() ([]model.Book, error)
	GetBook(ID int64) (book model.Book)
	GetTopBook() string
	GetAllCategory() ([]model.Category, error)
	CreateCategory(category model.Category) error
	UpdateCategory(category model.Category) error
	DeleteCategory(category model.Category) error
	GetCategory(ID int64) (category model.Category)
	UpdateGroupBook(groupBook model.GroupBook) error
	GetGroupBookByID(id int64) (groupBook []model.GroupBook)
	CreateGoupBook(groupBook model.GroupBook) error
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
