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
			t.Errorf("Got: %s,  want: %s", result.Body.String(), tt.quer)
		}
		// if result.Code != 200 {
		// 	t.Errorf("Error code ! Got: %d,  want: %d", result.Code, 200)
		// }
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
		{InitWs: "{\"user_id\": 2, \"ws_name\": \"test ws name\"}", res: "{\"ws_id\": 1}"},
		{InitWs: "{\"user_id\": 0, \"ws_name\": \"\"}", res: "{\"ws_id\": 2}"},
		// TODO: sould return an error -> no user associated
		{InitWs: "{\"usr_id\": 1, \"ws_name\": \"-1\"}", res: "{\"ws_id\": 3}"},
		// TODO: sould return an error -> no user associated
	}
	for _, tt := range createwsTest {
		request, _ := http.NewRequest("POST", "/createws", bytes.NewBufferString(tt.InitWs))
		result := httptest.NewRecorder()
		r1.ServeHTTP(result, request)
		if result.Body.String() != tt.res {
			t.Errorf("Got: %s,  want: %s", result.Body.String(), tt.res)
		}
		// if result.Code != 200 {
		// 	t.Errorf("Error code ! Got: %d,  want: %d", result.Code, 200)
		// }
	}
}

func Testws(t *testing.T) {
	r1 := httprouter.New()
	r1.POST("/ws", dbm.ws)
	var wsTest = []struct {
		InitWs string
		res    string
	}{
		{InitWs: "{\"user_id\": 1}", res: "{\"ws\": [{\"ws_id\": 1,\"ws_name\": \"test ws name\",\"user_id\": 1}]}"},
	}
	for _, tt := range wsTest {
		request, _ := http.NewRequest("POST", "/ws", bytes.NewBufferString(tt.InitWs))
		result := httptest.NewRecorder()
		r1.ServeHTTP(result, request)
		if result.Body.String() != tt.res {
			t.Errorf("Got: %s,  want: %s", result.Body.String(), tt.res)
		}
		// if result.Code != 200 {
		// 	t.Errorf("Error code ! Got: %d,  want: %d", result.Code, 200)
		// }
	}
}

func TestSave(t *testing.T) {
	r1 := httprouter.New()
	r1.POST("/save", dbm.save)
	var SaveTest = []struct {
		InitSave string
		res      string
	}{
		{InitSave: "{\"user_id\": 1,\"ws_id\": 1,\"groupes\": [{\"groupe_id\": 12,\"cards\": [{\"card_id\": 1,\"card_pub\": false,\"card\": {\"card_content\": \"{}\"}},{\"card_id\": 2,\"card_pub\": false,\"card\": {\"card_content\": \"{}\"}},{\"card_id\": 3,\"card_pub\": false,\"card\": {\"card_content\": \"{}\"}},{\"card_id\": 4,\"card_pub\": false,\"card\": {\"card_content\": \"{}\"}}]},{\"groupe_id\": 21,\"cards\": [{\"card_pub\": false,\"card_id\": 5,\"card\": {\"card_content\": \"{\\\"card_pos\\\":12}\"}}]}]}", res: "{\"saved\": true}"},
		{InitSave: "{\"user_id\": 1,\"ws_id\": 1,\"groupes\": [{\"groupe_id\": 12,\"cards\": [{\"card_id\": 1,\"card_pub\": false,\"card\": {\"card_content\": \"{}\"}},{\"card_id\": 2,\"card_pub\": false,\"card\": {\"card_content\": \"{}\"}},{\"card_id\": 3,\"card_pub\": false,\"card\": {\"card_content\": \"{}\"}},{\"card_id\": 4,\"card_pub\": false,\"card\": {\"card_content\": \"{}\"}}]},{\"groupe_id\": 21,\"cards\": [{\"card_pub\": false,\"card_id\": 5,\"card\": {\"card_content\": \"{\"card_pos\":12}\"}}]}]}", res: "{\"saved\": false, \"error\": invalid character 'c' after object key:value pair, \"code\":6.0"}},
	}
	for _, tt := range SaveTest {
		request, _ := http.NewRequest("POST", "/save", bytes.NewBufferString(tt.InitSave))
		result := httptest.NewRecorder()
		r1.ServeHTTP(result, request)
		if result.Body.String() != tt.res {
			t.Errorf("Got: %s,  want: %s", result.Body.String(), tt.res)
		}
		// if result.Code != 200 {
		// 	t.Errorf("Error code ! Got: %d,  want: %d", result.Code, 200)
		// }
	}
}

