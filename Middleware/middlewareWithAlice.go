package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/justinas/alice"
)

type City struct {
	Name string
	Area uint64
}

// Middleware to check content type as JSON
func FileterContentType(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("currently in the check content middleware")

		if r.Header.Get("Content-type") != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			w.Write([]byte("415 - Unsupported Media Type. Please send JSON"))

			return

		}

		handler.ServeHTTP(w, r)
	})

}

func setServerTimeCookie(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)

		cookie := http.Cookie{Name: "Server-Time(UTC)", Value: strconv.FormatInt(time.Now().Unix(), 10)}
		http.SetCookie(w, &cookie)
		log.Println("currently in set server time middleware")
	})
}

func Logic(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var tempCity City
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&tempCity)

		if err != nil {
			panic(err)
		}
		defer r.Body.Close()

		log.Printf("Got %s city with area of %d sq miles\n", tempCity.Name, tempCity.Area)

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("201 - Created"))
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("405 - Method not allowed"))
	}
}

func main() {
	mainLogicHandler := http.HandlerFunc(Logic)
	chain := alice.New(FileterContentType, setServerTimeCookie).Then(mainLogicHandler)

	http.Handle("/city", chain)

	fmt.Println("server running at port 8000")
	http.ListenAndServe(":8000", nil)

}

