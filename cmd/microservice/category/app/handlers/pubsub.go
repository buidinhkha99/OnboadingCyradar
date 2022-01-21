package handlers

import (
	"BookShop/model"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/spf13/viper"
)

func GetBookHttp(id string, data *model.DetailBook) error {

	resp, err := http.Get("http://localhost:8888/book/" + id + "?filter=getbookcat")
	if err != nil {
		return err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return err
	}
	return nil

}

func GetDetailCategoryPubSub(idCategory string) (data model.DetailBook, err error) {
	checkType := viper.GetBool("http_pubsub.status")
	if !checkType {
		categoryPut := model.CatPublish{
			IdCatergory: idCategory,
			Description: "GetBooks",
		}
		data, err = ListenPubSub(categoryPut)
		if err != nil {
			return
		}
		return
	}
	err = GetBookHttp(idCategory, &data)
	if err != nil {

		return
	}
	return
}
