package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"os"
	"strconv"
)

var id int = 0
var dataSource string

func showTable(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Connecting to db
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

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
	// Connecting to db
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

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
	// Connecting to db
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
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
	// Connecting to db
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	// Dropping database
	_, err = db.Exec("DROP DATABASE zeodine")
	if err != nil {
		panic(err)
	}
	fmt.Println("Database DROPED")
	os.Exit(1) // Shutdown api
}

func setupDB() {
	// Connecting to db
	db, err := sql.Open("mysql", dataSource)
	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Database connected!")
	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	// Creating a new database
	_, err = db.Exec("CREATE DATABASE zeodine")
	if err != nil {
		// fmt.Println(err)
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
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("Table created")
}

func main() {
	if os.Getenv("MYSQL_PORT_3306_TCP_PORT") == "" {
		dataSource = "tom:multipass@tcp(127.0.0.1:3306)/"
	} else {
		dataSource = "tom:multipass@tcp(" + os.Getenv("MYSQL_PORT_3306_TCP_ADDR") +
			":" + os.Getenv("MYSQL_PORT_3306_TCP_PORT") + ")/"
	}

	go setupDB()

	// Routing
	router := httprouter.New()
	router.GET("/", showTable)
	router.GET("/add/:name", add)
	router.GET("/rm/:id", rm)
	router.GET("/quit", quit)

	log.Fatal(http.ListenAndServe(":8081", router))
}
