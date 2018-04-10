package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type myBDD struct {
	dataSource string
	db         *sql.DB
}

func chooseRandomString() string {
	svrStrings := []string{"toto", "tom", "arnaud", "charle", "brahim", "lisa"}
	nb := rand.Intn(len(svrStrings))
	return svrStrings[nb]
}

func (mdb myBDD) initialise() myBDD {
	var err error
	mdb.dataSource = "arnaud:nono@tcp(db:3306)/"
	mdb.db, err = sql.Open("mysql", mdb.dataSource)
	if err != nil {
		fmt.Println("Erreur lors de l'ouverture de la BDD")
		log.Fatal(err)
	}
	fmt.Println("BDD ouverte " + mdb.dataSource)
	return mdb
}

func main() {
	var mdb myBDD
	var err error
	mdb = mdb.initialise()

	defer mdb.db.Close()
	i := 0

	for i < 3 {
		err = mdb.db.Ping()
		if err != nil {
			fmt.Println("On attend 5 sec")
			time.Sleep(5 * time.Second)
			i++
		} else {
			i = 10
		}
	}

	if i == 3 {
		fmt.Println("Connexion à la ddb trop longue")
		os.Exit(1)
	}

	/*
		_, err = mdb.db.Exec("CREATE DATABASE test")
		if err != nil {
			fmt.Println("Erreur lors de l'addition d'une BDD")
			log.Fatal(err)
		}
		fmt.Println("BDD créer")

		_, err = mdb.db.Exec("USE test")
		if err != nil {
			fmt.Println("Erreur lors du use")
			log.Fatal(err)
		}
		fmt.Println("USE test")
	*/
	_, err = mdb.db.Exec("CREATE TABLE bd1.example (id integer, name varchar(32))")
	if err != nil {
		fmt.Println("Table déjà existante")

	} else {
		fmt.Println("Table créée")
	}

	_, err = mdb.db.Exec("USE bd1")
	if err != nil {
		panic(err)
	}
	fmt.Println("Table bd1 selectionné")

	stmt, err := mdb.db.Prepare("INSERT INTO example(id, name) VALUES(?, ?)")
	if err != nil {
		fmt.Println("Erreur du statement")
		log.Fatal(err)
	}

	nb := rand.Intn(100)
	str := chooseRandomString()
	res, err := stmt.Exec(nb, str)
	if err != nil {
		fmt.Println("Erreur de l'insertion")
		log.Fatal(err)
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		fmt.Println("LastInsertId")
		log.Fatal(err)
	}

	fmt.Println(str, " ajouté ", nb, " lastID : ", lastID)

	http.HandleFunc("/", mdb.printPage)
	http.ListenAndServe(":8081", nil)

}

func (mdb myBDD) printPage(w http.ResponseWriter, r *http.Request) {
	var err error
	fmt.Println(mdb)

	stmt, err := mdb.db.Prepare("SELECT * FROM bd1.example")
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
