package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

func main() {

	newMux := http.NewServeMux()

	newMux.HandleFunc("/random_float", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "random number in :%f", rand.Float64())
	})

	newMux.HandleFunc("/random_int", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "random int number is : %d", rand.Intn(100))
	})

	fmt.Println("server running at port 8000")
	log.Fatal(http.ListenAndServe(":8000", newMux))

}
