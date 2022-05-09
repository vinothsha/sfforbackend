package signup

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sha/cassession"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var AllDataFromUser CreateAccount

func CreateAccountOtpVerify(w http.ResponseWriter, r *http.Request) {
	reqData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("error while data read during Create_Account")
	}
	json.Unmarshal(reqData, &AllDataFromUser)
	if ValidateEmail(UserData.Email) && UserData.Email != "" {
		var mailcheck CreateAccount
		result := cassession.Session.Query("select otp from otp where usermail=? allow filtering", AllDataFromUser.Email)
		result.Scan(&mailcheck.OtpNumber)
		if AllDataFromUser.OtpNumber == mailcheck.OtpNumber {
			p := Result{Status: true, Message: "OTP Verified Successfully"}
			if err := cassession.Session.Query("delete from otp where uid=?", UniqueId).Exec(); err != nil {
				fmt.Println("error while delete mobile otp")
				fmt.Println(err)
			}
			json.NewEncoder(w).Encode(p)
		} else {
			p := Result{Status: false, Message: "OTP Not Verified Successfully"}
			json.NewEncoder(w).Encode(p)
			return
		}
	} else if ValidateMobile(UserData.Mobile) && UserData.Mobile != "" {
		var mailcheck CreateAccount
		result := cassession.Session.Query("select otp from otp where mobile=? allow filtering", AllDataFromUser.Mobile)
		result.Scan(&mailcheck.OtpNumber)
		if AllDataFromUser.OtpNumber == mailcheck.OtpNumber {
			p := Result{Status: true, Message: "OTP Verified Successfully"}
			fmt.Println(UniqueId)
			if err := cassession.Session.Query("delete  from otp where uid=?", UniqueId).Exec(); err != nil {
				fmt.Println("error while delete mobile otp")
				fmt.Println(err)
			}
			json.NewEncoder(w).Encode(p)
		} else {
			p := Result{Status: false, Message: "OTP Verified Successfully"}
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
	}
	json.Unmarshal(reqData, &PasswordFromUser)
	if ValidateEmail(PasswordFromUser.Email) && PasswordFromUser.Email != "" {
		if len(PasswordFromUser.Password) >= 8 {
			p := Result{Status: true, Message: "Account Created  Successfully"}
			hashedPass := HashPassword(PasswordFromUser.Password)
			var mailcheck Otp
			result := cassession.Session.Query("select usermail from signup where usermail=? allow filtering", PasswordFromUser.Email)
			result.Scan(&mailcheck.Email)
			if UserData.Email != mailcheck.Email {
				if err := cassession.Session.Query("insert into signup (uid,usermail,password,createddatetime)values(?,?,?,?)", uuid.New().String(), PasswordFromUser.Email, hashedPass, Createddatetime).Exec(); err != nil {
					fmt.Println("error while insert password and usermail to DB table function is PasswordEnter_Signup")
					fmt.Println(err)
				}

			}
			json.NewEncoder(w).Encode(p)
		} else {
			p := Result{Status: false, Message: "password should greater than 8 character"}
			json.NewEncoder(w).Encode(p)
		}
	} else if ValidateMobile(PasswordFromUser.Mobile) {
		if len(PasswordFromUser.Password) >= 8 {
			p := Result{Status: true, Message: "Account Created  Successfully"}
			hashedPass := HashPassword(PasswordFromUser.Password)
			var mailcheck Otp
			result := cassession.Session.Query("select mobile from signup where mobile=? allow filtering", PasswordFromUser.Mobile)
			result.Scan(&mailcheck.Mobile)
			if PasswordFromUser.Mobile != mailcheck.Mobile {
				if err := cassession.Session.Query("insert into signup (uid,mobile,countrycode,password,createddatetime)values(?,?,?,?,?)", uuid.New().String(), PasswordFromUser.Mobile, PasswordFromUser.CountryCode, hashedPass, Createddatetime).Exec(); err != nil {
					fmt.Println("error while insert password and usermail to DB table function is PasswordEnter_Signup")
					fmt.Println(err)
				}
				json.NewEncoder(w).Encode(p)
			}
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
