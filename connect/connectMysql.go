package connect

import (
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectMysql() (*gorm.DB, error) {
	nameDB := viper.GetString("mysql.name")
	host := viper.GetString("mysql.host")
	port := viper.GetString("mysql.post")
	password := viper.GetString("mysql.password")
	dsn := "root:" + password + "@tcp(" + host + ":" + port + ")/" + nameDB + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
