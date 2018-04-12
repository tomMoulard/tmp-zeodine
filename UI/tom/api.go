package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

type jsonManage struct {
	err  error
	data map[string]interface{}
}

func (jsonS jsonManage) printCard(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	chaine := "card" + ps.ByName("id")
	fmt.Println(jsonS.data)
	fmt.Println(chaine)

	if jsonS.err != nil {
		mapErr := map[string]string{"erreur": "1", "id": ps.ByName("ch"), "img": "", "text": "", "card": chaine}
		res, _ := json.Marshal(mapErr)
		fmt.Fprintln(w, string(res))
		return
	}

	mapCard := jsonS.data[chaine].(map[string]interface{})

	fmt.Println(mapCard)

	mapRes := map[string]string{"erreur": "0", "id": ps.ByName("ch"), "img": mapCard["img"].(string), "text": mapCard["text"].(string), "card": chaine}
	res, _ := json.Marshal(mapRes)
	fmt.Fprintln(w, string(res))

}

func main() {
	var jsonS jsonManage

	file, err := os.Open("card.json")
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

	fmt.Println(jsonS.data)

	router := httprouter.New()
	router.GET("/card/:id", jsonS.printCard)
	log.Fatal(http.ListenAndServe(":8081", router))
}
