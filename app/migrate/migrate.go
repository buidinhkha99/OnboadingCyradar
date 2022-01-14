package migrate

import (
	"BookShop/app/connect"
	"BookShop/app/model"
)

func CreateTable() {
	db := connect.ConnectMysql()
	db.AutoMigrate(&model.Book{}, &model.Category{}, &model.GroupBook{})
}
