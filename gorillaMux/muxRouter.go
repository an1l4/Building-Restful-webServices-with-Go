package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func ArticleHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "category is %v\n", vars["category"])
	fmt.Fprintf(w, "id is %v\n", vars["id"])

}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/article/{category}/{id:[0-9]+}", ArticleHandler)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("server running at 8000")
	log.Fatal(srv.ListenAndServe())

}
