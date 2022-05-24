package videolikes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sha/cassession"
	"sha/signup"

	"github.com/gocql/gocql"
	"github.com/google/uuid"
)

type VideoLikes struct {
	LikeOrNot bool `json:"likeornot"`
	//below details from Db
	Videouid gocql.UUID `json:"videouid"`
	Useruid  gocql.UUID `json:"useruid"`
}

func VideoLikesEndPoint(w http.ResponseWriter, r *http.Request) {
	var Like VideoLikes
	reqData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("error while read like data")
	}
	json.Unmarshal(reqData, &Like)
	var userid gocql.UUID
	var vuid gocql.UUID
	takeuseruid := cassession.Session.Query("select useruid from likes where useruid=? allow filtering", Like.Useruid)
	takeuseruid.Scan(&userid)
	takevideouid := cassession.Session.Query("select videouid from likes where videouid=? allow filtering", Like.Videouid)
	takevideouid.Scan(&vuid)
	if Like.Useruid != userid || Like.Videouid != vuid {
		if err := cassession.Session.Query("insert into likes(likeuid,useruid,videouid)Values(?,?,?)", gocql.UUID(uuid.New()), Like.Useruid, Like.Videouid).Exec(); err != nil {
			fmt.Println("error while insert data into likes table")
			fmt.Println(err)
		}
		p := signup.Result{Status: true, Message: "Video Liked Successfully"}
		json.NewEncoder(w).Encode(p)
	} else {
		var likeuid gocql.UUID
		res := cassession.Session.Query("select likeuid from likes where useruid=? and videouid=? allow filtering", Like.Useruid, Like.Videouid)
		res.Scan(&likeuid)
		if err := cassession.Session.Query("delete from likes where likeuid=?", likeuid).Exec(); err != nil {
			fmt.Println("error while delete like data into likes table")
			fmt.Println(err)
		}
		p := signup.Result{Status: true, Message: "Video Like removed Successfully"}
		json.NewEncoder(w).Encode(p)
	}
	//Get Usermail/Mobile From the Browser Cokkie
	// tokenCookie, _ := r.Cookie("token")
	// a := string(tokenCookie.Value)

	// base64Text := make([]byte, base64.URLEncoding.DecodedLen(len(strings.Split(a, ".")[1])))
	// base64.StdEncoding.Decode(base64Text, []byte(strings.Split(a, ".")[1]))
	// str := string(base64Text)
	// findcolon := strings.Index(str, ":")
	// findcomma := strings.Index(str, ",")
	// EmailOrMobile := str[findcolon+2 : findcomma-1]
	// if signup.ValidateEmail(EmailOrMobile) {
	// takeuseruid := cassession.Session.Query("select uid from signup where usermail=? allow filtering", EmailOrMobile)
	// takeuseruid.Scan(&Like.Useruid)
	// takevideouid := cassession.Session.Query("select videouid from videos where videolink=? allow filtering", Like.VideoLink)
	// takevideouid.Scan(&Like.Videouid)
	// fmt.Println(Like.Useruid)
	// fmt.Println(Like.Videouid)
	// if Like.LikeOrNot {
	// 	if err := cassession.Session.Query("insert into likes(likeuid,useruid,videouid)Values(?,?,?)allow filtering;", gocql.UUID(uuid.New()), Like.Useruid, Like.Videouid); err != nil {
	// 		fmt.Println("error while insert data into likes table")
	// 		fmt.Println(err)
	// 	}
	// 	fmt.Println("like success")
	// }
	// }
}
