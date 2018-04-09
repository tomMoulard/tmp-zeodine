package main

import (
	"io"
	"net/http"
	"os"
)

func printPage(w http.ResponseWriter, r *http.Request) {
	//fmt.Println(r.URL.Path)
	file, err := os.Open("www" + r.URL.Path)
	if err != nil {
		io.WriteString(w, "Fichier non présent")
		return
	}
	io.Copy(w, file)

}

func main() {
	http.HandleFunc("/", printPage)

	http.ListenAndServe(":8080", nil)
}
