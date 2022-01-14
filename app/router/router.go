package router

import (
	"BookShop/app/handlers"
	"BookShop/app/middlewares"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Run() {
	r := mux.NewRouter().StrictSlash(true)
	// post := r.Methods(http.MethodPost).Subrouter()
	// post.Path("/book").HandlerFunc(middlewares.SetMiddleware(handlers.CreateBooks))
	// post.Path("/category").HandlerFunc(middlewares.SetMiddleware(handlers.CreateCategory))

	get := r.Methods(http.MethodGet).Subrouter()
	get.Path("/book").HandlerFunc(middlewares.SetMiddleware(handlers.GetAllProdcut))
	// get.Path("/category").HandlerFunc(middlewares.SetMiddleware(handlers.GetCategory))
	// get.Path("/book/{id}").HandlerFunc(middlewares.SetMiddleware(handlers.GetDetailBook))
	// get.Path("/book?filter=top").HandlerFunc(middlewares.SetMiddleware(handlers.GetTopBooks))
	// get.Path("/book?search={name}").HandlerFunc(middlewares.SetMiddleware(handlers.ElasticSearch))
	// delete := r.Methods(http.MethodDelete).Subrouter()
	// delete.Path("/book/{id}").HandlerFunc(middlewares.SetMiddleware(handlers.DeteleBook))
	// delete.Path("/category/{id}").HandlerFunc(middlewares.SetMiddleware(handlers.DeteleCategory))
	// put := r.Methods(http.MethodPut).Subrouter()
	// put.Path("/book/{id}").HandlerFunc(middlewares.SetMiddleware(handlers.UpdateBook))
	// put.Path("/category/{id}").HandlerFunc(middlewares.SetMiddleware(handlers.UpdateCategory))

	handler := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST", "DELETE", "PATCH", "OPTIONS", "PUT"},
		AllowedHeaders: []string{"*"},
	}).Handler(r)
	http.ListenAndServe(":8080", handler)
}

// GET POST
// GET POST PUT PATCH DELETE

// CRUD
// http://localhost:8080/book

// GET http://localhost:8080/book
// GET http://localhost:8080/book?filter=top
// GET http://localhost:8080/book?search=
// POST http://localhost:8080/book
// GET http://localhost:8080/book/{id}
// PUT/PATCH http://localhost:8080/book/{id}
// DELETE http://localhost:8080/book/{id}
