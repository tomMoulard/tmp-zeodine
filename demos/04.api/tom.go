package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strconv"
	"time"
)

//go get github.com/julienschmidt/httprouter

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func gettime(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprint(w, time.Now())
}

func reverse(str string) (res string) {
	for _, s := range str {
		res = string(s) + res
	}
	return
}

func getreverse(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprint(w, reverse(ps.ByName("str")))
}

func getold(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	i, err := strconv.Atoi(ps.ByName("year"))
	i = time.Now().Year() - i
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Fprint(w, i)
	}
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)
	router.GET("/gettime", gettime)
	router.GET("/getreverse/:str", getreverse)
	router.GET("/getold/:year", getold)

	log.Fatal(http.ListenAndServe(":8080", router))
}
