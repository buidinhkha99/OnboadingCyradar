package manager

import (
	"BookShop/manager/mongodb"
	"BookShop/manager/mysqldb"
	"BookShop/manager/postgesdb"
	"BookShop/model"
	"BookShop/pkg"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type ConnectionDB interface {
	Connect() error
	CreateBook(book *model.Book) error
	UpdateBook(book model.Book) error
	DeleteBook(book model.Book) error
	GetAllBook() ([]model.Book, error)
	GetBook(ID string) (book model.Book, err error)
	GetTopBook() string
	GetAllCategory() ([]model.Category, error)
	CreateCategory(category model.Category) error
	UpdateCategory(category model.Category) error
	DeleteCategory(category model.Category) error
	GetCategory(ID string) (category model.Category)
	UpdateGroupBook(groupBook model.GroupBook) error
	GetGroupBookByID(id string) (groupBook []model.GroupBook, err error)
	CreateGoupBook(groupBook model.GroupBook) error
	GetBookWithCatergory(ID string) ([]model.Book, error)
}

var DB ConnectionDB

func init() {
	err := pkg.LoadConfig()
	if err != nil {
		log.Error("Error loading cofig ")
		return

	}
	if viper.GetBool("mysql.status") {
		DB = &mysqldb.Mysql{}
	}
	if viper.GetBool("mongodb.status") {
		DB = &mongodb.Mongo{}
	}
	if viper.GetBool("postgres.status") {
		DB = &postgesdb.Postgres{}
	}
	err = DB.Connect()
	if err != nil {
		panic(err)
	}
}
