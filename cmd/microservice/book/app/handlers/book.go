package handlers

import (
	"BookShop/manager"
	"BookShop/model"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var db = manager.DB

func ManagerFilter(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filter := vars["filter"]
	if filter == "top" {
		GetTopBooks(w, r)
		return
	}
	if len(filter) == 0 {
		GetAllProdcut(w, r)
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintf(w, "No required !")
}
func GetAllProdcut(w http.ResponseWriter, r *http.Request) {
	book, err := db.GetAllBook()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error: %v", err)
		return
	}

	data, _ := json.Marshal(book)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(data))
}

func GetDetailBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idBook := vars["id"]
	bookDetail := model.DetailBook{}

	b, err := db.GetBook(idBook)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	bookDetail.Book = b
	if len(bookDetail.Book.ID) == 0 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "No data")
		return
	}
	bookPub := model.BookPublish{
		IdBook:      bookDetail.Book.ID,
		Description: "GetGroupCategory",
	}
	data, err := ListenPubSub(bookPub)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	bookDetail.GroupBooks = data.GroupBooks
	bookDetail.Category = data.Category
	dataBook, _ := json.Marshal(&bookDetail)

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, string(dataBook))

}

func GetTopBooks(w http.ResponseWriter, r *http.Request) {

	data := db.GetTopBook()
	fmt.Fprint(w, data)
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
		return
	}
	for _, category := range book.Category {
		var group model.GroupBook
		group.BookID = book.Book.ID
		group.CategoryID = category.ID

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
	}
	fmt.Fprint(w, "Create book successfull")

}
func DeteleBook(w http.ResponseWriter, r *http.Request) {
	var book model.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	err = db.DeleteBook(book)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	fmt.Fprint(w, "Delete successfull !")
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var detailBook model.DetailBook

	err := json.NewDecoder(r.Body).Decode(&detailBook)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error: %v", err)
		return
	}

	err = db.UpdateBook(detailBook.Book)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	for _, groupBook := range detailBook.GroupBooks {
		dataPub := model.BookPublish{
			GroupBook:   groupBook,
			Description: "UpdateGroupBook",
		}
		dataSub, err := ListenPubSub(dataPub)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, err)
			return
		}
		if !dataSub.Status {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, "Wrong data !")
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Update successfull !")
}
