package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"romanNumerals/coversion/romanNumerals"
	"strconv"
	"strings"
)

func RomanNumerals(w http.ResponseWriter, r *http.Request) {
	urlPathElement := strings.Split(r.URL.Path, "/")

	if urlPathElement[1] == "roman_number" {
		number, _ := strconv.Atoi(strings.TrimSpace(urlPathElement[2]))

		if number == 0 || number > 10 {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("404 - Not Found"))
		} else {
			fmt.Fprintf(w, "%q", html.EscapeString(romanNumerals.Numerals[number]))
		}

	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("400 - Bad Request"))
	}

}

func main() {

	http.HandleFunc("/", RomanNumerals)

	fmt.Println("server starting at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)

	}
}
