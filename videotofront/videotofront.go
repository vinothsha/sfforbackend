package videotofront

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sha/cassession"
	"strconv"
	"strings"

	"github.com/gocql/gocql"
)

type GetVideo struct {
	VideouID  gocql.UUID `json:"videouid"`
	UserId    gocql.UUID `json:"useruid"`
	VideoLink string     `json:"videolink"`
	// VideoLength   string `json:"videolength"`
	VideoSizeInMb   float64 `json:"videosizeinmb"`
	Title           string  `json:"title"`
	Description     string  `json:"description"`
	Language        string  `json:"language"`
	Genres          string  `json:"genres"`
	AgeGroup        string  `json:"agegroup"`
	Tags            string  `json:"tags"`
	Thumnail        string  `json:"thumnail"`
	LikesCount      string  `json:"likescount"`
	Viewsc          string  `json:"views"`
	Ratings         string  `json:"ratings"`
	ProfileImage    string  `json:"profileimage"`
	Createddatetime string  `json:"createddatetime"`
	// Lastupdatedatetime string `json:"lastupdatedatetime"`
}

func VideoToHomePage(w http.ResponseWriter, r *http.Request) {
	// this function send videos to the front End
	var AllVides []GetVideo
	var n int
	var c string
	var z float32
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
		iter := cassession.Session.Query("SELECT videoid,views FROM views").Iter()
		for iter.Scan(&Videoidv, &Views) {
			if m["videouid"] == Videoidv {
				Viewsc = Views
			}
		}
		iter1 := cassession.Session.Query("SELECT * FROM likes;").Iter()
		for iter1.Scan(&Videoidl, &Likes) {
			if m["videouid"] == Videoidl {
				LikesCount = strconv.Itoa(len(Likes))
			}
		}
		iter2 := cassession.Session.Query("SELECT * FROM ratings;").Iter()
		for iter2.Scan(&Videoidr, &Rating) {
			if m["videouid"] == Videoidr {
				for i, v := range Rating {
					n = n + v
					i.Clock()
				}
				z = float32(n / len(Rating))
				c = strconv.Itoa(int(z))
				strings.Split(c, "")
				Ratings = c
			}
			strings.Split(Viewsc, "")
		}
		iter3 := cassession.Session.Query("SELECT profileimage,useruid FROM userprofiledetails;").Iter()
		for iter3.Scan(&Profile, &UserIdp) {
			if m["useruid"] == UserIdp {
				ProfileImage = Profile
			}
		}

		AllVides = append(AllVides, GetVideo{
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
