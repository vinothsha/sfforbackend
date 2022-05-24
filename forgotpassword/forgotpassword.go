package forgotpassword

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sha/cassession"
	"sha/signup"

	"github.com/google/uuid"
)

var UniqueId string = uuid.New().String()

func PasswordResetOtpSender(w http.ResponseWriter, r *http.Request) {

	var PasswdResetData PasswordReset
	reqData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("error while read data from user when Email_mobile_Verifier")
	}
	json.Unmarshal(reqData, &PasswdResetData)
	if signup.ValidateEmail(PasswdResetData.Email) && PasswdResetData.Email != "" {
		p := signup.Result{Status: true, Message: "enter your OTP received by your email"}
		var mailcheck PasswordReset
		result := cassession.Session.Query("select usermail from signup where usermail=? allow filtering", PasswdResetData.Email)
		result.Scan(&mailcheck.Email)
		if PasswdResetData.Email == mailcheck.Email {
			var OtpAlreadyIn PasswordReset
			result1 := cassession.Session.Query("select usermail from otp where usermail=? allow filtering", PasswdResetData.Email)
			result1.Scan(&OtpAlreadyIn.Email)
			if OtpAlreadyIn.Email == PasswdResetData.Email {

				fmt.Println("OTP Already Sent to your Email Enter the Received OTP or with for 5 mintues")
			} else {
				var genOtp = signup.RandomGenerater()
				signup.SendOtpToEmail(PasswdResetData.Email, genOtp)

				if err := cassession.Session.Query("insert into otp(uid,usermail,otp)values(?,?,?)USING TTL 300", UniqueId, PasswdResetData.Email, genOtp).Exec(); err != nil {
					fmt.Println("error while insert otp to OTP Table")
					fmt.Println(err)
				}
				json.NewEncoder(w).Encode(p)
			}
		} else {
			w.WriteHeader(http.StatusNotAcceptable)
			p := signup.Result{Status: false, Message: "invalid mobile or email"}
			json.NewEncoder(w).Encode(p)
		}

	} else if signup.ValidateMobile(PasswdResetData.Mobile) && PasswdResetData.Mobile != "" {
		p := signup.Result{Status: true, Message: "enter your OTP received by your Mobile"}
		var mailcheck PasswordReset
		result := cassession.Session.Query("select mobile from signup where mobile=? allow filtering", PasswdResetData.Mobile)
		result.Scan(&mailcheck.Mobile)
		if PasswdResetData.Mobile == mailcheck.Mobile {

			var OtpAlreadyIn PasswordReset
			result1 := cassession.Session.Query("select mobile from otp where mobile=? allow filtering", PasswdResetData.Mobile)
			result1.Scan(&OtpAlreadyIn.Mobile)
			result1.Scan(&OtpAlreadyIn.Email)
			if OtpAlreadyIn.Mobile == PasswdResetData.Mobile {

				fmt.Println("OTP Already Sent to your Mobile Enter the Received OTP or with for 5 mintues")
			} else {
				var genOtp = signup.RandomGenerater()
				signup.SendOtpToMobile(PasswdResetData.Mobile, genOtp)
				if err := cassession.Session.Query("insert into otp(uid,mobile,countrycode,otp)values(?,?,?,?)USING TTL 300", UniqueId, PasswdResetData.Mobile, PasswdResetData.CountryCode, genOtp).Exec(); err != nil {
					fmt.Println("error while insert otp to OTP Table")
					fmt.Println(err)
				}
				json.NewEncoder(w).Encode(p)
			}

		} else {
			w.WriteHeader(http.StatusNotAcceptable)
			p := signup.Result{Status: false, Message: "invalid mobile or email"}
			json.NewEncoder(w).Encode(p)
		}

	} else {
		w.WriteHeader(http.StatusNotAcceptable)
		p := signup.Result{Status: false, Message: "invalid mobile or email"}
		json.NewEncoder(w).Encode(p)
	}

}
func ResetPasswordOtpVerify(w http.ResponseWriter, r *http.Request) {
	var AllDataFromUser OtpVerify
	reqData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("error while data read during Reset_Password_OtpVerify")
	}
	json.Unmarshal(reqData, &AllDataFromUser)
	if signup.ValidateEmail(AllDataFromUser.Email) && AllDataFromUser.Email != "" {
		var mailcheck OtpVerify
		result := cassession.Session.Query("select otp from otp where usermail=? allow filtering", AllDataFromUser.Email)
		result.Scan(&mailcheck.Otp)
		if AllDataFromUser.Otp == mailcheck.Otp {
			p := signup.Result{Status: true, Message: "OTP Verified Successfully"}
			if err := cassession.Session.Query("delete from otp where uid=?", UniqueId).Exec(); err != nil {
				fmt.Println("error while delete mobile otp")
				fmt.Println(err)
			}
			json.NewEncoder(w).Encode(p)
		} else {
			w.WriteHeader(http.StatusNotAcceptable)
			p := signup.Result{Status: false, Message: "OTP Not Verified Successfully"}
			json.NewEncoder(w).Encode(p)
			return
		}
	} else if signup.ValidateMobile(AllDataFromUser.Mobile) && AllDataFromUser.Mobile != "" {
		var mailcheck OtpVerify
		result := cassession.Session.Query("select otp from otp where mobile=? allow filtering", AllDataFromUser.Mobile)
		result.Scan(&mailcheck.Otp)
		if AllDataFromUser.Otp == mailcheck.Otp {
			p := signup.Result{Status: true, Message: "OTP Verified Successfully"}
			fmt.Println(UniqueId)
			if err := cassession.Session.Query("delete  from otp where uid=?", UniqueId).Exec(); err != nil {
				fmt.Println("error while delete mobile otp")
				fmt.Println(err)
			}
			json.NewEncoder(w).Encode(p)
		} else {
			w.WriteHeader(http.StatusNotAcceptable)
			p := signup.Result{Status: true, Message: "OTP Verified Successfully"}
			json.NewEncoder(w).Encode(p)
		}
	}
}
func EnterNewPassword(w http.ResponseWriter, r *http.Request) {
	var PasswordFromUser NewPassword
	reqData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("error while read password from user and function is EnterNewPassword")
	}
	json.Unmarshal(reqData, &PasswordFromUser)
	if signup.ValidateEmail(PasswordFromUser.Email) && PasswordFromUser.Email != "" {
		if len(PasswordFromUser.Password) >= 8 {
			p := signup.Result{Status: true, Message: "Pasword Reset Successfully"}
			hashedPass := signup.HashPassword(PasswordFromUser.Password)
			var mailcheck NewPassword
			result := cassession.Session.Query("select uid from signup where usermail=? allow filtering", PasswordFromUser.Email)
			result.Scan(&mailcheck.UniqueId)
			result = cassession.Session.Query("select usermail from signup where usermail=? allow filtering", PasswordFromUser.Email)
			result.Scan(&mailcheck.Email)
			if mailcheck.Email == PasswordFromUser.Email {
				if err := cassession.Session.Query("update signup set password=? where uid=?", hashedPass, mailcheck.UniqueId).Exec(); err != nil {
					fmt.Println("error while insert password and usermail to DB table function is PasswordEnter_Signup")
					fmt.Println(err)
				}
			}
			json.NewEncoder(w).Encode(p)
		} else {
			w.WriteHeader(http.StatusNotAcceptable)
			p := signup.Result{Status: false, Message: "password should greater than 8 character"}
			json.NewEncoder(w).Encode(p)
		}
	} else if signup.ValidateMobile(PasswordFromUser.Mobile) {
		if len(PasswordFromUser.Password) >= 8 {
			hashedPass := signup.HashPassword(PasswordFromUser.Password)
			var mailcheck NewPassword
			result := cassession.Session.Query("select uid from signup where mobile=? allow filtering", PasswordFromUser.Mobile)
			result.Scan(&mailcheck.UniqueId)
			result = cassession.Session.Query("select mobile from signup where mobile=? allow filtering", PasswordFromUser.Mobile)
			result.Scan(&mailcheck.Mobile)
			if PasswordFromUser.Mobile == mailcheck.Mobile {
				p := signup.Result{Status: true, Message: "Password Reset Successfully"}
				if err := cassession.Session.Query("update signup set password=? where uid=?", hashedPass, mailcheck.UniqueId).Exec(); err != nil {
					fmt.Println("error while insert password and usermail to DB table function is EnterNewPassword ")
					fmt.Println(err)
				}
				json.NewEncoder(w).Encode(p)
			}

		} else {
			w.WriteHeader(http.StatusNotAcceptable)
			p := signup.Result{Status: false, Message: "password greater than 8 character"}
			json.NewEncoder(w).Encode(p)
		}
	} else {
		w.WriteHeader(http.StatusNotAcceptable)
		p := signup.Result{Status: false, Message: "plz Reset Password again"}
		json.NewEncoder(w).Encode(p)
	}
}
