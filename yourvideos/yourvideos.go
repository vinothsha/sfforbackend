package yourvideos

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sha/cassession"
	"sha/videotofront"

	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
)

func GiveYourVideosToFront(w http.ResponseWriter, r *http.Request) {
	userid := mux.Vars(r)["id"]
	var AllVides []videotofront.GetVideo
	m := map[string]interface{}{} //make(map[string]interface{})
	iter := cassession.Session.Query("SELECT * FROM videos where useruid=? allow filtering", userid).Iter()
	for iter.MapScan(m) {
		AllVides = append(AllVides, videotofront.GetVideo{
			VideouID:        m["videouid"].(gocql.UUID),
			VideoLink:       m["videolink"].(string),
			Thumnail:        m["thumnail"].(string),
			VideoSizeInMb:   m["videosizeinmb"].(float64),
			Title:           m["title"].(string),
			Description:     m["description"].(string),
			Language:        m["language"].(string),
			Genres:          m["genres"].([]string),
			AgeGroup:        m["agegroup"].(string),
			UserId:          m["useruid"].(gocql.UUID),
			Tags:            m["tags"].([]string),
			Createddatetime: m["createddatetime"].(string),
		})
		m = map[string]interface{}{}
	}

	Conv, _ := json.MarshalIndent(AllVides, "", "  ")
	fmt.Fprintf(w, "%s", string(Conv))
	fmt.Println(len(AllVides))
}
