package recentlywatched

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sha/cassession"
	"sha/signup"
	"time"

	"github.com/gocql/gocql"
	"github.com/google/uuid"
)

type RecentlyWatched struct {
	RecentlyWatchedUid gocql.UUID `json:"recentlywatcheduid"`
	UserId             gocql.UUID `json:"userid"`
	VideoId            gocql.UUID `json:"videoid"`
	DateTime           string     `json:"datetime"`
}

func RecentlyWatchedVideos(w http.ResponseWriter, r *http.Request) {
	var Watched RecentlyWatched

	Watched.DateTime = time.Now().Format("2006-01-02 15:04:05")
	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		fmt.Println("error while read recentlywatched data")
	}
	json.Unmarshal(req, &Watched)
	var vidDb gocql.UUID
	res := cassession.Session.Query("select videoid from recentlywatched where videoid=? and userid=? allow filtering;", Watched.VideoId, Watched.UserId)
	res.Scan(&vidDb)
	var uidDb gocql.UUID
	res = cassession.Session.Query("select userid from recentlywatched where videoid=? and userid=? allow filtering;", Watched.VideoId, Watched.UserId)
	res.Scan(&uidDb)
	res = cassession.Session.Query("select useruid from videos where videouid=? and useruid=? allow filtering", Watched.VideoId, Watched.UserId)
	var owncontentView gocql.UUID
	res.Scan(&owncontentView)
	if (Watched.VideoId != vidDb || Watched.UserId != uidDb) && Watched.UserId != owncontentView {
		Watched.RecentlyWatchedUid = gocql.UUID(uuid.New())
		if err = cassession.Session.Query("insert into RecentlyWatched(recentlywatcheduid ,userid ,videoid ,datetime)Values(?,?,?,?);", Watched.RecentlyWatchedUid, Watched.UserId, Watched.VideoId, Watched.DateTime).Exec(); err != nil {
			fmt.Println("error while insert data into recentlywatched table")
			fmt.Println(err)
		}
		p := signup.Result{Status: true, Message: "This video is add to recently watched list"}
		json.NewEncoder(w).Encode(p)
	} else {
		// fmt.Println("owncontentwatch or already watched")
		p := signup.Result{Status: false, Message: "own content watch or already watched"}
		json.NewEncoder(w).Encode(p)
	}
}
