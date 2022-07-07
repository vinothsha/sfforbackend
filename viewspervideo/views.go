package viewspervideo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"sha/cassession"

	// "log"
	"net/http"
	"time"

	"github.com/gocql/gocql"
	// "github.com/gorilla/mux"
)

type View struct {
	Videoid  gocql.UUID `json:"videoid"`
	Userid   gocql.UUID `json:"userid"`
	Datetime string     `json:"datetime"`
	Views    int        `json:"views"`
}
type GetView struct {
	Videoid gocql.UUID `json:"videoid"`
}
type Result struct {
	Views int `json:"views"`
}

func Views(w http.ResponseWriter, r *http.Request) {
	var Event View
	Event.Datetime = time.Now().Format("2006-01-02 15:04:05")
	Event.Views = 1
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Enter Data")
	}
	json.Unmarshal(reqBody, &Event)
	var Views int
	if err := cassession.Session.Query("SELECT views from views WHERE videoid = ?", Event.Videoid).Scan(&Views); err != nil {
		cassession.Session.Query("UPDATE views SET views = ?, userid[?] = ? WHERE videoid = ?;", Event.Userid, Event.Datetime, Event.Views, Event.Videoid).Exec()
	}
	x := Views + 1
	if err := cassession.Session.Query("UPDATE views SET userid[?] = ? , views = ? WHERE videoid = ? ;", Event.Userid, Event.Datetime, x, Event.Videoid).Exec(); err != nil {
		fmt.Println("views not updated")
		fmt.Println(err)
	}
	var views int
	if err = cassession.Session.Query("SELECT views from views WHERE videoid = ?", Event.Videoid).Scan(&views); err != nil {
		fmt.Println(err)
	}
	var Res Result
	Res.Views = views
	json.NewEncoder(w).Encode(Res)
	// w.WriteHeader(http.StatusCreated)
	fmt.Println("Data processed sucessfully")
}

// func GetViews(w http.ResponseWriter, r *http.Request) {
// 	var Views int
// 	var GEvent GetView
// 	var Res Result
// 	reqBody, err := ioutil.ReadAll(r.Body)
// 	if err != nil {
// 		fmt.Println("Enter Data")
// 	}
// 	json.Unmarshal(reqBody, &GEvent)
// 	if err = cassession.Session.Query("SELECT views from views WHERE videoid = ?", GEvent.Videoid).Scan(&Views); err != nil {
// 		fmt.Println(err)
// 	}
// 	Res.Views = Views
// 	json.NewEncoder(w).Encode(Res)
// }
