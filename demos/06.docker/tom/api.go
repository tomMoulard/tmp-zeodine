package main

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strconv"
)

var db *sql.DB
var id int = 0

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
		"INSERT INTO example VALUES (" + i + ",'" + ps.ByName("name") + "')")
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
		"DELETE FROM example WHERE id = " + ps.ByName("id") + ")")
	if err != nil {
		fmt.Println(err)
	}
	defer ins.Close()

	fmt.Fprintln(w, "You've just deleted the name corresponding of the id:", id)

}

func quit(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Dropping database
	_, err := db.Exec("DROP DATABASE zeodine")
	if err != nil {
		panic(err)
	}
	fmt.Println("Database DROPED")
}

func main() {
	// Connecting to db
	var err error
	db, err = sql.Open("mysql", "tom:multipass@tcp(127.0.0.1:3306)/")
	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	// Creating a new database
	_, err = db.Exec("CREATE DATABASE zeodine")
	if err != nil {
		panic(err)
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
		panic(err)
	}
	fmt.Println("Table created")

	// Routing
	router := httprouter.New()
	router.GET("/", showTable)
	router.GET("/add/:name", add)
	router.GET("/rm/:id", rm)
	router.GET("/quit", quit)

	log.Fatal(http.ListenAndServe(":8080", router))
}
