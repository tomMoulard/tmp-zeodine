package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

type jsonManage struct {
	err    error
	data   map[string]interface{}
	nbcard int
}

type DbManager struct {
	id         int
	dataSource string
	db         *sql.DB
	err        error
}

// router.GET("/card/:userID/:id", dbm.card)
func (dbm DbManager) card(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	card_id := ps.ByName("id")
	user_id := ps.ByName("userID")
  
	que, err := dbm.db.Query("SELECT * FROM zeodine.cards WHERE card_id = " + card_id +" AND user_id = " + user_id)
	if err != nil {
		log.Println(err)
		// panic(err)
	}
	for que.Next() {
		var card_id int
		var name string
		var img_url string
		var description string
		var pos_x int
		var pos_y int
		var ui map[string]interface{}
		err = que.Scan(&id, &name, &img_url, &description, &pos_x, , &pos_y, &ui, &user_id)
		if err != nil {
			log.Println(err)
			// panic(err)
		}
		fmt.Fprintln(w, "{id:", id, ",name:", name, ",img_url:", img_url, ",description:", description, ",pos_x:", pos_x, ",pos_y:", pos_y, ",ui:", ui.(string), ",user_id:", user_id"}")
	}
}

// func (jsonS jsonManage) printCard(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	chaine := "card" + ps.ByName("id")
// 	fmt.Println(chaine)

// 	if jsonS.err != nil {
// 		mapErr := map[string]string{"erreur": "1", "id": ps.ByName("id"), "img": "", "text": "", "card": chaine}
// 		res, err := json.Marshal(mapErr)
// 		if err != nil {
// 			fmt.Println("Erreur du Marshal ", err)
// 			return
// 		}
// 		fmt.Fprintln(w, string(res))
// 		return
// 	}

// 	mapCard0 := jsonS.data[chaine]
// 	if mapCard0 == nil {
// 		mapErr := map[string]string{"erreur": "-1", "id": ps.ByName("id"), "img": "", "text": "", "card": chaine}
// 		res, err := json.Marshal(mapErr)
// 		if err != nil {
// 			fmt.Println("Erreur du Marshal ", err)
// 			return
// 		}
// 		fmt.Fprintln(w, string(res))
// 		return
// 	}

// 	mapCard := mapCard0.(map[string]interface{})

// 	mapRes := map[string]string{"erreur": "0", "id": ps.ByName("id"), "img": mapCard["img"].(string), "text": mapCard["text"].(string), "card": chaine}
// 	res, err := json.Marshal(mapRes)
// 	if err != nil {
// 		fmt.Println("Erreur du Marshal ", err)
// 		return
// 	}
// 	fmt.Fprintln(w, string(res))

// }

// router.GET("/nbcard/:userID", dbm.printNBCard)
// printNBCard output a json style file
// Query all rows to count them
func (dbm DbManager) printNBCard(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	user_id := ps.ByName("userID")

	que, err := dbm.db.Query("SELECT * FROM zeodine.cards WHERE user_id = " + user_id)
	if err != nil {
		log.Println(err)
		// panic(err)
	}
	nbCards := 0
	for que.Next() {
		nbCards += 1
		var card_id int
		var name string
		var img_url string
		var description string
		var pos_x int
		var pos_y int
		var ui map[string]interface{}
		err = que.Scan(&id, &name, &img_url, &description, &pos_x, , &pos_y, &ui, &user_id)
		if err != nil {
			log.Println(err)
			// panic(err)
		}
	}
	fmt.Fprintln(w, "{nbCards: ", nbCards, "}")
}

// func (jsonS jsonManage) printNBCard(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	chaine := "card" + ps.ByName("id")
// 	fmt.Println(chaine)

// 	if jsonS.err != nil {
// 		mapErr := map[string]string{"erreur": "1", "nbcard": "-1"}
// 		res, err := json.Marshal(mapErr)
// 		if err != nil {
// 			fmt.Println("Erreur du Marshal ", err)
// 			return
// 		}
// 		fmt.Fprintln(w, string(res))
// 		return
// 	}

// 	mapRes := map[string]string{"erreur": "0", "nbcard": strconv.Itoa(jsonS.nbcard)}
// 	res, err := json.Marshal(mapRes)
// 	if err != nil {
// 		fmt.Println("Erreur du Marshal ", err)
// 		return
// 	}
// 	fmt.Fprintln(w, string(res))

// }

// router.GET("/save/:jsonCards", dbm.save)
func (dbm DbManager) save(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	
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

	// Creating a new USERS table 
	q := "CREATE TABLE IF NOT EXISTS zeodine.cards (card_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,  name VARCHAR(64) DEFAULT NULL,  img_url VARCHAR(128) DEFAULT NULL,  description VARCHAR(200) DEFAULT NULL,  pos_x INT(8) NOT NULL,  pos_y INT(8) NOT NULL,  ui JSON DEFAULT NULL,  user_id INT DEFAULT NULL, FOREIGN KEY user_id(user_id) REFERENCES users(user_id))"
	_, dbm.err = dbm.db.Exec(q)
	if dbm.err != nil {
		log.Println("Error when creating table:", dbm.err)
	}
	log.Println("Table ready to be used")

	// Creating a new CARDS table
	q = "CREATE TABLE IF NOT EXISTS zeodine.cards (card_id INT(16) NOT NULL AUTO_INCREMENT, name VARCHAR(64) DEFAULT NULL, img_url VARCHAR(128) DEFAULT NULL, description VARCHAR(200) DEFAULT NULL, pos_x INT(8) NOT NULL AUTO_INCREMENT, pos_y INT(8) NOT NULL AUTO_INCREMENT, ui JSON DEFAULT NULL, user_id INT(16) DEFAULT NULL, PRIMARY KEY (task_id, user_id))"
	_, dbm.err = dbm.db.Exec(q)
	if dbm.err != nil {
		log.Println("Error when creating table:", dbm.err)ll
	}
	log.Println("Table ready to be used")

	//inserting the default cards
	quer, err := dbm.db.Prepare("INSERT INTO zeodine.cards VALUES (?, ?, ?, ?, ?, ?, ?, NULL)")
	if err != nil {
		log.Println(err)
	}
	defer quer.Close()

	// //Parsing global cards
	// file, err := os.Open("card.json")
	// defer file.Close()

	// if err != nil {
	// 	panic("Erreur ouverture fichier")
	// }
	// info, _ := file.Stat()
	// b       := make([]byte, info.Size())
	// n, err := file.Read(b)

	// if err != nil {
	// 	panic("Erreur lecture fichier")
	// }
	// // fmt.Println(n, " octets lus")
	// json.Unmarshal(b, &jsonS.data)

	// jsonS.nbcard = len(jsonS.data)
	// fmt.Println(jsonS)

	// Query
	_, err = quer.Exec(69, "test_name", "test_url", "test_desc", 50, 50, '{"test":"json"}', 42)
	if err != nil {
		log.Println(err)
	}

	return dbm
}

func main() {
	var jsonS jsonManage
	var dbm DbManager

	dbm = dbm.setupDB()

	defer dbm.close()


	router := httprouter.New()
	router.GET("/nbcard/:userID", dbm.printNBCard)
	router.GET("/save/:jsonCards", dbm.save)
	router.GET("/load/:userID", dbm.load)
	router.GET("/newuser/:userID", dbm.newuser)
	router.GET("/card/:userID/:id", dbm.card)
	log.Fatal(http.ListenAndServe(":8080", router))
}
