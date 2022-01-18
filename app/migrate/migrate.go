package migrate

import (
	"BookShop/app/connect"
	"BookShop/app/model"
)

func CreateTableMySql() {
	db, _ := connect.ConnectMysql()
	db.AutoMigrate(&model.Book{}, &model.Category{}, &model.GroupBook{})
}

func CreateTablePostgres() {
	db, _ := connect.ConnectPostgres()
	db.AutoMigrate(&model.Book{}, &model.Category{}, &model.GroupBook{})
}
