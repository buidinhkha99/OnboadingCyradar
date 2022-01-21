package router

import (
	"BookShop/cmd/microservice/book/app/handlers"
	"BookShop/cmd/microservice/book/app/middlewares"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Run() {
	r := mux.NewRouter().StrictSlash(true)
	post := r.Methods(http.MethodPost).Subrouter()
	post.Path("/book").HandlerFunc(middlewares.SetMiddleware(handlers.CreateBooks))

	get := r.Methods(http.MethodGet).Subrouter()
	get.Path("/book").Queries("filter", "{filter}").HandlerFunc(middlewares.SetMiddleware(handlers.ManagerFilter))
	get.Path("/book/{id}").Queries("filter", "{filter}").HandlerFunc(middlewares.SetMiddleware(handlers.ManagerFilter))
	get.Path("/book/{id}").HandlerFunc(middlewares.SetMiddleware(handlers.GetDetailBook))
	delete := r.Methods(http.MethodDelete).Subrouter()
	delete.Path("/book").HandlerFunc(middlewares.SetMiddleware(handlers.DeteleBook))
	put := r.Methods(http.MethodPut).Subrouter()
	put.Path("/book").HandlerFunc(middlewares.SetMiddleware(handlers.UpdateBook))

	handler := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "DELETE", "PATCH", "OPTIONS", "PUT"},
		AllowedHeaders: []string{"*"},
	}).Handler(r)
	http.ListenAndServe(":8888", handler)
}
