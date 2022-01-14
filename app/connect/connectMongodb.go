package connect

import (
	"context"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() (*mongo.Client, context.Context, context.CancelFunc, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	url := viper.GetString("mongodb.url")
	username := viper.GetString("mongodb.username")
	password := viper.GetString("mongodb.password")
	authSource := viper.GetString("mongodb.auth_source")
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://"+username+":"+password+"@"+url+"/config?authSource="+authSource))

	return client, ctx, cancel, err
}
func Close(client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {
	defer cancel()
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()
}

func CreateMonggodb() {

}
func UpdateMonggodb() {

}
func DeleteMongodb() {

}
func GetDataMongodb() {

}

func GetAllMongodb() {

}
