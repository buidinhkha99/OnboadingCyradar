package router

import (
	"BookShop/cmd/microservice/category/app/handlers"
	"BookShop/cmd/microservice/category/app/middlewares"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Run() {
	go handlers.CheckPubSub()
	fmt.Println(1)
	r := mux.NewRouter().StrictSlash(true)
	post := r.Methods(http.MethodPost).Subrouter()
	post.Path("/category").HandlerFunc(middlewares.SetMiddleware(handlers.CreateCategory))

	get := r.Methods(http.MethodGet).Subrouter()
	get.Path("/category").HandlerFunc(middlewares.SetMiddleware(handlers.GetCategory))
	// get.Path("/book").Queries("filter", "{filter}").HandlerFunc(middlewares.SetMiddleware(handlers.GetTopBooks))
	delete := r.Methods(http.MethodDelete).Subrouter()
	delete.Path("/category").HandlerFunc(middlewares.SetMiddleware(handlers.DeteleCategory))
	put := r.Methods(http.MethodPut).Subrouter()
	put.Path("/category").HandlerFunc(middlewares.SetMiddleware(handlers.UpdateCategory))

	handler := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "DELETE", "PATCH", "OPTIONS", "PUT"},
		AllowedHeaders: []string{"*"},
	}).Handler(r)
	http.ListenAndServe(":8080", handler)
}
