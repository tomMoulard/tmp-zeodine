package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
	// "encoding/json"

	"log"
	"net/http"
	"os"
	// "strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

type jsonManage struct {
	err    error
	data   map[string]interface{}
	nbcard int
}

type DbManager struct {
	dataSource string
	db         *sql.DB
	err        error
}

// // router.GET("/card/:userID/:id", dbm.card)
// func (dbm DbManager) card(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	card_id := ps.ByName("id")
// 	user_id := ps.ByName("userID")

// 	que, err := dbm.db.Query("SELECT * FROM zeodine.cards WHERE card_id = " + card_id + " AND user_id = " + user_id)
// 	if err != nil {
// 		log.Println(err)
// 		// panic(err)
// 	}
// 	for que.Next() {
// 		var card_id int
// 		var name string
// 		var img_url string
// 		var description string
// 		var pos_x int
// 		var pos_y int
// 		var ui string
// 		err = que.Scan(&card_id, &name, &img_url, &description, &pos_x, &pos_y, &ui, &user_id)
// 		if err != nil {
// 			log.Println("error when query a card", err)
// 			// panic(err)
// 		}
// 		fmt.Fprintln(w, "{card_id:", card_id, ",name:", name, ",img_url:", img_url, ",description:", description, ",pos_x:", pos_x, ",pos_y:", pos_y, ",ui:", ui, ",user_id:", user_id, "}")
// 	}
// }

// // router.GET("/nbcard/:userID", dbm.printNBCard)
// // printNBCard output a json style file
// // Query all rows to count them
// func (dbm DbManager) printNBCard(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
// 	user_id := ps.ByName("userID")

// 	que, err := dbm.db.Query("SELECT COUNT(*) FROM zeodine.cards WHERE user_id = " + user_id)
// 	if err != nil {
// 		log.Println(err)
// 		// panic(err)
// 	}
// 	for que.Next() {
// 		var nbCards int
// 		err = que.Scan(&nbCards)
// 		if err != nil {
// 			log.Println(err)
// 			// panic(err)
// 		}
// 		fmt.Fprintln(w, "{nbCards: ", nbCards, "}")
// 	}
// }

// router.GET("/newuser/:userID ", dbm.newuser)
func (dbm DbManager) newuser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	quer, err := dbm.db.Prepare("INSERT INTO zeodine.users VALUES ( ? )")
	if err != nil {
		fmt.Fprint(w, "{ saved:false, err: %s, userID: %s }", err.Error(), ps.ByName("userID"))
		return
	}
	defer quer.Close()
	i, err := strconv.Atoi(ps.ByName("userID"))
	// log.Println(i, err, ps.ByName("userID"))
	if err != nil {
		fmt.Fprint(w, "{ saved:false, err: %s, userID: %s }", err.Error(), ps.ByName("userID"))
		return
	}
	_, err = quer.Exec(i)
	if err != nil {
		fmt.Fprint(w, "{ saved:false, err: %s, userID: %s }", err.Error(), ps.ByName("userID"))
		return
	}
	fmt.Fprintln(w, "{saved:True}")
}

// router.GET("/ws/:userID ", dbm.ws)
func (dbm DbManager) ws(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	que, err := dbm.db.Query("SELECT ws_id, ws_name, user_id FROM zeodine.ws WHERE user_id = " + ps.ByName("userID"))
	if err != nil {
		fmt.Fprint(w, "{ err: %s, userID: %s }", err.Error(), ps.ByName("userID"))
		return
	}
	res := "{"
	for que.Next() {
		var ws_id int
		var ws_name string
		var user_id int
		err = que.Scan(&ws_id, &ws_name, &user_id)
		if err != nil {
			res += "{err:" + err.Error() + ", userID:" + ps.ByName("userID") + "}"
		} else {
			res += "{ ws_id:" + strconv.Itoa(ws_id) + ", ws_name:" + ws_name + ", user_id:" + strconv.Itoa(user_id) + "}"
		}
	}
	res += "}"
	fmt.Fprintln(w, res)
}

// router.GET("/createws/:userID/:wsName ", dbm.createws)
func (dbm DbManager) createws(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	user_id, err := strconv.Atoi(ps.ByName("userID"))
	wsName := ps.ByName("wsName")

	if err != nil {
		fmt.Fprint(w, "{ws_id: -1, err: %s, userID: %s }", err.Error(), ps.ByName("userID"))
		return
	}
	//getting ws.lengh
	nbWS := time.Now().Format("20060102150405")

	quer, err := dbm.db.Prepare("INSERT INTO zeodine.ws VALUES (?, ?, ?)")
	if err != nil {
		fmt.Fprint(w, "{ws_id: -1, err: %s, userID: %s }", err.Error(), ps.ByName("userID"))
		return
	}

	_, err = quer.Exec(nbWS, wsName, user_id)
	if err != nil {
		fmt.Fprint(w, "{ws_id: -1, err: %s, userID: %s }", err.Error(), ps.ByName("userID"))
		return
	}
	fmt.Fprint(w, "{ws_id: %d}", nbWS)
}

