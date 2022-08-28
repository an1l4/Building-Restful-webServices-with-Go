package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func QueryHandler(w http.ResponseWriter, r *http.Request) {

	queryParams := r.URL.Query()

	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "got parameter id is :%s!\n", queryParams["id"][0])
	fmt.Fprintf(w, "got parameter category is : %s!\n", queryParams["category"][0])

}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/articles", QueryHandler)
	r.Queries("id", "category")

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("server running at port 8080")
	log.Fatal(srv.ListenAndServe())

}
