package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var dataSource string

	if os.Getenv("MYSQL_PORT_3306_TCP_PORT") == "" {
		dataSource = "arnaud:nono@tcp(127.0.0.1:3306)/"
	} else {
		dataSource = "arnaud:nono@tcp(" + os.Getenv("MYSQL_PORT_3306_TCP_ADDR") +
			":" + os.Getenv("MYSQL_PORT_3306_TCP_PORT") + ")/"
	}

	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture de la BDD")
		log.Fatal(err)
	}
	fmt.Println("BDD ouverte " + dataSource)
	defer db.Close()

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
		fmt.Println("Erreur lors de la creation d'une table")
		log.Fatal(err)
	}
	fmt.Println("Table créer")

}
