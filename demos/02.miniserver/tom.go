package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Hello struct{}

func (h Hello) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	dat, err := ioutil.ReadFile(r.URL.Path[1:])
	if err != nil {
		fmt.Fprintln(w, "Something went wrong: ", err)
	} else {
		fmt.Fprint(w, string(dat))
	}

}

func main() {
	var h Hello
	http.ListenAndServe("localhost:8080", h)
}
