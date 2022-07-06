package ratings

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sha/cassession"

	"github.com/gocql/gocql"
)

type Rate struct {
	Videoid gocql.UUID `json:"videoid"`
	Userid  gocql.UUID `json:"userid"`
	Rating  int        `json:"rating"`
}

type Result1 struct {
	Rating string `json:"rating"`
	// UserId gocql.UUID `json:"userid"`
}
type Result struct {
	Videoid gocql.UUID         `json:"videoid"`
	Rating  string             `json:"rating"`
	Ratings map[gocql.UUID]int `json:"userid"`
}

func Rating(w http.ResponseWriter, r *http.Request) {
	var Event Rate
	var Res Result1
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Enter Data")
	}
	json.Unmarshal(reqBody, &Event)
	if err = cassession.Session.Query("UPDATE ratings SET rating[?] = ? WHERE videoid = ?", Event.Userid, Event.Rating, Event.Videoid).Exec(); err != nil {
		fmt.Println("error in updating e]rating")
		fmt.Println(err)
	}
	var Rating map[gocql.UUID]int
	if err := cassession.Session.Query("select rating from ratings where videoid = ?;", Event.Videoid).Scan(&Rating); err != nil {
		fmt.Println("Error while getting")
		fmt.Println(err)
	}
	var n float32
	for i, v := range Rating {
		n = n + float32(v)
		i.Clock()
	}
	z := fmt.Sprintf("%.1f", (n / float32(len(Rating))))
	Res.Rating = z
	json.NewEncoder(w).Encode(Res)
}

// func Getrate(w http.ResponseWriter, r *http.Request) {
// 	var Res Result
// 	var Rating map[gocql.UUID]int
// 	var Videoid gocql.UUID
// 	var z string
// 	iter := cassession.Session.Query("SELECT * FROM videos").Iter()
// 	for iter.Scan(&Videoid, &Rating) {
// 		var n float32
// 		for i, v := range Rating {
// 			n = n + float32(v)
// 			i.Clock()
// 		}
// 		z = fmt.Sprintf("%.1f", (n / float32(len(Rating))))
// 	}
// 	Res.Videoid = Videoid
// 	Res.Rating = z
// 	Res.Ratings = Rating
// 	json.NewEncoder(w).Encode(Res)
// }
