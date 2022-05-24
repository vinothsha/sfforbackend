package userprofiledetails

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sha/cassession"
	"sha/signup"
	"strings"
)

func UserProfileDetails(w http.ResponseWriter, r *http.Request) {
	var TakeProfileDetails UserProfile
	reqdata, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
		fmt.Println("error while read userprofile function in UserProfileDetails")
	}
	json.Unmarshal(reqdata, &TakeProfileDetails)
	// fmt.Println(GetProfileDetails)
	// Get Usermail/Mobile From the Browser Cokkie
	tokenCookie, _ := r.Cookie("token")
	a := string(tokenCookie.Value)

	base64Text := make([]byte, base64.URLEncoding.DecodedLen(len(strings.Split(a, ".")[1])))
	base64.StdEncoding.Decode(base64Text, []byte(strings.Split(a, ".")[1]))
	str := string(base64Text)
	findcolon := strings.Index(str, ":")
	findcomma := strings.Index(str, ",")
	EmailOrMobile := str[findcolon+2 : findcomma-1]
	// EmailOrMobile := "+919655373273"
	if signup.ValidateEmail(EmailOrMobile) {
		GetEmailUid := cassession.Session.Query("select profileuid from userprofiledetails where email=? allow filtering", EmailOrMobile)
		GetEmailUid.Scan(&TakeProfileDetails.ProfileUid)
		fmt.Println(TakeProfileDetails.ProfileUid)
		var GetMobileDb string
		getemail := cassession.Session.Query("select mobile from userprofiledetails where mobile=? allow filtering", TakeProfileDetails.Mobile)
		getemail.Scan(&GetMobileDb)
		fmt.Println(GetMobileDb)
		fmt.Println(TakeProfileDetails.Mobile)
		if GetMobileDb != TakeProfileDetails.Mobile {
			if err = cassession.Session.Query("update userprofiledetails set firstname=?,lastname=?,dateofbirth=?,gender=?,country=?,state=?,mobile=?,countrycode=? where profileuid=?",
				TakeProfileDetails.FirstName, TakeProfileDetails.LastName, TakeProfileDetails.DateOfBirth, TakeProfileDetails.Gender, TakeProfileDetails.Country,
				TakeProfileDetails.State, TakeProfileDetails.Mobile, TakeProfileDetails.CountryCode, TakeProfileDetails.ProfileUid).Exec(); err != nil {
				fmt.Println("error while update profile details")
				fmt.Println(err)
			}
		} else {
			p := signup.Result{Status: false, Message: "mobile already used by another user"}
			json.NewEncoder(w).Encode(p)
			fmt.Println("mobile already used by another account")
		}

		// GetProfileDetails.Mobile = "null"
	} else if signup.ValidateMobile(EmailOrMobile) {
		GetEmailUid := cassession.Session.Query("select profileuid from userprofiledetails where mobile=? allow filtering", EmailOrMobile)
		GetEmailUid.Scan(&TakeProfileDetails.ProfileUid)
		fmt.Println(TakeProfileDetails.ProfileUid)
		var GetEmailDb string
		getemail := cassession.Session.Query("select email from userprofiledetails where email=? allow filtering", TakeProfileDetails.Email)
		getemail.Scan(&GetEmailDb)
		fmt.Println(GetEmailDb)
		fmt.Println(TakeProfileDetails.Email)
		if GetEmailDb != TakeProfileDetails.Email {
			if err = cassession.Session.Query("update userprofiledetails set firstname=?,lastname=?,dateofbirth=?,gender=?,country=?,state=?,email=? where profileuid=?",
				TakeProfileDetails.FirstName, TakeProfileDetails.LastName, TakeProfileDetails.DateOfBirth, TakeProfileDetails.Gender, TakeProfileDetails.Country,
				TakeProfileDetails.State, TakeProfileDetails.Email, TakeProfileDetails.ProfileUid).Exec(); err != nil {
				fmt.Println("error while update profile details")
				fmt.Println(err)
			}
		} else {
			p := signup.Result{Status: false, Message: "memail already used in another user"}
			json.NewEncoder(w).Encode(p)
			fmt.Println("email already used in another account")
		}
	}
}

//"countrycode":"+91","dateofbirth":"22-10-1997","firstname":"vinoth","lastname":"sha","gender":"male","country":"india","state":"tamilnadu"
