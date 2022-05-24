package userprofiledetails

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sha/cassession"

	"github.com/gocql/gocql"
	"github.com/gorilla/mux"
)

func GiveprofileDetailsToFrontEnd(w http.ResponseWriter, r *http.Request) {
	var getuid = mux.Vars(r)["id"]
	var UserProfileDetails []UserProfile
	fmt.Println(getuid)
	m := map[string]interface{}{}
	iter := cassession.Session.Query("select * from userprofiledetails where useruid=? allow filtering", getuid).Iter()
	for iter.MapScan(m) {
		fmt.Println("called")
		UserProfileDetails = append(UserProfileDetails, UserProfile{
			ProfileUid:  m["profileuid"].(gocql.UUID),
			CountryCode: m["countrycode"].(string),
			Mobile:      m["mobile"].(string),
			DateOfBirth: m["dateofbirth"].(string),
			Email:       m["email"].(string),
			FirstName:   m["firstname"].(string),
			LastName:    m["lastname"].(string),
			Gender:      m["gender"].(string),
			Country:     m["country"].(string),
			State:       m["state"].(string),
		})
		m = map[string]interface{}{}
	}
	Conv, _ := json.MarshalIndent(UserProfileDetails, "", "  ")
	fmt.Fprintf(w, "%s", string(Conv))
}
