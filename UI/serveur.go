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

type serv struct {
	path      string
	mime_type string
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
	fmt.Fprintln(w, res)

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

func (s serv) printFile(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type: ", s.mime_type)

	info := ps.ByName("info")

	fmt.Println(s.path + info)

	file, err := ioutil.ReadFile(s.path + info)
	if err != nil {
		w.WriteHeader(404)
		fmt.Println("Erreur fichier")
		fmt.Fprintf(w, "<h1>Error 404 : page not found</h1>")
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, string(file))
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

	js := serv{path: "client/js/", mime_type: "application/javascript"}
	css := serv{path: "client/css/", mime_type: "text/css"}
	imgP := serv{path: "client/assets/productif/", mime_type: "image/png"}
	imgS := serv{path: "client/assets/souverain/", mime_type: "image/png"}
	imgG := serv{path: "client/assets/guerrier/", mime_type: "image/png"}

	router := httprouter.New()
	router.GET("/js/:info", js.printFile)
	router.GET("/css/:info", css.printFile)

	router.GET("/guerrier/:info", imgG.printFile)
	router.GET("/souverain/:info", imgS.printFile)
	router.GET("/productif/:info", imgP.printFile)

	router.GET("/", printPage4)
	router.GET("/cards/:id", jsonS.printCard)
	router.GET("/mybiblio/:userid", printPage3)
	router.GET("/workspace/:userid", printPage2)
	router.GET("/workspace/:userid/:wsid", printPage)

	api.ExecuteAPI(router)

	log.Fatal(http.ListenAndServe(":8080", router))
}
