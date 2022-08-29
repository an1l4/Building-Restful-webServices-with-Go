package main

import (
	"fmt"
	"net/http"
)

func middleware(handler http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("executing middleware before request phase")
		handler.ServeHTTP(w, r)
		fmt.Println("executing middleware after response phase")

	})

}

func mainLogic(w http.ResponseWriter, r *http.Request) {
	fmt.Println("executing main logic")
	w.Write([]byte("OK"))
}

func main() {
	mainLogicHandler := http.HandlerFunc(mainLogic)
	http.Handle("/", middleware(mainLogicHandler))
	fmt.Println("server running at port 8080")
	http.ListenAndServe(":8000", nil)
}