func TestLoad(t *testing.T) {
	r1 := httprouter.New()
	r1.POST("/load", dbm.load)
	var loadTest = []struct {
		InitLoad string
		res      string
	}{
		{InitLoad: "{\"user_id\": 1,\"ws_id\": 1}", res: "{ \"card_id\": [1,2,3,4,5]}"},
		{InitLoad: "{\"user_id\": 2,\"ws_id\": 1}", res: "{ \"card_id\": []}"},
	}
	for _, tt := range loadTest {
		request, _ := http.NewRequest("POST", "/load", bytes.NewBufferString(tt.InitLoad))
		result := httptest.NewRecorder()
		r1.ServeHTTP(result, request)
		if result.Body.String() != tt.res {
			t.Errorf("Got: %s,  want: %s", result.Body.String(), tt.res)
		}
		// if result.Code != 200 {
		// 	t.Errorf("Error code ! Got: %d,  want: %d", result.Code, 200)
		// }
	}
}

func TestNbCard(t *testing.T) {
	r1 := httprouter.New()
	r1.POST("/nbcard", dbm.nbcard)
	var NbCardTest = []struct {
		InitNbCard string
		res        string
	}{
		{InitNbCard: "{\"user_id\": 1,\"ws_id\": 1}", res: "{\"nb_card\": 5}"},
		{InitNbCard: "{\"user_id\": 2,\"ws_id\": 1}", res: "{\"nb_card\": 0}"},
	}
	for _, tt := range NbCardTest {
		request, _ := http.NewRequest("POST", "/nbcard", bytes.NewBufferString(tt.InitNbCard))
		result := httptest.NewRecorder()
		r1.ServeHTTP(result, request)
		if result.Body.String() != tt.res {
			t.Errorf("Got: %s,  want: %s", result.Body.String(), tt.res)
		}
		// if result.Code != 200 {
		// 	t.Errorf("Error code ! Got: %d,  want: %d", result.Code, 200)
		// }
	}
}

func TestTag(t *testing.T) {
	r1 := httprouter.New()
	r1.POST("/tag", dbm.tag)
	var TagTest = []struct {
		InitTag string
		res        string
	}{
		{InitTag: "{\"stack_id\":1, \"tag_val\":\"I am tagged\"}", res: "{\"tagged\": true}"},
		{InitTag: "{\"stack_id\":1, \"tag_val\":\"Second tag for me !\"}", res: "{\"tagged\": true}"},
	}
	for _, tt := range TagTest {
		request, _ := http.NewRequest("POST", "/tag", bytes.NewBufferString(tt.InitTag))
		result := httptest.NewRecorder()
		r1.ServeHTTP(result, request)
		if result.Body.String() != tt.res {
			t.Errorf("Got: %s,  want: %s", result.Body.String(), tt.res)
		}
		// if result.Code != 200 {
		// 	t.Errorf("Error code ! Got: %d,  want: %d", result.Code, 200)
		// }
	}
}

func TestGetTag(t *testing.T) {
	r1 := httprouter.New()
	r1.POST("/gettag", dbm.gettag)
	var GetTagTest = []struct {
		InitGetTag string
		res        string
	}{
		{InitGetTag: "{\"stack_id\":1}", res: "{\"tags:\":[\"I am tagged\",\"Second tag for me !\"]}"},
	}
	for _, tt := range GetTagTest {
		request, _ := http.NewRequest("POST", "/gettag", bytes.NewBufferString(tt.InitGetTag))
		result := httptest.NewRecorder()
		r1.ServeHTTP(result, request)
		if result.Body.String() != tt.res {
			t.Errorf("Got: %s,  want: %s", result.Body.String(), tt.res)
		}
		// if result.Code != 200 {
		// 	t.Errorf("Error code ! Got: %d,  want: %d", result.Code, 200)
		// }
	}
}
