package commentmoderation

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sha/cassession"
	"time"

	"github.com/gocql/gocql"
)

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	var newEvent Cassschema
	var kafkaEvent Proschema
	Datetime := time.Now().Format("2006-01-02 15:04:05")
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
	}

	json.Unmarshal(reqBody, &newEvent)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newEvent)

	json.Unmarshal(reqBody, &kafkaEvent)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(kafkaEvent)

	var err1 = gocql.UUID.String(newEvent.CommentId)
	var err2 = gocql.UUID.String(newEvent.CreatorId)
	if err1 != "00000000-0000-0000-0000-000000000000" && err2 != "00000000-0000-0000-0000-000000000000" {
		//Push data in to producer
		comment, _ := json.Marshal(kafkaEvent)
		mainPro(comment)
	}
	if err = cassession.Session.Query("insert into comments (creatorid,videoid,commentid,comment,polarity,datetime )Values(?,?,?,?,?,?);", newEvent.CreatorId, newEvent.Videoid, newEvent.CommentId, newEvent.Comment, newEvent.Polarity, Datetime).Exec(); err != nil {
		fmt.Println("error while insert data into comment table")
		fmt.Println(err)

	}

}
