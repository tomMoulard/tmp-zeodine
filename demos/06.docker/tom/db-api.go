package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

var id int = 0
var dataSource string
var db *sql.DB
var err error

func init() {

	fmt.Println("Initializing the database...")

	dataSource = "tom:multipass@tcp(db:3306)/"

	fmt.Printf("dataSource = %v\n", dataSource)

	// Connecting to db
	db, err = sql.Open("mysql", dataSource)
	if err != nil {
		panic(err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
}

func showTable(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	// Query the table
	que, err := db.Query("SELECT * FROM zeodine.example")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintln(w, "Table content:")
	for que.Next() {
		var id int
		var name string
		err = que.Scan(&id, &name)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Fprintln(w, id, ":", name)
	}
}

func add(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	// Inserting some stuff
	i := strconv.Itoa(id)
	ins, err := db.Query(
		"INSERT INTO zeodine.example VALUES (" + i + ",'" + ps.ByName("name") + "')")
	if err != nil {
		fmt.Println(err)
	}
	defer ins.Close()

	fmt.Fprintln(w, "You've just insered", ps.ByName("name"), "with the id of", id)
	id += 1
}

func rm(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	// Deleting some stuff
	ins, err := db.Query(
		"DELETE FROM zeodine.example WHERE id = " + ps.ByName("id"))
	if err != nil {
		fmt.Println(err)
	}
	defer ins.Close()

	fmt.Fprintln(w, "You've just deleted the name corresponding of the id:", ps.ByName("id"))

}

func quit(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	// Dropping database
	_, err := db.Exec("DROP DATABASE zeodine")
	if err != nil {
		panic(err)
	}
	fmt.Println("Database DROPED")
	os.Exit(1) // Shutdown api
}

func setupDB() {

	// Creating a new database
	_, err := db.Exec("CREATE DATABASE zeodine")
	if err != nil {
		// fmt.Println(err)
		log.Printf("A la création de la database zeodine, j'ai obtenu l'erreur : %v\n", err)
	}
	fmt.Println("Database created")

	// Using the fleshly created database:
	_, err = db.Exec("USE zeodine")
	if err != nil {
		panic(err)
	}

	// Creating a new table to insert some elements
	_, err = db.Exec("CREATE TABLE zeodine.example ( id integer, data varchar(32) )")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("Table created")
}

func main() {
	defer db.Close()

	setupDB()

	// Routing
	router := httprouter.New()
	router.GET("/", showTable)
	router.GET("/add/:name", add)
	router.GET("/rm/:id", rm)
	router.GET("/quit", quit)

	log.Fatal(http.ListenAndServe(":8081", router))
}
