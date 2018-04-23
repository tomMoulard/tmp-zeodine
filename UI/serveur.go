package main

import (
	"app/api"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"text/template"

	"github.com/julienschmidt/httprouter"
)

type jsonManage struct {
	err    error
	data   map[string]interface{}
	nbcard int
}

func (jsonS jsonManage) printCard(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	chaine := "card" + ps.ByName("id")
	fmt.Println(chaine)

	if jsonS.err != nil {
		mapErr := map[string]string{"erreur": "1", "id": ps.ByName("id"), "img": "", "text": "", "card": chaine}
		res, err := json.Marshal(mapErr)
		if err != nil {
			fmt.Println("Erreur du Marshal ", err)
			return
		}
		fmt.Fprintln(w, string(res))
		return
	}

	mapCard0 := jsonS.data[chaine]
	if mapCard0 == nil {
		mapErr := map[string]string{"erreur": "-1", "id": ps.ByName("id"), "img": "", "text": "", "card": chaine}
		res, err := json.Marshal(mapErr)
		if err != nil {
			fmt.Println("Erreur du Marshal ", err)
			return
		}
		fmt.Fprintln(w, string(res))
		return
	}

	mapCard := mapCard0.(map[string]interface{})

	mapRes := map[string]string{"erreur": "0", "id": ps.ByName("id"), "img": mapCard["img"].(string), "text": mapCard["text"].(string), "card": chaine}
	res, err := json.Marshal(mapRes)
	if err != nil {
		fmt.Println("Erreur du Marshal ", err)
		return
	}
	fmt.Fprintln(w, string(res))

}

func (jsonS jsonManage) printNBCard(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	chaine := "card" + ps.ByName("id")
	fmt.Println(chaine)

	if jsonS.err != nil {
		mapErr := map[string]string{"erreur": "1", "nbcard": "-1"}
		res, err := json.Marshal(mapErr)
		if err != nil {
			fmt.Println("Erreur du Marshal ", err)
			return
		}
		fmt.Fprintln(w, string(res))
		return
	}

	mapRes := map[string]string{"erreur": "0", "nbcard": strconv.Itoa(jsonS.nbcard)}
	res, err := json.Marshal(mapRes)
	if err != nil {
		fmt.Println("Erreur du Marshal ", err)
		return
	}
	fmt.Fprintln(w, string(res))

}

func printCSSFile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Print("css : ")
	w.Header().Set("Content-Type: ", "text/css")
	cssFile := ps.ByName("css")

	fmt.Println(cssFile)

	file, err := ioutil.ReadFile("client/css/" + cssFile)
	if err != nil {
		fmt.Println("Erreur fichier")
		fmt.Fprintf(w, "<h1>Fichier non présent</h1>")
		return
	}
	fmt.Fprintf(w, string(file))

}

func printJSFile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type: ", "text/javascript")
	jsFile := ps.ByName("js")

	file, err := ioutil.ReadFile("client/js/" + jsFile)
	if err != nil {
		fmt.Fprintf(w, "<h1>Fichier non présent</h1>")
		return
	}
	fmt.Fprintf(w, string(file))

}

func printPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	user0 := ps.ByName("userid")
	ws0 := ps.ByName("wsid")

	vars := map[string]interface{}{
		"user": user0,
		"ws":   ws0,
	}

	t, _ := template.ParseFiles("client/index.html")

	t.Execute(w, vars)
}

func printPage2(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	user0 := ps.ByName("userid")

	vars := map[string]interface{}{
		"user": user0,
	}

	t, _ := template.ParseFiles("client/ws.html")

	t.Execute(w, vars)
}

func printPage3(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	user0 := ps.ByName("userid")

	vars := map[string]interface{}{
		"user": user0,
	}

	t, _ := template.ParseFiles("client/library.html")

	t.Execute(w, vars)
}

func printPage4(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	//user0 := ps.ByName("userid")

	vars := map[string]interface{}{
		"user": 1524134993,
	}

	t, _ := template.ParseFiles("client/index.html")

	t.Execute(w, vars)
}

func main() {
	var jsonS jsonManage

	file, err := os.Open("client/card.json")
	defer file.Close()

	if err != nil {
		fmt.Println("Erreur ouverture fichier")
		jsonS.err = err
	} else {
		info, _ := file.Stat()
		b := make([]byte, info.Size())
		n, err := file.Read(b)

		fmt.Println(n, " octets lus")
		if err != nil {
			fmt.Println("Erreur lecture fichier")
			jsonS.err = err
		}
		json.Unmarshal(b, &jsonS.data)
	}

	jsonS.nbcard = len(jsonS.data)
	fmt.Println(jsonS)

	router := httprouter.New()
	router.GET("/js/:js", printJSFile)
	router.GET("/css/:css", printCSSFile)

	router.GET("/", printPage4)
	router.GET("/mybiblio/:userid", printPage3)
	router.GET("/workspace/:userid", printPage2)
	router.GET("/workspace/:userid/:wsid", printPage)

	api.ExecuteAPI(router)

	log.Fatal(http.ListenAndServe(":8080", router))
}
