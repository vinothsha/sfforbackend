package recentlywatched

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sha/cassession"

	s "sha/commonstruct"

	"github.com/gocql/gocql"
)

type GetRecentWatch struct {
	VideouID  gocql.UUID `json:"videouid"`
	UserId    gocql.UUID `json:"useruid"`
	VideoLink string     `json:"videolink"`
	// VideoLength   string `json:"videolength"`
	VideoSizeInMb   float64  `json:"videosizeinmb"`
	Title           string   `json:"title"`
	Description     string   `json:"description"`
	Language        string   `json:"language"`
	Genres          []string `json:"genres"`
	AgeGroup        string   `json:"agegroup"`
	Mail            string   `json:"mail"`
	Tags            []string `json:"tags"`
	Thumnail        string   `json:"thumnail"`
	Createddatetime string   `json:"createddatetime"`
	// Lastupdatedatetime string `json:"lastupdatedatetime"`
}

// type GetUserId struct {
// 	Userid  gocql.UUID `json:"userid"`
// 	VideoId gocql.UUID `json:"videoid"`
// }

func GiveRecentlyWatchedVideosToFront(w http.ResponseWriter, r *http.Request) {
	var getid s.GetUserId
	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("error while userid for GiveRecentlyWatchedVideosToFront")
	}
	json.Unmarshal(req, &getid)
	fmt.Println(getid.Userid)
	var s gocql.UUID
	iter1 := cassession.Session.Query("select videoid from recentlywatched where userid=? allow filtering;", getid.Userid).Iter()
	var AllRecentWatched []GetRecentWatch
	m := map[string]interface{}{} //make(map[string]interface{})

	for iter1.Scan(&s) {
		iter := cassession.Session.Query("SELECT * FROM videos where videouid=?", s).Iter()
		for iter.MapScan(m) {
			AllRecentWatched = append(AllRecentWatched, GetRecentWatch{
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

	}
	Conv, _ := json.MarshalIndent(AllRecentWatched, "", "  ")
	fmt.Fprintf(w, "%s", string(Conv))
	fmt.Println("finish")

}
