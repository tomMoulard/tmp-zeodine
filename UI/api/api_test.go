package api

import (
	"bytes"
	"database/sql"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/julienschmidt/httprouter"
)

// func Router() *httprouter {
// 	r1 := httprouter.New()
// 	r1.GET("/testing-http/:test", Rt)
// 	return r1
// }

var dbm DbManager = SetupDB()

func SetupDB() DbManager {
	var dbm DbManager
	dbm.dataSource = "server:zeodine@tcp(db:3306)/zeodine"

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
	// log.Println("Connected")
	return dbm
}

func TestNewuser(t *testing.T) {
	r1 := httprouter.New()
	r1.GET("/newuser", dbm.newuser)
	var NewuserTest = []struct {
		quer string
	}{
		{"{\"user_id\":1}"},
		{"{\"user_id\":2}"},
		{"{\"user_id\":3}"},
	}
	for _, tt := range NewuserTest {
		request, _ := http.NewRequest("GET", "/newuser", nil)
		result := httptest.NewRecorder()
		r1.ServeHTTP(result, request)
		if result.Body.String() != tt.quer {
			t.Errorf("Rt: Got: %s,  want: %s", result.Body.String(), tt.quer)
		}
	}
}

func TestCreateWs(t *testing.T) {
	r1 := httprouter.New()
	r1.POST("/createws", dbm.createws)
	var createwsTest = []struct {
		InitWs string
		res    string
	}{
		{InitWs: "{\"user_id\": 1, \"ws_name\": \"test ws name\"}", res: "{\"ws_id\": 1}"},
		{InitWs: "{\"user_id\": 0, \"ws_name\": \"\"}", res: "{\"ws_id\": 1}"},
		{InitWs: "{\"usr_id\": 1, \"ws_name\": \"-1\"}", res: "{\"ws_id\": 1}"},
	}
	for _, tt := range createwsTest {
		request, _ := http.NewRequest("POST", "/createws", bytes.NewBufferString(tt.InitWs))
		result := httptest.NewRecorder()
		r1.ServeHTTP(result, request)
		if result.Body.String() != tt.res {
			t.Errorf("Rt: Got: %s,  want: %s", result.Body.String(), tt.res)
		}
	}
}
