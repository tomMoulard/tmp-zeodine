package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"
)

type Hello struct{}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles("index.html")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		fmt.Fprint(w, "Bonjour, ", r.Form["userName"][0], ".\n")
		i, err := strconv.Atoi(r.Form["birthYear"][0])
		i = time.Now().Year() - i
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Fprintln(w, "Vous avez", i, "ans.")
		}
	}

}

func main() {
	var h Hello
	http.ListenAndServe("localhost:8080", h)
}
