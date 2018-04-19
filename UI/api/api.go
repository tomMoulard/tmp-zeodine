package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

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

func (dbm DbManager) getLastId() uint64 {
	var res uint64
	que, err := dbm.db.Prepare("SELECT LAST_INSERT_ID()")
	if err != nil {
		log.Printf("[0]Error when getLastID(): %s", err.Error())
		return 0
	}
	defer que.Close()

	quer, err := que.Query()
	if err != nil {
		log.Printf("[1]Error when getLastID(): %s", err.Error())
		return 0
	}

	for quer.Next() {
		err = quer.Scan(&res)
		if err != nil {
			log.Printf("[2]Error when getLastID(): %s", err.Error())
			return 0
		}
	}
	return res
}

// router.GET("/newuser ", dbm.newuser)
func (dbm DbManager) newuser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)
	quer, err := dbm.db.Prepare("INSERT INTO zeodine.users VALUES ( NULL )")
	if err != nil {
		fmt.Fprintf(w, "{ \"user_id\":0, \"err\": \"%s\", code:0}", err.Error())
		return
	}
	defer quer.Close()

	_, err = quer.Exec()
	if err != nil {
		fmt.Fprintf(w, "{ \"user_id\":0, \"err\": \"%s\", code:1 }", err.Error())
		return
	}

	userID := dbm.getLastId()

	w.WriteHeader(200)
	fmt.Fprintf(w, "{\"user_id\":%d}", userID)
}

// router.GET("/ws/:userID ", dbm.ws)
func (dbm DbManager) ws(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)
	que, err := dbm.db.Prepare("SELECT ws_id, ws_name, user_id FROM zeodine.ws WHERE user_id = ?")
	if err != nil {
		fmt.Fprintf(w, "{ \"err\": \"%s\", \"userID\": %s, \"code\":0 }", err.Error(), ps.ByName("userID"))
		return
	}
	defer que.Close()
	quer, err := que.Query(ps.ByName("userID"))
	if err != nil {
		fmt.Fprintf(w, "{ \"err\": \"%s\", \"userID\": %s, \"code\":1 }", err.Error(), ps.ByName("userID"))
		return
	}
	firstWs := true
	res := "{ \"ws\": ["
	for quer.Next() {
		if firstWs {
			firstWs = false
		} else {
			res += ","
		}
		var ws_id int
		var ws_name string
		var user_id int
		err = quer.Scan(&ws_id, &ws_name, &user_id)
		if err != nil {
			res += "{\"err\": \"" + err.Error() + "\", \"userID\":" + ps.ByName("userID") + "}"
		} else {
			res += "{\"ws_id\":" + strconv.Itoa(ws_id) + ", \"ws_name\": \"" + ws_name + "\", \"user_id\":" + strconv.Itoa(user_id) + "}"

		}
	}
	res += "]}"
	w.WriteHeader(200)
	fmt.Fprintln(w, res)
}

// router.GET("/createws/:userID/:wsName ", dbm.createws)
func (dbm DbManager) createws(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)
	user_id, err := strconv.Atoi(ps.ByName("userID"))
	wsName := ps.ByName("wsName")

	if err != nil {
		fmt.Fprintf(w, "{\"ws_id\": 0, \"err\": \"%s\", \"userID\": %s, \"code\":0 }", err.Error(), ps.ByName("userID"))
		return
	}

	quer, err := dbm.db.Prepare("INSERT INTO zeodine.ws VALUES (NULL, ?, ?)")
	if err != nil {
		fmt.Fprintf(w, "{\"ws_id\": 0, \"err\": \"%s\", \"userID\": %s, \"code\":1 }", err.Error(), ps.ByName("userID"))
		return
	}
	defer quer.Close()

	_, err = quer.Exec(wsName, user_id)
	if err != nil {
		fmt.Fprintf(w, "{\"ws_id\": 0, \"err\": \"%s\", \"userID\": %s, \"code\":2 }", err.Error(), ps.ByName("userID"))
		return
	}

	nbWS := dbm.getLastId()

	w.WriteHeader(200)
	fmt.Fprintf(w, "{\"ws_id\": %d}", nbWS)
}

