package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
)

// CustomServeMux is a struct which can be a multiplexer
type CustomServeMux struct {
}

// This is the function handler to be overridden
func (p *CustomServeMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		giveRandom(w, r)
		return

	}
	http.NotFound(w, r)
	return
}

func giveRandom(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "random number is : %f", rand.Float64())
}

func main() {

	fmt.Println("server is running at port 8000")

	// Any struct that has serveHTTP function can be a multiplexer
	mux := &CustomServeMux{}
	log.Fatal(http.ListenAndServe(":8000", mux))

}
