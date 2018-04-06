package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter" //go get github.com/julienschmidt/httprouter
	"html/template"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
)

func parseData(s string) [][2]string {
	lines := strings.Split(s, "\n") //split lines
	res := make([][2]string, len(lines))
	for i, _ := range res { //foreach lines split comas
		a := strings.Split(lines[i], ",")
		if len(a) != 2 {
			log.Fatal("CSV file not right on line :", i)
		} else {
			res[i] = [2]string{a[0], a[1]}
		}
	}
	return res
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//Choosing a question
	pos := rand.Intn(len(parsed))
	//Adding the file template
	t, err := template.New("index.html").ParseFiles("index.html")
	if err != nil {
		panic(err)
	}
	//Replacing parameters directly into the template
	var i int
	if right == 0 && wrong == 0 {
		i = 0
	} else {
		i = right / (wrong + right) * 100
	}
	q := map[string]string{
		"Question": parsed[pos][0],
		"Pos":      strconv.Itoa(pos),
		"Right":    strconv.Itoa(right),
		"Wrong":    strconv.Itoa(wrong),
		"Percent":  strconv.Itoa(i),
	}
	//Sending the template to the client
	err = t.Execute(w, q)
	if err != nil {
		panic(err)
	}
	last = pos
}

func answer(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//getting the form answer
	r.ParseForm()
	answer := r.Form["answer"][0]
	if answer == parsed[last][1] { //good answer
		right += 1
		page := "<head></head><body><p>Congrats, You are right !</p><p><a href=\"http://localhost:8080/\">Back Home</a></p></body>"
		t, err := template.New("Test").Parse(page)
		if err != nil {
			panic(err)
		}
		err = t.Execute(w, nil)
		if err != nil {
			panic(err)
		}
	} else { //wrong answer
		wrong += 1
		fmt.Fprintln(w, "You wrote :", answer, "But the answer was this:", parsed[last][1])
	}
}

var last int = 0
var right int = 0
var wrong int = 0
var parsed [][2]string

func main() {
	//Reading the question.csv
	data, err := ioutil.ReadFile("questions.csv")
	if err != nil {
		panic(err)
	}
	//Parsing the file
	parsed = parseData(string(data))
	//Creating the router
	router := httprouter.New()
	//routing
	router.GET("/", Index)
	router.POST("/answer/:pos", answer)

	log.Fatal(http.ListenAndServe(":8080", router))
}
