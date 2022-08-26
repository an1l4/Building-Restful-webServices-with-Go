package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

//hello world web server

func Myserver(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World!")
}
func main() {

	fmt.Println("server running at port 8080")

	http.HandleFunc("/hello", Myserver)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
