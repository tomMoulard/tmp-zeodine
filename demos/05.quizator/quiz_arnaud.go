package main

import (
	"encoding/csv"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strconv"
)

func creatForm(q string, nb, godRep, totRep int) string {

	form := "<form action=\"\" method=\"post\">\nVotre réponse: <input type=\"text\" name=\"rep\">\n<input type=\"hidden\" name=\"nb\" value=\"" +
		strconv.Itoa(nb) +
		"\">\n<input type=\"hidden\" name=\"godRep\" value=\"" + strconv.Itoa(godRep) +
		"\">\n<input type=\"hidden\" name=\"totRep\" value=\"" + strconv.Itoa(totRep) +
		"\">\n<input type=\"submit\" value=\"Repondre\">\n</form>"
	return "<p>" + q + "</p>\n" + form
}

func analyseQuiz(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		printQuiz(w, r)
	} else {
		r.ParseForm()
		nb, _ := strconv.Atoi(r.Form["nb"][0])
		rep := r.Form["rep"][0]
		analyseRep(w, r, rep, nb)
	}
}

func analyseRep(w http.ResponseWriter, r *http.Request, rep string, nb int) {
	file, err := os.Open("questions.csv")

	if err != nil {
		io.WriteString(w, "Problème avec le fichier de quiz")
	} else {
		cvsFile := csv.NewReader(file)
		records, _ := cvsFile.ReadAll()
		goodRep, _ := strconv.Atoi(r.Form["godRep"][0])
		totRep, _ := strconv.Atoi(r.Form["totRep"][0])

		if records[nb][1] == rep {
			nb = rand.Intn(len(records))
			io.WriteString(w, "<h1 style=\"color:green;\">Bonne réponse !</h1>\n")
			io.WriteString(w, creatForm(records[nb][0], nb, goodRep+1, totRep+1))
			addRep(w, goodRep+1, totRep+1)
		} else {
			io.WriteString(w, "<h1 style=\"color:red;\">Mauvaise réponse : réesayer !</h1>\n")
			io.WriteString(w, creatForm(records[nb][0], nb, goodRep, totRep+1))
			addRep(w, goodRep, totRep+1)
		}

	}
	file.Close()
}

func addRep(w http.ResponseWriter, godRep, totRep int) {
	io.WriteString(w, "<p>"+strconv.Itoa(godRep)+"/"+strconv.Itoa(totRep)+"</p>\n")
}

func printQuiz(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("questions.csv")

	if err != nil {
		io.WriteString(w, "Problème avec le fichier de quiz")
	} else {
		var nb int
		cvsFile := csv.NewReader(file)
		records, _ := cvsFile.ReadAll()
		nb = rand.Intn(len(records))

		io.WriteString(w, creatForm(records[nb][0], nb, 0, 0))

	}

	addRep(w, 0, 0)
	file.Close()
}

func main() {
	http.HandleFunc("/", analyseQuiz)
	http.ListenAndServe(":8080", nil)
}
