package handlers

import (
	"BookShop/app_book/model"
	"encoding/json"
	"fmt"
	"net/http"
)

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
