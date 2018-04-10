package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var dataSource string

	fmt.Println("20 sec")
	time.Sleep(20 * time.Second)

	dataSource = "arnaud:nono@tcp(db:3306)/"

	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture de la BDD")
		log.Fatal(err)
	}
	fmt.Println("BDD ouverte " + dataSource)
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("Erreur lors du ping")
		panic(err.Error())
	}

	/*
		_, err = db.Exec("CREATE DATABASE test")
		if err != nil {
			fmt.Println("Erreur lors de l'addition d'une BDD")
			log.Fatal(err)
		}
		fmt.Println("BDD créer")

		_, err = db.Exec("USE test")
		if err != nil {
			fmt.Println("Erreur lors du use")
			log.Fatal(err)
		}
		fmt.Println("USE test")
	*/
	_, err = db.Exec("CREATE TABLE bd1.example (id integer, name varchar(32))")
	if err != nil {
		fmt.Println("Table déjà existante")

	} else {
		fmt.Println("Table créée")
	}

	_, err = db.Exec("USE bd1")
	if err != nil {
		panic(err)
	}
	fmt.Println("Table bd1 selectionné")

	stmt, err := db.Prepare("INSERT INTO example(name) VALUES(?)")
	if err != nil {
		fmt.Println("Erreur du statement")
		log.Fatal(err)
	}

	res, err := stmt.Exec("Arnaud")
	if err != nil {
		fmt.Println("Erreur de l'insertion")
		log.Fatal(err)
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		fmt.Println("LastInsertId")
		log.Fatal(err)
	}

	fmt.Println("Arnaud ajouté %i", lastID)

	http.HandleFunc("/", printPage)
	http.ListenAndServe(":8081", nil)

}

func printPage(w http.ResponseWriter, r *http.Request) {
	var dataSource string

	dataSource = "arnaud:nono@tcp(db:3306)/"

	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture de la BDD (2)")
		log.Fatal(err)
	}
	fmt.Println("BDD ouverte (2) " + dataSource)
	defer db.Close()

	err = db.Ping()
	if err != nil {
		fmt.Println("Erreur lors du ping (2)")
		//panic(err.Error()) // proper error handling instead of panic in your app
	}

	fmt.Println("BDD ouverte (2) " + dataSource)

	stmt, err := db.Prepare("SELECT * FROM bd1.example")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintln(w, "Table bd1.example:")
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Fprintln(w, id, ":", name)
	}
}
