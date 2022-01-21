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
	if filter == "getbook" {
		GetCategoryByBook(w, r)
		return
	}
	if len(filter) == 0 {
		GetDetailCategory(w, r)
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprintf(w, "No required !")
}

func GetCategory(w http.ResponseWriter, r *http.Request) {
	category, err := db.GetAllCategory()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error: %v", err)
		return
	}

	dataCategory, _ := json.Marshal(&category)
	fmt.Fprint(w, string(dataCategory))
}

func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	var catergory model.Category
	err := json.NewDecoder(r.Body).Decode(&catergory)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}

	if len(catergory.Name) != 0 {
		err = db.UpdateCategory(catergory)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprint(w, err)
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Update successfull !")
		return
	}

	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(w, "Error: Invalid name")

}
func CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category model.Category
	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Create Images of book have error: %v", err)
		return
	}

	err = db.CreateCategory(category)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Create Catergory  have error: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Create catergory successfull")

}
func DeteleCategory(w http.ResponseWriter, r *http.Request) {
	var category model.Category

	err := json.NewDecoder(r.Body).Decode(&category)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Delete Catergory  have error: %v", err)
		return
	}

	err = db.DeleteCategory(category)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Update Catergory  have error: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Delete successfull !")
}

func GetDetailCategory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idCategory := vars["id"]
	data, err := GetDetailCategoryPubSub(idCategory)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	dataCategory, _ := json.Marshal(&data.Books)
	fmt.Fprint(w, string(dataCategory))
}
func GetCategoryByBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idBook := vars["id"]
	var bookDetail model.DetailBook
	bookDetail.GroupBooks, _ = db.GetGroupBookByID(idBook)
	for _, group := range bookDetail.GroupBooks {
		bookDetail.Category = append(bookDetail.Category, db.GetCategory(group.CategoryID))
	}
	dataBook, _ := json.Marshal(&bookDetail)
	fmt.Fprint(w, string(dataBook))
}
func CreateGoupBook(w http.ResponseWriter, r *http.Request) {
	var group model.GroupBook
	bookDetail := model.DetailBook{
		Status: true,
	}
	err := json.NewDecoder(r.Body).Decode(&group)
	if err != nil {
		bookDetail.Status = false
		dataBook, _ := json.Marshal(&bookDetail)
		fmt.Fprint(w, string(dataBook))
		return
	}
	err = db.CreateGoupBook(group)
	if err != nil {
		bookDetail.Status = false
		dataBook, _ := json.Marshal(&bookDetail)
		fmt.Fprint(w, string(dataBook))
		return
	}
	dataBook, _ := json.Marshal(&bookDetail)
	fmt.Fprint(w, string(dataBook))
}

func UpGroupBook(w http.ResponseWriter, r *http.Request) {
	fmt.Print(2)
	var group model.GroupBook
	bookDetail := model.DetailBook{
		Status: true,
	}
	err := json.NewDecoder(r.Body).Decode(&group)
	if err != nil {
		bookDetail.Status = false
		dataBook, _ := json.Marshal(&bookDetail)
		fmt.Fprint(w, string(dataBook))
		return
	}
	err = db.UpdateGroupBook(group)
	if err != nil {
		bookDetail.Status = false
		dataBook, _ := json.Marshal(&bookDetail)
		fmt.Fprint(w, string(dataBook))
		return
	}
	dataBook, _ := json.Marshal(&bookDetail)
	fmt.Fprint(w, string(dataBook))
}
