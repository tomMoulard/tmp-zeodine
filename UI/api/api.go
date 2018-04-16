package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

type jsonManage struct {
	err  error
	data map[string]interface{}
}

func (jsonS jsonManage) printCard(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	chaine := "card" + ps.ByName("id")

	if jsonS.err != nil {
		mapErr := map[string]string{"erreur": "1", "id": ps.ByName("id"), "img": "", "text": "", "card": chaine}
		res, _ := json.Marshal(mapErr)
		fmt.Fprintln(w, string(res))
		return
	}

	mapCard0 := jsonS.data[chaine]
	if mapCard0 == nil {
		mapErr := map[string]string{"erreur": "-1", "id": ps.ByName("id"), "img": "", "text": "", "card": chaine}
		res, _ := json.Marshal(mapErr)
		fmt.Fprintln(w, string(res))
		return
	}

	mapCard := mapCard0.(map[string]interface{})

	mapRes := map[string]string{"erreur": "0", "id": ps.ByName("id"), "img": mapCard["img"].(string), "text": mapCard["text"].(string), "card": chaine}
	res, _ := json.Marshal(mapRes)
	fmt.Fprintln(w, string(res))

}

type DbManager struct {
	id         int
	dataSource string
	db         *sql.DB
	err        error
}

func (dbm DbManager) setupDB() DbManager {
	log.Println("Initializing the database...")

	dbm.dataSource = "server:zeodine@tcp(db:3306)/zeodine"

	log.Printf("dataSource = %v\n", dbm.dataSource)

	// Connecting to db
	dbm.db, dbm.err = sql.Open("mysql", dbm.dataSource)
	if dbm.err != nil {
		panic(dbm.err.Error())
	}

	var pinged bool = false
	for i := 0; i < 3 && !pinged; i++ {
		log.Println("Connecting to database ... try:", i)
		err := dbm.db.Ping()
		if err != nil {
			log.Println(err.Error())
			time.Sleep(time.Second * 5)
			pinged = false
		} else {
			pinged = true
		}
	}

	if !pinged {
		log.Println("Could not connect to db in time")
		os.Exit(0)
	}
	log.Println("Connected")
	// Creating a new database
	// _, dbm.err = dbm.db.Exec("CREATE DATABASE IF NOT EXISTS zeodine")
	// _, dbm.err = dbm.db.Exec("CREATE DATABASE zeodine")
	// if dbm.err != nil {
	// log.Println(err)
	// 	log.Printf("Error when creating db: %v\n", dbm.err)
	// }
	// log.Println("Database created")

	// Using the fleshly created database:
	// _, dbm.err = dbm.db.Exec("USE zeodine")
	// if dbm.err != nil {
	// 	log.Println(dbm.err)
	// }

	// Creating a new table to insert some elements
	q := "CREATE TABLE IF NOT EXISTS zeodine.cards (card_id INT(16) NOT NULL AUTO_INCREMENT, name VARCHAR(64) DEFAULT NULL, img_url VARCHAR(128) DEFAULT NULL, description VARCHAR(200) DEFAULT NULL, pos_x INT(8) NOT NULL AUTO_INCREMENT, pos_y INT(8) NOT NULL AUTO_INCREMENT, ui JSON DEFAULT NULL, user_id INT(16) DEFAULT NULL, PRIMARY KEY (task_id, user_id))"
	_, dbm.err = dbm.db.Exec(q)
	if dbm.err != nil {
		log.Println("Error when creating table:", dbm.err)ll
	}
	log.Println("Table ready to be used")

	return dbm
}

func main() {
	var jsonS jsonManage
	var dbm DbManager

	dbm = dbm.setupDB()

	defer dbm.close()

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
	log.Fatal(http.ListenAndServe(":8080", router))
}
