package model

import (
	"SavingBooks/app/connect"

	log "github.com/sirupsen/logrus"
)

type Category struct {
	ID   uint64 `json:"id, omitempty"`
	Name string `json:"name, omitempty"`
}

func (category Category) CreateMysql() {
	db := connect.ConnectMysql()
	result := db.Create(&category)
	if result.Error != nil {
		log.Error("Create category have error: %v", result.Error)
	}
}

func (category Category) UpdateMysql() {

}

func (category Category) DeleteMysql() {
	db := connect.ConnectMysql()
	result := db.Delete(&category)
	if result.Error != nil {
		log.Error("Delete category have error: %v", result.Error)
	}
}

func GetAllCategoryMysql() ([]Category, error) {
	db := connect.ConnectMysql()
	var category []Category
	result := db.Find(&category)
	if result.Error != nil {
		return nil, result.Error
	}
	return category, nil
}
func (category Category) GetMysql() {
	db := connect.ConnectMysql()
	db.Where("id = ?", category.ID).Find(&category)
}