// router.GET("/nbcard/:userID/:wsID ", dbm.nbcard)
func (dbm DbManager) nbcard(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// user_id, err1 := strconv.Atoi(ps.ByName("userID"))
	// ws_id, err2 := strconv.Atoi(ps.ByName("wsID"))
	// if err1 != nil {
	// 	fmt.Fprint(w, "{nb_card: -1, err: %s, userID: %s }", err1.Error(), ps.ByName("userID"))
	// 	return
	// }
	// if err2 != nil {
	// 	fmt.Fprint(w, "{nb_card: -1, err: %s, userID: %s }", err2.Error(), ps.ByName("userID"))
	// 	return
	// }
	que, err := dbm.db.Query("SELECT stack_id FROM zeodine.stacks where user_id = " + ps.ByName("userID") + " AND ws_id = " + ps.ByName("wsID"))
	if err != nil {
		fmt.Fprint(w, "{nb_card: -1, err: %s, userID: %s }", err.Error(), ps.ByName("userID"))
		return
	}
	nbcard := 0
	for que.Next() {
		var stack_id int
		err = que.Scan(&stack_id)
		if err != nil {
			fmt.Fprint(w, "{nb_card: -1, err: %s, userID: %s }", err.Error(), ps.ByName("userID"))
			return
		}
		nbcard += 1
	}
	fmt.Fprintln(w, "{nb_card: "+strconv.Itoa(nbcard)+"}")

}

// router.GET("/load/:userID/:wsID ", dbm.load)
func (dbm DbManager) load(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {}

// router.GET("/card/:userID/:wsID/:cardID", dbm.card)
func (dbm DbManager) card(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {}

// router.GET("/save/:json", dbm.save)
func (dbm DbManager) save(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {}

func (dbm DbManager) createTable(tableName, tableContent string) {
	q := "CREATE TABLE IF NOT EXISTS " + tableName + " (" + tableContent + ")"
	_, dbm.err = dbm.db.Exec(q)
	if dbm.err != nil {
		log.Println("Error when creating "+tableName+" table:", dbm.err)
	}
	log.Println(tableName, "table ready to be used")
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

	// Creating a new CARDS table
	dbm.createTable("zeodine.cards", "card_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY, body JSON DEFAULT NULL, stack_id INT DEFAULT NULL")

	// Creating a new USERS table
	dbm.createTable("zeodine.users", "user_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY")

	// Creating a new WS table
	dbm.createTable("zeodine.ws", "ws_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,ws_name VARCHAR(64) DEFAULT NULL,user_id INT DEFAULT NULL,FOREIGN KEY user_id(user_id) REFERENCES users(user_id)")

	// Creating a new WS table
	dbm.createTable("zeodine.stacks", "stack_id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,group_id INT DEFAULT NULL,user_id INT DEFAULT NULL,card_id INT DEFAULT NULL,ws_id INT DEFAULT NULL,FOREIGN KEY user_id(user_id) REFERENCES users(user_id),FOREIGN KEY card_id(card_id) REFERENCES cards(card_id),FOREIGN KEY ws_id(ws_id) REFERENCES ws(ws_id)")

	var jsonS jsonManage
	// //Parsing global cards
	file, err := os.Open("card.json")
	defer file.Close()

	if err != nil {
		panic("Erreur ouverture fichier")
	}
	info, _ := file.Stat()
	b := make([]byte, info.Size())
	_, err = file.Read(b)

	if err != nil {
		panic("Erreur lecture fichier")
	}
	// fmt.Println(n, " octets lus")
	json.Unmarshal(b, &jsonS.data)

	jsonS.nbcard = len(jsonS.data)
	// fmt.Println(jsonS)

	//inserting the default cards
	quer, err := dbm.db.Prepare("INSERT INTO zeodine.cards VALUES (?, ?, NULL)")
	if err != nil {
		log.Println(err)
	}
	defer quer.Close()

	// fmt.Println(jsonS.data, jsonS.data.(map[string]interface{})["card2"])

	// Query
	// for i := 1; i < 6; i++ {
	// 	// j := strconv.Itoa(i)
	// 	// _, err = quer.Exec(i, jsonS.data.(map[string]interface{})["card"+j])
	// 	_, err = quer.Exec(i, "{card:42}")
	// 	if err != nil {
	// 		log.Println(err)
	// 	}
	// }

	return dbm
}

func main() {
	// var jsonS jsonManage
	var dbm DbManager

	dbm = dbm.setupDB()

	defer dbm.db.Close()

	router := httprouter.New()

	router.GET("/newuser/:userID ", dbm.newuser)
	router.GET("/ws/:userID ", dbm.ws)
	router.GET("/createws/:userID/:wsName ", dbm.createws)
	router.GET("/nbcard/:userID/:wsID ", dbm.nbcard)
	router.GET("/load/:userID/:wsID ", dbm.load)
	router.GET("/card/:userID/:wsID/:cardID", dbm.card)
	router.GET("/save/:json", dbm.save)
	log.Fatal(http.ListenAndServe(":8080", router))
}