// router.GET("/nbcard/:userID/:wsID ", dbm.nbcard)
func (dbm DbManager) nbcard(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	que, err := dbm.db.Prepare("SELECT stack_id FROM zeodine.stacks WHERE user_id = ? AND ws_id = ?")
	if err != nil {
		fmt.Fprintf(w, "{\"nb_card\": 0, \"err\": \"%s\", \"userID\": %s, \"code\":0 }", err.Error(), ps.ByName("userID"))
		return
	}
	defer que.Close()

	quer, err := que.Query(ps.ByName("userID"), ps.ByName("wsID"))
	if err != nil {
		fmt.Fprintf(w, "{\"nb_card\": 0, \"err\": \"%s\", \"userID\": %s, \"code\":1 }", err.Error(), ps.ByName("userID"))
		return
	}
	nbcard := 0
	for quer.Next() {
		var stack_id int
		err = quer.Scan(&stack_id)
		if err != nil {
			fmt.Fprintf(w, "{\"nb_card\": 0, \"err\": \"%s\", \"userID\": %s, \"code\":2 }", err.Error(), ps.ByName("userID"))
			return
		}
		nbcard += 1
	}
	w.WriteHeader(200)
	fmt.Fprintln(w, "{\"nb_card\": "+strconv.Itoa(nbcard)+"}")
}

// router.GET("/load/:userID/:wsID ", dbm.load)
func (dbm DbManager) load(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)
	// Query stacks
	que, err := dbm.db.Prepare("SELECT stack_id FROM zeodine.stacks where user_id = ? AND ws_id =  ?")
	if err != nil {
		fmt.Fprint(w, "{\"err\": \"%s\", \"userID\": %s, \"code\":0 }", err.Error(), ps.ByName("userID"))
		return
	}
	defer que.Close()

	quer, err := dbm.db.Query(ps.ByName("userID"), ps.ByName("wsID"))
	if err != nil {
		fmt.Fprint(w, "{\"err\": \"%s\", \"userID\": %s, \"code\":1 }", err.Error(), ps.ByName("userID"))
		return
	}
	res := "{ \"cards\": ["
	for quer.Next() {
		var stack_id int
		err = quer.Scan(&stack_id)
		if err != nil {
			fmt.Fprint(w, "{\"err\": \"%s\", \"userID\": %s, \"code\":2 }", err.Error(), ps.ByName("userID"))
			return
		}
		stack_id_str := strconv.Itoa(stack_id)
		// Query card
		que2, err := dbm.db.Prepare("SELECT body FROM zeodine.cards where stack_id = ?")
		if err != nil {
			fmt.Fprint(w, "{\"err\": \"%s\", \"userID\": %s, \"code\":3 }", err.Error(), ps.ByName("userID"))
			return
		}
		defer que2.Close()

		quer2, err := dbm.db.Query(stack_id_str)
		if err != nil {
			fmt.Fprint(w, "{\"err\": \"%s\", \"userID\": %s, \"code\":4 }", err.Error(), ps.ByName("userID"))
			return
		}
		firstCard := true
		for quer2.Next() {
			if firstCard {
				firstCard = false
			} else {
				res += ","
			}
			var card string
			err = quer2.Scan(&card)
			if err != nil {
				res += "{\"err\": \"" + err.Error() + "\", \"userID\": " + ps.ByName("userID") + ", \"code\":5 }"
			} else {
				res += card
			}
		}
	}
	w.WriteHeader(200)
	fmt.Fprintln(w, res+"]}")
}

