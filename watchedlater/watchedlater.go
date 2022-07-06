package watchedlater

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sha/cassession"
	s "sha/commonstruct"
	"time"

	"github.com/gocql/gocql"
	"github.com/google/uuid"
)

type watchedlater struct {
	WatchedlaterUid gocql.UUID `json:"watchedlateruid"`
	VideouID        gocql.UUID `json:"videouid"`
	Datetime        string     `json:"datetime"`
	UseruId         gocql.UUID `json:"useruid"`
}

func Watchedlater(w http.ResponseWriter, r *http.Request) {
	var Watched watchedlater
	Watched.Datetime = time.Now().Format("2006-01-02 15:04:05")
	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}
	json.Unmarshal(req, &Watched)

	fmt.Println(Watched.UseruId)
	var vidDb gocql.UUID
	res := cassession.Session.Query("Select videoid from watchedlater where videoid=? and userid=? allow filtering;", Watched.VideouID, Watched.UseruId)
	res.Scan(&vidDb)
	// fmt.Println(vidDb)
	var uidDb gocql.UUID
	res = cassession.Session.Query("Select userid from watchedlater where videoid=? and userid=? allow filtering;", Watched.VideouID, Watched.UseruId)
	res.Scan(&uidDb)
	// fmt.Println(uidDb)
	if Watched.VideouID != vidDb || Watched.UseruId != uidDb {
		Watched.WatchedlaterUid = gocql.UUID(uuid.New())
		if err = cassession.Session.Query("insert into watchedlater(watchedlateruid,userid,videoid,datetime)VALUES(?,?,?,?); ", Watched.WatchedlaterUid, Watched.UseruId, Watched.VideouID, Watched.Datetime).Exec(); err != nil {
			fmt.Println("error")
			fmt.Println(err)
		}
	} else {
		fmt.Println("already exists")
	}
	p := s.ErrorResult{Status: true, Message: "This video is add to watchedlater "}
	json.NewEncoder(w).Encode(p)

}
