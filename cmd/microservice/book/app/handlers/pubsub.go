package handlers

import (
	"BookShop/model"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/spf13/viper"
)

func GetBookHttp(ID string, data *model.DetailBook) error {
	resp, err := http.Get("http://localhost:8080/category/" + ID + "?filter=getbook")
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

func CreateGoupBookHttp(group model.GroupBook, data *model.DetailBook) error {
	dataJson, _ := json.Marshal(group)
	responseBody := bytes.NewBuffer(dataJson)
	resp, err := http.Post("http://localhost:8080/goupbook", "application/json", responseBody)
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

func UpdateBookHttp(group model.GroupBook) (model.DetailBook, error) {
	var bookSub model.DetailBook
	dataJson, _ := json.Marshal(group)
	resp, err := http.Post("http://localhost:8080/groupbookup", "application/json", bytes.NewBuffer(dataJson))
	if err != nil {
		return bookSub, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return bookSub, err
	}
	err = json.Unmarshal(body, &bookSub)
	if err != nil {
		return bookSub, err
	}
	return bookSub, nil
}
func UpdateBookPubSub(groupBook model.GroupBook) error {
	checkType := viper.GetBool("http_pubsub.status")
	if !checkType {
		dataPub := model.BookPublish{
			GroupBook:   groupBook,
			Description: "UpdateGroupBook",
		}
		dataSub, err := ListenPubSub(dataPub)
		if err != nil {
			return err
		}
		if !dataSub.Status {
			return err
		}
		return nil
	}
	bookSub, err := UpdateBookHttp(groupBook)
	if err != nil {
		return err
	}
	if !bookSub.Status {
		return errors.New("Can't update category in service Catergory")
	}
	return nil
}

func CreateBookPubSub(book model.DetailBook, category model.Category, w http.ResponseWriter, r *http.Request) {
	var group model.GroupBook
	group.BookID = book.Book.ID
	group.CategoryID = category.ID
	var bookSub model.DetailBook
	checkType := viper.GetBool("http_pubsub.status")
	if !checkType {
		data := model.BookPublish{
			GroupBook:   group,
			Description: "CreateGroupBook",
		}
		bookSub, err := ListenPubSub(data)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, err)
			return
		}
		if !bookSub.Status {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Wrong data !")
			return
		}
		return
	}
	err := CreateGoupBookHttp(group, &bookSub)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	if !bookSub.Status {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Wrong data !")
		return
	}
}

func GetDetailBookPubSub(bookDetail model.DetailBook, w http.ResponseWriter, r *http.Request) (model.DetailBook, error) {
	var data model.DetailBook
	checkType := viper.GetBool("http_pubsub.status")
	if !checkType {

		bookPub := model.BookPublish{
			IdBook:      bookDetail.Book.ID,
			Description: "GetGroupCategory",
		}
		data, err := ListenPubSub(bookPub)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, err)
			return data, err
		}
		return data, nil
	}
	err := GetBookHttp(bookDetail.Book.ID, &data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return data, nil
	}
	return data, nil
}
