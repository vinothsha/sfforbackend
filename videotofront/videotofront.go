package videotofront

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sha/cassession"
	"strconv"
	s "sha/commonstruct"
	
	"github.com/gocql/gocql"
)

func VideoToHomePage(w http.ResponseWriter, r *http.Request) {
	// this function send videos to the front End
	var AllVides []s.GetVideo
	var n int
	var z string
	var LikesCount string
	var Viewsc string
	var Views string
	var Ratings string
	var Rating map[gocql.UUID]int
	var Videoidv gocql.UUID
	var Videoidl gocql.UUID
	var Videoidr gocql.UUID
	var Likes map[gocql.UUID]string
	var UserIdp gocql.UUID
	var Profile string
	var ProfileImage string
	m := map[string]interface{}{} //make(map[string]interface{})
	iter := cassession.Session.Query("SELECT * FROM videos").Iter()
	for iter.MapScan(m) {
		LikesCount = "0"
		Viewsc = "0"
		Ratings = "0"
		n = 0
		iter := cassession.Session.Query("SELECT videoid,views FROM views").Iter()
		for iter.Scan(&Videoidv, &Views) {
			if m["videouid"] == Videoidv {
				Viewsc = Views
			}
			// fmt.Println("v", Viewsc)
		}
		iter1 := cassession.Session.Query("SELECT * FROM likes;").Iter()
		for iter1.Scan(&Videoidl, &Likes) {
			if m["videouid"] == Videoidl {
				LikesCount = strconv.Itoa(len(Likes))
			}
			// fmt.Println("l", LikesCount)
		}
		iter2 := cassession.Session.Query("SELECT * FROM ratings;").Iter()
		for iter2.Scan(&Videoidr, &Rating) {
			if m["videouid"] == Videoidr {
				for i, v := range Rating {
					n = n + v
					i.Clock()
				}
				var n float32
				for i, v := range Rating {
					n = n + float32(v)
					i.Clock()
				}
				z = fmt.Sprintf("%.1f", (n / float32(len(Rating))))
				Ratings = z
			}
			// fmt.Println("r", Ratings)
		}
		iter3 := cassession.Session.Query("SELECT profileimage,useruid FROM userprofiledetails;").Iter()
		for iter3.Scan(&Profile, &UserIdp) {
			if m["useruid"] == UserIdp {
				ProfileImage = Profile
			}
		}

		AllVides = append(AllVides, s.GetVideo{
			VideouID:        m["videouid"].(gocql.UUID),
			VideoLink:       m["videolink"].(string),
			Thumnail:        m["thumnail"].(string),
			VideoSizeInMb:   m["videosizeinmb"].(float64),
			Title:           m["title"].(string),
			Description:     m["description"].(string),
			Language:        m["language"].(string),
			Genres:          m["genres"].(string),
			AgeGroup:        m["agegroup"].(string),
			UserId:          m["useruid"].(gocql.UUID),
			Tags:            m["tags"].(string),
			LikesCount:      LikesCount,
			Viewsc:          Viewsc,
			Ratings:         Ratings,
			ProfileImage:    ProfileImage,
			Createddatetime: m["createddatetime"].(string),
		})
		m = map[string]interface{}{}
	}

	Conv, _ := json.MarshalIndent(AllVides, "", "  ")
	fmt.Fprintf(w, "%s", string(Conv))
}