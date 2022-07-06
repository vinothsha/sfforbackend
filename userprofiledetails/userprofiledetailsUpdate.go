package userprofiledetails

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sha/cassession"
	e "sha/commonservices/commonfunctions"
	s "sha/commonstruct"
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
	// tokenCookie, _ := r.Cookie("token")
	// a := string(tokenCookie.Value)

	// base64Text := make([]byte, base64.URLEncoding.DecodedLen(len(strings.Split(a, ".")[1])))
	// base64.StdEncoding.Decode(base64Text, []byte(strings.Split(a, ".")[1]))
	// str := string(base64Text)
	// findcolon := strings.Index(str, ":")
	// findcomma := strings.Index(str, ",")
	// EmailOrMobile := str[findcolon+2 : findcomma-1]
	var email string
	var mobile string
	var EmailOrMobile string
	if err := cassession.Session.Query("select usermail,mobile from signup where uid=?", TakeProfileDetails.Useruid).Scan(&email, &mobile); err != nil {
		fmt.Println("error while get email/mobile in userprofile_function")
	}
	if email != "" {
		EmailOrMobile = email
	} else {
		EmailOrMobile = mobile
	}
	if e.ValidateEmail(EmailOrMobile) {
		GetEmailUid := cassession.Session.Query("select profileuid from userprofiledetails where email=? allow filtering", EmailOrMobile)
		GetEmailUid.Scan(&TakeProfileDetails.ProfileUid)
		fmt.Println(TakeProfileDetails.ProfileUid)

		if err = cassession.Session.Query("update userprofiledetails set firstname=?,lastname=?,dateofbirth=?,gender=?,mobile=?,profileimage=?,countrycode=? where profileuid=?",
			TakeProfileDetails.FirstName, TakeProfileDetails.LastName, TakeProfileDetails.DateOfBirth, TakeProfileDetails.Gender,
			TakeProfileDetails.Mobile, TakeProfileDetails.Profileimage, TakeProfileDetails.CountryCode, TakeProfileDetails.ProfileUid).Exec(); err != nil {
			fmt.Println("error while update profile details")
			fmt.Println(err)
		}
		p := s.ErrorResult{Status: true, Message: "changed mobile or other details"}
		json.NewEncoder(w).Encode(p)
		fmt.Println("changed")

	} else if e.ValidateMobile(EmailOrMobile) {
		GetEmailUid := cassession.Session.Query("select profileuid from userprofiledetails where mobile=? allow filtering", EmailOrMobile)
		GetEmailUid.Scan(&TakeProfileDetails.ProfileUid)
		fmt.Println(TakeProfileDetails.ProfileUid)
		var GetEmailDb string
		getemail := cassession.Session.Query("select email from userprofiledetails where email=? allow filtering", TakeProfileDetails.Email)
		getemail.Scan(&GetEmailDb)
		fmt.Println(GetEmailDb)
		fmt.Println(TakeProfileDetails.Email)

		if err = cassession.Session.Query("update userprofiledetails set firstname=?,lastname=?,dateofbirth=?,gender=?,profileimage=?,email=? where profileuid=?",
			TakeProfileDetails.FirstName, TakeProfileDetails.LastName, TakeProfileDetails.DateOfBirth, TakeProfileDetails.Gender, TakeProfileDetails.Profileimage,
			TakeProfileDetails.Email, TakeProfileDetails.ProfileUid).Exec(); err != nil {
			fmt.Println("error while update profile details")
			fmt.Println(err)
		}

		p := s.ErrorResult{Status: true, Message: "changed email or other details"}
		json.NewEncoder(w).Encode(p)
		fmt.Println("changed")
	} else {
		p := s.ErrorResult{Status: false, Message: "error in data "}
		json.NewEncoder(w).Encode(p)
	}
}

//"countrycode":"+91","dateofbirth":"22-10-1997","firstname":"vinoth","lastname":"sha","gender":"male","country":"india","state":"tamilnadu"
