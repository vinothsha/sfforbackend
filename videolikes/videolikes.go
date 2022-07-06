package videolikes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sha/cassession"
	"strings"
	"time"

	"github.com/gocql/gocql"
)

type Like struct {
	Videoid  gocql.UUID `json:"videoid"`
	Likes    gocql.UUID `json:"userid"`
	Datetime string     `json:"datetime"`
}
type ResultLikes struct {
	Likescount int  `json:"likescount"`
	Like       bool `json:"like"`
}
type GetLike struct {
	Videoid gocql.UUID `json:"videoid"`
}
type GetRes struct {
	Likescount int          `json:"likescount"`
	Users      []gocql.UUID `json:"users"`
}

func AddLikes(w http.ResponseWriter, r *http.Request) {
	var UpEvent Like
	var x bool = false
	UpEvent.Datetime = time.Now().Format("2006-01-02 15:04:05")
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Enter Data")
	}
	json.Unmarshal(reqBody, &UpEvent)
	var Videoid gocql.UUID
	var Likes map[gocql.UUID]string
	n := 0
	v := 0
	iter := cassession.Session.Query("SELECT * from likes WHERE videoid = ?", UpEvent.Videoid).Iter()
	for iter.Scan(&Videoid, &Likes) {
		n += 1
	}
	if n == 0 {
		if err := cassession.Session.Query("UPDATE likes SET likes[?] = ? WHERE videoid = ?", UpEvent.Likes, UpEvent.Datetime, UpEvent.Videoid).Exec(); err != nil {
			fmt.Println("Error while updating")
			fmt.Println(err)
		}
	}
	if n != 0 {
		iter := cassession.Session.Query("SELECT * from likes WHERE videoid = ?", UpEvent.Videoid).Iter()
		for iter.Scan(&Videoid, &Likes) {
			for k, z := range Likes {
				if k == UpEvent.Likes {
					v += 1
					strings.Split(z, "")
				} else {
					v += 0
				}
			}
		}
		// fmt.Println(v)
		if v == 0 {
			if err := cassession.Session.Query("UPDATE likes SET likes[?] = ? WHERE videoid = ?", UpEvent.Likes, UpEvent.Datetime, UpEvent.Videoid).Exec(); err != nil {
				fmt.Println("Error while updating")
				fmt.Println(err)
			}
		}
		if v != 0 {
			if err := cassession.Session.Query("DELETE likes[?] FROM likes WHERE videoid = ?", UpEvent.Likes, UpEvent.Videoid).Exec(); err != nil {
				fmt.Println("Error while updating")
				fmt.Println(err)
			}
		}
	}
	if err := cassession.Session.Query("SELECT likes FROM likes WHERE videoid = ?", UpEvent.Videoid).Scan(&Likes); err != nil {
		return
	}
	for k, v := range Likes {
		strings.Split(v, "")
		if k == UpEvent.Likes {
			x = true
		}
	}
	var Res ResultLikes
	Res.Like = x
	Res.Likescount = len(Likes)
	json.NewEncoder(w).Encode(Res)
	w.WriteHeader(http.StatusCreated)
	fmt.Println("Data processed sucessfully")
}

func GetLikes(w http.ResponseWriter, r *http.Request) {
	var Event GetLike
	var Res GetRes
	var Likes map[gocql.UUID]string
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Enter Data")
	}
	json.Unmarshal(reqBody, &Event)

	if err := cassession.Session.Query("SELECT likes FROM likes WHERE videoid = ?", Event.Videoid).Scan(&Likes); err != nil {
		fmt.Println("Error in getting likes")
		fmt.Println(err)
	}
	Res.Likescount = len(Likes)
	var d []gocql.UUID
	for v, k := range Likes {
		d = append(d, v)
		strings.Split(k, "")
	}
	json.NewEncoder(w).Encode(Res)
}