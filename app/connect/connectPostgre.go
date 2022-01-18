package connect

import (
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectPostgres() (*gorm.DB, error) {
	nameDB := viper.GetString("postgres.name")
	host := viper.GetString("postgres.host")
	port := viper.GetString("postgres.post")
	user := viper.GetString("postgres.user")
	password := viper.GetString("postgres.password")
	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + nameDB + " port=" + port + " sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