// router.GET("/card/:userID/:wsID/:cardID", dbm.card)
func (dbm DbManager) card(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)
	// Query stacks
	que, err := dbm.db.Prepare("SELECT card_id FROM zeodine.stacks where user_id = ? AND ws_id = ?")
	if err != nil {
		fmt.Fprint(w, "{\"err\": \"%s\", \"userID\": %s, \"code\":0 }", err.Error(), ps.ByName("userID"))
		return
	}
	defer que.Close()

	quer, err := dbm.db.Query(ps.ByName("userID"), ps.ByName("wsID"))
	if err != nil {
		fmt.Fprint(w, "{\"err\": \"%s\", \"userID\": %s, \"code\":1 }", err.Error(), ps.ByName("userID"))
		return
	}
	exist := false
	card_id_str := ""
	for quer.Next() && !exist {
		var card_id int
		err = quer.Scan(&card_id)
		if err != nil {
			fmt.Fprint(w, "{\"err\": \"%s\", \"userID\": %s, \"code\":2 }", err.Error(), ps.ByName("userID"))
			return
		}
		card_id_str = strconv.Itoa(card_id)
		if card_id_str == ps.ByName("cardID") {
			exist = true
		}
	}
	if exist {
		// Query card
		que, err = dbm.db.Prepare("SELECT body FROM zeodine.cards where card_id = ?")
		if err != nil {
			fmt.Fprint(w, "{\"err\": \"%s\", \"userID\": %s, \"code\":3 }", err.Error(), ps.ByName("userID"))
			return
		}
		defer que.Close()
		quer, err = dbm.db.Query(card_id_str)
		if err != nil {
			fmt.Fprint(w, "{\"err\": \"%s\", \"userID\": %s, \"code\":4 }", err.Error(), ps.ByName("userID"))
			return
		}
		for quer.Next() {
			var card string
			err = quer.Scan(&card)
			if err != nil {
				fmt.Fprint(w, "{\"err\": \"%s\", \"userID\": %s, \"code\":5 }", err.Error(), ps.ByName("userID"))
				return
			}
			w.WriteHeader(200)
			fmt.Fprintln(w, card)
		}
	} else {
		fmt.Fprint(w, "{\"err\": Error: There is no card with this id, \"userID\": %s, \"code\":6 }", ps.ByName("userID"))
	}
}

type Save struct {
	Groupes []struct {
		Cards []struct {
			Card struct {
				CardContent string `json:"card_content"`
			} `json:"card"`
			CardID int `json:"card_id"`
		} `json:"cards"`
		GroupeID int `json:"groupe_id"`
	} `json:"groupes"`
	UserID int `json:"user_id"`
	WsID   int `json:"ws_id"`
}

