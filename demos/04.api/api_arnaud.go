package main

import (
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/golang/example/stringutil"
	"github.com/julienschmidt/httprouter"
)

func printTime(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	io.WriteString(w, "<p>"+time.Now().String()+"/<p>")
}

func printReverse(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	chaine := ps.ByName("ch")
	io.WriteString(w, "<p>"+stringutil.Reverse(chaine)+"</p>")
}

func printAge(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	age := ps.ByName("age")
	annee, _ := strconv.Atoi(age)
	if len(age) != 4 {
		io.WriteString(w, "<h1>Année de naissance incorrecte</h1>\n<p>Veiullez rentrer une année de naissance correcte</p>")
	} else {
		io.WriteString(w, "Vous avez "+strconv.Itoa(2018-annee)+" ans")
	}

}

func main() {
	router := httprouter.New()
	router.GET("/gettime", printTime)
	router.GET("/getreverse/:ch", printReverse)
	router.GET("/getold/:age", printAge)

	log.Fatal(http.ListenAndServe(":8080", router))
}
