package commonstruct

import "github.com/gocql/gocql"

type Otp struct {
	Email       string `json:"email"`
	CountryCode string `json:"countrycode"`
	Mobile      string `json:"mobile"`
}

type CreateAccount struct {
	Email       string `json:"email"`
	CountryCode string `json:"countrycode"`
	Mobile      string `json:"mobile"`
	OtpNumber   string `json:"otp"`
}

type Passwd struct {
	Email       string `json:"email"`
	CountryCode string `json:"countrycode"`
	Mobile      string `json:"mobile"`
	Password    string `json:"password"`
}
type ErrorResult struct {
	Status  bool   `json:"content"`
	Message string `json:"message"`
	UserId  string `josn:"userid"`
}
type ResultEmail struct{
	Status  bool   `json:"content"`
	Message string `json:"message"`
	UserId  string `josn:"userid"`
	Email   string `json:"email"`
}
type ResultMobile struct{
	Status  bool   `json:"content"`
	Message string `json:"message"`
	UserId  string `josn:"userid"`
	Mobile   string `json:"mobile"`
}
type GetUserId struct {
	Userid  gocql.UUID `json:"userid"`
	VideoId gocql.UUID `json:"videoid"`
}
type RecentlyWatched struct {
	RecentlyWatchedUid gocql.UUID `json:"recentlywatcheduid"`
	UserId             gocql.UUID `json:"userid"`
	VideoId            gocql.UUID `json:"videoid"`
	DateTime           string     `json:"datetime"`
}
type UploadVideo struct {
	VideouID  gocql.UUID `json:"videouid"`
	UserId    string     `json:"useruid"`
	VideoLink string     `json:"videolink"`
	Etag      string     `json:"etag"`
	// VideoLength   string `json:"videolength"`
	VideoSizeInMb   float64 `json:"videosizeinmb"`
	Title           string  `json:"title"`
	Description     string  `json:"description"`
	Language        string  `json:"language"`
	Genres          string  `json:"geners"`
	AgeGroup        string  `json:"agegroup"`
	Tags            string  `json:"tags"`
	Createddatetime string  `json:"createddatetime"`
	Thumnail        string  `json:"thumnail"`

	// Lastupdatedatetime string `json:"lastupdatedatetime"`
}
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
	Createddatetime string  `json:"createddatetime"`
}
