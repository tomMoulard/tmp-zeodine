package main

import (
	"io"
	"net/http"
)

func printPage(w http.ResponseWriter, r *http.Request) {
	//fmt.Println(r.URL)
	switch {
	case "/" == r.URL.Path:
		io.WriteString(w, "Bonjour le monde")
	default:
		io.WriteString(w, r.URL.Path)
	}

}

func main() {
	http.HandleFunc("/", printPage)

	http.ListenAndServe(":8080", nil)
}
