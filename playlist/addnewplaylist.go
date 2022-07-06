package playlist

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sha/cassession"
	"time"

	"github.com/gocql/gocql"
	"github.com/google/uuid"
)

type Addplaylist struct {
	Addplaylistid gocql.UUID `json:"addplaylistid"`
	Videoid       gocql.UUID `json:"videoid"`
	Userid        gocql.UUID `json:"userid"`
	Playlistname  string     `json:"playlistname"`
	Datetime      string     `json:"datetime"`
}
type Result struct {
	Status  bool   `json:"ststus"`
	Message string `json:"message"`
}

func Addplaylistid(w http.ResponseWriter, r *http.Request) {
	var addplaylist Addplaylist
	// datetime
	addplaylist.Datetime = time.Now().Format("2006-01-02 15:04:05")
	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)

	}
	json.Unmarshal(req, &addplaylist)
	// var Videoid gocql.UUID
	// res := Session.Query("select videoid from addplaylist where videoid=? and userid=? allow filtering;", addplaylist.Videoid, addplaylist.Userid)
	// res.Scan(&Videoid)
	// var Userid gocql.UUID
	// res = Session.Query("select userid from addplaylist where videoid=? and userid=? allow filtering:", addplaylist.Videoid, addplaylist.Userid)
	// res.Scan(&Userid)
	// fmt.Println("success")
	// var Playlistname string
	// res = Session.Query("select playlistname from addplaylist where playlistname =?and userid=? allow filtering:")
	// res.Scan(&Playlistname)
	// addplaylistid
	addplaylist.Addplaylistid = gocql.UUID(uuid.New())

	if err = cassession.Session.Query("insert into addplaylist(addplaylistid,videoid,userid,playlistname,datetime)VALUES(?,?,?,?,?);", addplaylist.Addplaylistid, addplaylist.Videoid, addplaylist.Userid, addplaylist.Playlistname, addplaylist.Datetime).Exec(); err != nil {
		fmt.Println("error")
		fmt.Println(err)
		// fmt.Println("success")

	} else {
		fmt.Println("already saved ")
	}
	p := Result{Status: true, Message: "This video is saved"}
	json.NewEncoder(w).Encode(p)

	fmt.Println("working")
}
