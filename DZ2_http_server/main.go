package main

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"log"
	"net/http"
	"tasks/DZ2_http_server/http_server"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	commentService := &http_server.CommentService{}
	commentCtrl := &http_server.CommentController{Service: commentService}

	commentCtrl.RegisterRoutes(r)

	log.Println("Сервер запущен на :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
