package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

type city struct {
	Name string
	Area uint64
}

// Middleware to check content type as JSON

func FileterContent(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("currently in the check content middleware")

		if r.Header.Get("Content-type") != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			w.Write([]byte("415 - Unsupported Media Type. Please send JSON"))
		}
		handler.ServeHTTP(w, r)

	})

}

// Middleware to add server timestamp for response cookie

func SetServerTimeCookie(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)

		cookie := http.Cookie{Name: "Server-Time(UTC)", Value: strconv.FormatInt(time.Now().Unix(), 10)}
		http.SetCookie(w, &cookie)

		log.Println("currently in the set server time middleware")
	})

}

func MainLogic(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var tempCity city
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&tempCity)

		if err != nil {
			panic(err)
		}
		defer r.Body.Close()

		log.Printf("Got %s city with area %d sq miles", tempCity.Name, tempCity.Area)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("201 - Created"))
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("405 - Method not allowed"))
	}

}

func main() {
	mainLogicHandler := http.HandlerFunc(MainLogic)
	http.Handle("/city", FileterContent(SetServerTimeCookie(mainLogicHandler)))

	fmt.Println("server running at port 8000")
	http.ListenAndServe(":8000", nil)
}
