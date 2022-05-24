package videotofront

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sha/cassession"

	"github.com/gocql/gocql"
)

func VideoToHomePage(w http.ResponseWriter, r *http.Request) {
	// this function send videos to the front End
	var AllVides []GetVideo
	m := map[string]interface{}{} //make(map[string]interface{})
	iter := cassession.Session.Query("SELECT * FROM videos").Iter()
	for iter.MapScan(m) {
		AllVides = append(AllVides, GetVideo{
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
}
