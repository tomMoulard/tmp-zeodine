package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

type DbManager struct {
	id         int
	dataSource string
	db         *sql.DB
	err        error
}

func (dbm DbManager) showTable(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	// Query the table
	que, err := dbm.db.Query("SELECT * FROM zeodine.example")
	if err != nil {
		log.Println(err)
	}
	fmt.Fprintln(w, "Table content:")
	for que.Next() {
		var id int
		var name string
		err = que.Scan(&id, &name)
		if err != nil {
			log.Println(err)
		}
		fmt.Fprintln(w, id, ":", name)
	}
}

func (dbm DbManager) add(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	// Inserting some stuff
	i := strconv.Itoa(dbm.id)

	// Prepare
	quer, err := dbm.db.Prepare("INSERT INTO zeodine.example VALUES (?, ?)")
	if err != nil {
		log.Println(err)
	}
	defer quer.Close()

	// Query
	_, err = quer.Exec(i, ps.ByName("name"))
	if err != nil {
		log.Println(err)
	}

	fmt.Fprintln(w, "You've just insered", ps.ByName("name"), "with the id of", dbm.id)
	dbm.id += 1
}

func (dbm DbManager) rm(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	// Deleting some stuff
	quer, err := dbm.db.Prepare("DELETE FROM zeodine.example WHERE id = ?")
	if err != nil {
		log.Println(err)
	}
	defer quer.Close()

	_, err = quer.Exec(ps.ByName("id"))
	if err != nil {
		log.Println(err)
	}

	fmt.Fprintln(w, "You've just deleted the name corresponding of the id:", ps.ByName("id"))

}

func (dbm DbManager) quit(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	// Dropping database
	_, err := dbm.db.Exec("DROP DATABASE zeodine")
	if err != nil {
		panic(err)
	}
	log.Println("Database DROPED")
	os.Exit(1) // Shutdown api
}

func (dbm DbManager) pingDB() bool {
	time.Sleep(time.Second * 5)
	err := dbm.db.Ping()
	if err != nil {
		log.Println(err.Error())
		return false
	}
	return true
}

func (dbm DbManager) setupDB() {
	log.Println("Initializing the database...")

	dbm.dataSource = "tom:multipass@tcp(db:3306)/"

	log.Printf("dataSource = %v\n", dbm.dataSource)

	// Connecting to db
	dbm.db, dbm.err = sql.Open("mysql", dbm.dataSource)
	if dbm.err != nil {
		panic(dbm.err.Error())
	}
	defer dbm.db.Close()

	var pinged bool = false
	for i := 0; i < 3 && !pinged; i++ {
		log.Println("Connecting to database ... try:", i)
		if dbm.pingDB() {
			pinged = true
		}
	}
	if !pinged {
		log.Println("Could not connect to db in time")
		os.Exit(0)
	}
	log.Println("Connected")
	// Creating a new database
	_, dbm.err = dbm.db.Exec("CREATE DATABASE IF NOT EXISTS zeodine")
	if dbm.err != nil {
		// log.Println(err)
		log.Printf("Error when creating db: %v\n", dbm.err)
	}
	log.Println("Database created")

	// Using the fleshly created database:
	_, dbm.err = dbm.db.Exec("USE zeodine")
	if dbm.err != nil {
		log.Println(dbm.err)
	}

	// Creating a new table to insert some elements
	_, dbm.err = dbm.db.Exec("CREATE TABLE IF NOT EXISTS zeodine.example ( id integer, data varchar(32) )")
	if dbm.err != nil {
		log.Println("Error when creating table:", dbm.err)
	}
	log.Println("Table ready to be used")
}

func main() {
	var dbm DbManager

	dbm.setupDB()

	// Routing
	router := httprouter.New()
	router.GET("/", dbm.showTable)
	router.GET("/add/:name", dbm.add)
	router.GET("/rm/:id", dbm.rm)
	router.GET("/quit", dbm.quit)

	log.Fatal(http.ListenAndServe(":8081", router))
}
