package router

import (
	"BookShop/app_book/handlers"
	"BookShop/app_book/middlewares"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Run() {
	r := mux.NewRouter().StrictSlash(true)
	post := r.Methods(http.MethodPost).Subrouter()
	post.Path("/book").HandlerFunc(middlewares.SetMiddleware(handlers.CreateBooks))
	post.Path("/category").HandlerFunc(middlewares.SetMiddleware(handlers.CreateCategory))

	get := r.Methods(http.MethodGet).Subrouter()
	get.Path("/book").Queries("filter", "{filter}").HandlerFunc(middlewares.SetMiddleware(handlers.ManagerFilter))
	get.Path("/category").HandlerFunc(middlewares.SetMiddleware(handlers.GetCategory))
	// get.Path("/book").Queries("filter", "{filter}").HandlerFunc(middlewares.SetMiddleware(handlers.GetTopBooks))
	get.Path("/book/{id}").HandlerFunc(middlewares.SetMiddleware(handlers.GetDetailBook))
	delete := r.Methods(http.MethodDelete).Subrouter()
	delete.Path("/book").HandlerFunc(middlewares.SetMiddleware(handlers.DeteleBook))
	delete.Path("/category").HandlerFunc(middlewares.SetMiddleware(handlers.DeteleCategory))
	put := r.Methods(http.MethodPut).Subrouter()
	put.Path("/book").HandlerFunc(middlewares.SetMiddleware(handlers.UpdateBook))
	put.Path("/category").HandlerFunc(middlewares.SetMiddleware(handlers.UpdateCategory))

	handler := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "DELETE", "PATCH", "OPTIONS", "PUT"},
		AllowedHeaders: []string{"*"},
	}).Handler(r)
	http.ListenAndServe(":8888", handler)
}
