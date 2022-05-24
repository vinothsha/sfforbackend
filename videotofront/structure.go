package videotofront

import "github.com/gocql/gocql"

type GetVideo struct {
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
	Tags            []string `json:"tags"`
	Thumnail        string   `json:"thumnail"`
	Createddatetime string   `json:"createddatetime"`
	// Lastupdatedatetime string `json:"lastupdatedatetime"`
}