// router.GET("/save/:save", dbm.save)
// curl localhost:8080/save/{"user_id": 42,"ws_id": 69,"groupes": [{"groupe_id": 12,"cards": [{"card_id": 1,"card": {"card_content": ""}},{"card_id": 2,"card": {"card_content": ""}},{"card_id": 3,"card": {"card_content": ""}},{"card_id": 4,"card": {"card_content": ""}}]},{"groupe_id": 21,"cards": [{"card_id": 1,"card": {"card_content": ""}}]}]}
func (dbm DbManager) save(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(500)

	// save := ps.ByName("save")
	save := `{"user_id": 1524134267,"ws_id": 1524134288,"groupes": [{"groupe_id": 12,"cards": [{"card_id": 1,"card": {"card_content": "{}"}},{"card_id": 2,"card": {"card_content": "{}"}},{"card_id": 3,"card": {"card_content": "{}"}},{"card_id": 4,"card": {"card_content": "{}"}}]},{"groupe_id": 21,"cards": [{"card_id": 5,"card": {"card_content": "{\"card_pos\":12}"}}]}]}`

	var saveStruct Save

	json.Unmarshal([]byte(save), &saveStruct)
	// Foreach group
	//      add the stack
	//      replace all cards with right the stack id
	for _, group := range saveStruct.Groupes {
		// For all cards
		for _, card := range group.Cards {
			time.Sleep(1 * time.Second) // Remove a bug when generating new id
			// query the right stack
			que, err := dbm.db.Prepare("SELECT stack_id FROM zeodine.stacks WHERE group_id = ? AND user_id = ? AND card_id = ? AND ws_id = ?")
			if err != nil {
				fmt.Fprintf(w, "{\"saved\": false, \"error\": %s, \"code\":0}", err.Error())
				return
			}
			defer que.Close()

			quer, err := que.Query(group.GroupeID, saveStruct.UserID, card.CardID, saveStruct.WsID)
			if err != nil {
				fmt.Fprintf(w, "{\"saved\": false, \"error\": %s, \"code\":1}", err.Error())
				return
			}
			var stack_id uint64
			stack_id = 0
			for quer.Next() {
				err := quer.Scan(&stack_id)
				if err != nil {
					fmt.Fprintf(w, "{\"saved\": false, \"error\": %s, \"code\":2}", err.Error())
					return
				}
			}
			if stack_id == 0 {
				que2, err := dbm.db.Prepare("INSERT INTO zeodine.stacks VALUE (NULL, ?, ?, ?, ?)")
				if err != nil {
					fmt.Fprintf(w, "{\"saved\": false, \"error\": %s, \"code\":3}", err.Error())
					return
				}
				defer que2.Close()
				_, err = que2.Query(group.GroupeID, saveStruct.UserID, card.CardID, saveStruct.WsID)
				if err != nil {
					fmt.Fprintf(w, "{\"saved\": false, \"error\": %s, \"code\":4}", err.Error())
					return
				}
				stack_id = dbm.getLastId()
			}
			// Replacing card || creating card
			// query the card -> if !exist -> create card
			que3, err := dbm.db.Prepare("SELECT card_id FROM zeodine.cards where card_id = ? AND body = ?")
			if err != nil {
				fmt.Fprintf(w, "{ \"saved\":false, \"err\": \"%s\", code:5}", err.Error())
				return
			}
			defer que3.Close()
			quer3, err := que3.Query(card.CardID, card.Card.CardContent)
			if err != nil {
				fmt.Fprintf(w, "{ \"saved\":false, \"err\": \"%s\", code:6}", err.Error())
				return
			}
			var cardID uint64
			for quer3.Next() {
				err = quer3.Scan(&cardID)
				if err != nil {
					fmt.Fprintf(w, "{ \"saved\":false, \"err\": \"%s\", code:7}", err.Error())
					return
				}
			}
			if cardID == 0 { // no card found -> card created
				que4, err := dbm.db.Prepare("INSERT INTO zeodine.cards value (?, ?, ?)")
				if err != nil {
					fmt.Fprintf(w, "{ \"saved\":false, \"err\": \"%s\", code:8}", err.Error())
					return
				}
				defer que4.Close()
				_, err = que4.Query(card.CardID, card.Card.CardContent, stack_id)
				if err != nil {
					fmt.Fprintf(w, "{ \"saved\":false, \"err\": \"%s\", code:9}", err.Error())
					return
				}
				//else crad created !
			}
		}
	}
	w.WriteHeader(200)
	fmt.Fprintln(w, "{\"saved\": true}")

}

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
	dbm.createTable("zeodine.cards", "card_id INT(32) NOT NULL AUTO_INCREMENT PRIMARY KEY, body JSON DEFAULT NULL, stack_id INT(32) DEFAULT NULL")

	// Creating a new USERS table
	dbm.createTable("zeodine.users", "user_id INT(32) NOT NULL AUTO_INCREMENT PRIMARY KEY")

	// Creating a new WS table
	dbm.createTable("zeodine.ws", "ws_id INT(32) NOT NULL AUTO_INCREMENT PRIMARY KEY,ws_name VARCHAR(64) DEFAULT NULL,user_id INT(32) DEFAULT NULL,FOREIGN KEY user_id(user_id) REFERENCES users(user_id)")

	// Creating a new WS table
	dbm.createTable("zeodine.stacks", "stack_id INT(32) NOT NULL AUTO_INCREMENT PRIMARY KEY, group_id INT(32) DEFAULT NULL, user_id INT(32) DEFAULT NULL, card_id INT(32) DEFAULT NULL, ws_id INT(32) DEFAULT NULL, FOREIGN KEY user_id(user_id) REFERENCES users(user_id), FOREIGN KEY ws_id(ws_id) REFERENCES ws(ws_id)")

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
	//  // j := strconv.Itoa(i)
	//  // _, err = quer.Exec(i, jsonS.data.(map[string]interface{})["card"+j])
	//  _, err = quer.Exec(i, "{card:42}")
	//  if err != nil {
	//      log.Println(err)
	//  }
	// }

	return dbm
}

func main() {
	// var jsonS jsonManage
	var dbm DbManager

	dbm = dbm.setupDB()

	defer dbm.db.Close()

	router := httprouter.New()

	router.GET("/newuser", dbm.newuser)
	router.GET("/ws/:userID", dbm.ws)
	router.GET("/createws/:userID/:wsName", dbm.createws)
	router.GET("/nbcard/:userID/:wsID", dbm.nbcard)
	router.GET("/load/:userID/:wsID", dbm.load)
	router.GET("/card/:userID/:wsID/:cardID", dbm.card)
	router.GET("/save/:save", dbm.save)
	log.Fatal(http.ListenAndServe(":8080", router))
}
