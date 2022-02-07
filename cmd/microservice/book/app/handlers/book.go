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
	filter := r.URL.Query().Get("filter")
	if filter == "top" {
		GetTopBooks(w, r)
		return
	}
	if filter == "getbookcat" {
		GetBookByCatergory(w, r)
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
	fmt.Print(1)
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

	// pub/sub data
	data, err := GetDetailBookPubSub(bookDetail, w, r)
	if err != nil {
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
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	err = db.CreateBook(&book.Book)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	for _, category := range book.Category {
		CreateBookPubSub(book, category, w, r)
	}
	dataBook, _ := json.Marshal(&book.Book)
	fmt.Fprintf(w, string(dataBook))

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
		if len(groupBook.ID) == 0 {
			err := db.CreateGoupBook(groupBook)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprint(w, err)
				return
			}
			continue
		}
		err = UpdateBookPubSub(groupBook)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, err)
			return
		}
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Update successfull !")
}

func GetBookByCatergory(w http.ResponseWriter, r *http.Request) {
	bookDetail := model.DetailBook{
		Status: true,
	}
	vars := mux.Vars(r)
	idCategory := vars["id"]
	book, err := db.GetBookWithCatergory(idCategory)
	if err != nil {
		bookDetail.Status = false
		dataBook, _ := json.Marshal(&bookDetail)
		fmt.Fprint(w, string(dataBook))
		return
	}
	bookDetail.Books = book
	dataBook, _ := json.Marshal(&bookDetail)
	fmt.Fprint(w, string(dataBook))
}
