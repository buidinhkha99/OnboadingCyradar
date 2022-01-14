package handlers

import (
	"BookShop/app/manager"
	"BookShop/app/model"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var db = manager.DB

func GetAllProdcut(w http.ResponseWriter, r *http.Request) {
	book, err := db.GetAllBook()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error: %v", err)
	}
	data, _ := json.Marshal(book)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(data))
}

func GetDetailBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idBook, _ := strconv.ParseInt(vars["id"], 10, 64)
	bookDetail := model.DetailBook{}

	bookDetail.Book = db.GetBook(idBook)
	bookDetail.GroupBooks = db.GetGroupBookByID(idBook)
	for _, group := range bookDetail.GroupBooks {
		bookDetail.Category = append(bookDetail.Category, db.GetCategory(int64(group.CategoryID)))
	}
	dataBook, _ := json.Marshal(&bookDetail)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(dataBook))

}

func GetTopBooks(w http.ResponseWriter, r *http.Request) {
	// key := viper.GetString("redis.key")
	// if data := cache.ServeJQueryWithRemoteCache(w, key); data == "" {
	// 	db := connect.ConnectMysql()
	// 	var books []model.Books
	// 	db.Limit(9).Order("rate desc").Find(&books)
	// 	b, _ := json.Marshal(books)
	// 	err := cache.InsertData(key, string(b))
	// 	if err != nil {
	// 		log.Error("Error when add data in remote cache")
	// 	}
	// 	fmt.Fprintln(w, string(b))

	// } else {
	// 	fmt.Fprintln(w, data)
	// }

}

func CreateBooks(w http.ResponseWriter, r *http.Request) {
	var book model.DetailBook
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		fmt.Fprint(w, err)
		return
	}
	err = db.CreateBook(&book.Book)
	if err != nil {
		fmt.Fprint(w, err)
	}
	for _, category := range book.Category {
		var group model.GroupBook
		group.BookID = book.Book.ID
		group.CategoryID = category.ID
		err = db.CreateGoupBook(group)
	}
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Wrong data !")
	}
	fmt.Fprint(w, "Create book successfull")

}
func DeteleBook(w http.ResponseWriter, r *http.Request) {
	// db := connect.ConnectMysql()
	// var book model.Books
	// var image model.Images
	// vars := mux.Vars(r)
	// idBook, err := strconv.ParseInt(vars["id"], 10, 64)
	// if err != nil {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	fmt.Fprintf(w, "Error: %v", err)
	// }
	// db.Delete(&book, idBook)
	// db.Where("book_id =?", idBook).Delete(&image)
	// bookElt := model.BookELT{
	// 	ID: int(idBook),
	// }
	// // delete in Elasticsearch
	// err = bookElt.DeleteDoc()
	// if err != nil {
	// 	log.Error("Error add data in Elasticsearch, ID: ", idBook)
	// }
	// w.WriteHeader(http.StatusOK)
	// fmt.Fprint(w, "Delete successfull !")
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var detailBook model.DetailBook

	err := json.NewDecoder(r.Body).Decode(&detailBook)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error: %v", err)
	}
	db.UpdateBook(detailBook.Book)
	for _, groupBook := range detailBook.GroupBooks {
		db.UpdateGroupBook(groupBook)
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Update successfull !")
}
