package videolikes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sha/cassession"
	"strings"

	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
)

type Videores struct {
	Videoid gocql.UUID `json:"videoid"`
	Userid  gocql.UUID `json:"userid"`
}

type Result struct {
	Likescount int  `json:"likescount"`
	Like       bool `json:"like"`
}

func LikedByThatUser(w http.ResponseWriter, r *http.Request) {
	var Res Result
	var x bool = false
	var UpEvent Videores
	var Likes map[gocql.UUID]string
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Enter Data")
	}
	json.Unmarshal(reqBody, &UpEvent)
	if err := cassession.Session.Query("SELECT likes FROM likes WHERE videoid = ?", UpEvent.Videoid).Scan(&Likes); err != nil {
		fmt.Println("Error")
		fmt.Println(err)
	}
	for k, d := range Likes {
		strings.Split(d, "")
		if UpEvent.Userid == k {
			x = true
		}
	}
	Res.Likescount = (len(Likes))
	Res.Like = x
	json.NewEncoder(w).Encode(Res)
}
func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/videores", GetLikes).Methods("POST")
	// router.HandleFunc("/gettop", GetTop).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
