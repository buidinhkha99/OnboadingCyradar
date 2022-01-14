package manager

import (
	"SavingBooks/app/model"
	"SavingBooks/app/pkg"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var database string

func init() {
	err := pkg.LoadConfig()
	if err != nil {
		log.Error("Error loading cofig")
	}

	if viper.GetBool("mysql.status") {
		database = "mysql"
	}

	if viper.GetBool("mongodb.status") {
		database = "mongodb"
	}
}

func ManagerDatabase() interface{} {
	if database == "mysql" {
		return model.HandlerMysql
	}
	return model.HandlerMongodb
}
