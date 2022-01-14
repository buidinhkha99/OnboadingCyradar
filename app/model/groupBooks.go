package model

import (
	"BookShop/app/connect"

	log "github.com/sirupsen/logrus"
)

type GroupBook struct {
	CategoryID uint64 `json:"category_id,omitempty"`
	BookID     uint64 `json:"book_id,omitempty"`
}

func (groupBook GroupBook) CreateMysql() {
	db := connect.ConnectMysql()
	result := db.Create(&groupBook)
	if result.Error != nil {
		log.Error("Create groupBook have error: %v", result.Error)
	}
}

func (groupBook GroupBook) UpdateMysql() {

}

func (groupBook GroupBook) DeleteMysql() {
	db := connect.ConnectMysql()
	result := db.Delete(&groupBook)
	if result.Error != nil {
		log.Error("Delete groupBook have error: %v", result.Error)
	}
}
