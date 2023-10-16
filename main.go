package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println(os.Getenv("name"))
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("welcome"))
	})
	log.Fatal(http.ListenAndServe(":3000", router))
}
