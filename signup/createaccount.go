package signup

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sha/cassession"
	"time"

	"github.com/gocql/gocql"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func CreateAccountOtpVerify(w http.ResponseWriter, r *http.Request) {
	var AllDataFromUser CreateAccount
	reqData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("error while data read during Create_Account")
	}
	json.Unmarshal(reqData, &AllDataFromUser)
	if ValidateEmail(AllDataFromUser.Email) && AllDataFromUser.Email != "" {
		var mailcheck CreateAccount
		result := cassession.Session.Query("select otp from otp where usermail=? allow filtering", AllDataFromUser.Email)
		result.Scan(&mailcheck.OtpNumber)
		if AllDataFromUser.OtpNumber == mailcheck.OtpNumber {
			p := Result{Status: true, Message: "OTP Verified Successfully"}
			json.NewEncoder(w).Encode(p)
		} else {
			p := Result{Status: false, Message: "OTP Not Verified Successfully"}
			json.NewEncoder(w).Encode(p)
			return
		}
	} else if ValidateMobile(AllDataFromUser.Mobile) && AllDataFromUser.Mobile != "" {
		var mailcheck CreateAccount
		result := cassession.Session.Query("select otp from otp where mobile=? allow filtering", AllDataFromUser.Mobile)
		result.Scan(&mailcheck.OtpNumber)
		if AllDataFromUser.OtpNumber == mailcheck.OtpNumber {
			p := Result{Status: true, Message: "OTP Verified Successfully"}
			json.NewEncoder(w).Encode(p)
		} else {
			p := Result{Status: false, Message: "OTP Not Verified Successfully"}
			json.NewEncoder(w).Encode(p)
		}
	}
}

func PasswordEnterSignup(w http.ResponseWriter, r *http.Request) {
	var PasswordFromUser Passwd
	reqData, err := ioutil.ReadAll(r.Body)
	Createddatetime := time.Now().Format("2006-01-02 15:04:05")
	if err != nil {
		fmt.Println("error while read password from user and function is PasswordEnter_Signup")
		fmt.Println(err)
	}
	json.Unmarshal(reqData, &PasswordFromUser)
	if ValidateEmail(PasswordFromUser.Email) && PasswordFromUser.Email != "" {
		if len(PasswordFromUser.Password) >= 8 {

			hashedPass := HashPassword(PasswordFromUser.Password)
			var UserUid gocql.UUID = gocql.UUID(uuid.New())
			if err := cassession.Session.Query("insert into signup (uid,usermail,password,createddatetime)values(?,?,?,?)", UserUid, PasswordFromUser.Email, hashedPass, Createddatetime).Exec(); err != nil {
				fmt.Println("error while insert password and usermail to DB table function is PasswordEnter_Signup")
				fmt.Println(err)
			}
			if err =
				cassession.Session.Query("insert into userprofiledetails(profileuid,useruid,email)values(?,?,?);", gocql.UUID(uuid.New()), UserUid, PasswordFromUser.Email).Exec(); err != nil {
				fmt.Println(err)
				fmt.Println("error while insert email signin unserprofiledetails")
			}
			p := Result{Status: true, Message: "Account Created  Successfully Using Email Number"}
			json.NewEncoder(w).Encode(p)
		} else {
			p := Result{Status: false, Message: "password should greater than 8 character"}
			json.NewEncoder(w).Encode(p)
		}
	} else if ValidateMobile(PasswordFromUser.Mobile) && PasswordFromUser.Mobile != "" {
		if len(PasswordFromUser.Password) >= 8 {

			hashedPass := HashPassword(PasswordFromUser.Password)
			var UserUid gocql.UUID = gocql.UUID(uuid.New())
			if err := cassession.Session.Query("insert into signup (uid,mobile,countrycode,password,createddatetime)values(?,?,?,?,?)", UserUid, PasswordFromUser.Mobile, PasswordFromUser.CountryCode, hashedPass, Createddatetime).Exec(); err != nil {
				fmt.Println("error while insert password and usermail to DB table function is PasswordEnter_Signup")
				fmt.Println(err)
			}
			if err =
				cassession.Session.Query("insert into userprofiledetails(profileuid,useruid,countrycode,mobile)values(?,?,?,?);", gocql.UUID(uuid.New()), UserUid, PasswordFromUser.CountryCode, PasswordFromUser.Mobile).Exec(); err != nil {
				fmt.Println(err)
				fmt.Println("error while insert mobile signin unserprofiledetails")
			}
			p := Result{Status: true, Message: "Account Created  Successfully Using Mobile Number"}
			fmt.Println("this is a country code", PasswordFromUser.CountryCode, " this is a mobile ", PasswordFromUser.Mobile)
			json.NewEncoder(w).Encode(p)
		} else {
			p := Result{Status: false, Message: "password greater than 8 character"}
			json.NewEncoder(w).Encode(p)
		}
	} else {
		p := Result{Status: false, Message: "plz signup agin"}
		json.NewEncoder(w).Encode(p)
	}
}
func HashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}
