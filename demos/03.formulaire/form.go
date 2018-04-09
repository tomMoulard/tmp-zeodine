package main

import (
	"io"
	"net/http"
	"os"
	"strconv"
)

func printPage(w http.ResponseWriter, r *http.Request) {
	//fmt.Println(r)
	indx, _ := os.Open("index.html")
	if r.Method == "GET" {
		io.Copy(w, indx)
	} else {
		r.ParseForm()
		io.WriteString(w, "Bonjour, "+r.Form["userName"][0]+".\n")
		i, err := strconv.Atoi(r.Form["birthYear"][0])
		i = 2018 - i
		if err != nil {
			io.WriteString(w, "Erreur lors de la saisie de l'age")
		} else {
			io.WriteString(w, "Vous avez "+string(i)+" ans.")
		}
	}

}

func main() {
	http.HandleFunc("/", printPage)
	http.ListenAndServe(":8080", nil)
}
