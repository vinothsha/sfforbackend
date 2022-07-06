package signup

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"

	"sha/cassession"
	e "sha/commonservices/commonfunctions"
	otp "sha/commonservices/otpservices"
	s "sha/commonstruct"
)

func OtpSenderInputverifier(w http.ResponseWriter, r *http.Request) {
	var UserData s.Otp
	var UniqueId string = uuid.New().String()
	reqData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("error while read data from user when Email_mobile_Verifier")
	}
	json.Unmarshal(reqData, &UserData)
	if e.ValidateEmail(UserData.Email) && UserData.Email != "" {
		var mailcheck s.Otp
		result := cassession.Session.Query("select usermail from signup where usermail=? allow filtering", UserData.Email)
		result.Scan(&mailcheck.Email)
		if UserData.Email != mailcheck.Email {
			var OtpAlreadyIn s.Otp
			result1 := cassession.Session.Query("select usermail from otp where usermail=? allow filtering", UserData.Email)
			result1.Scan(&OtpAlreadyIn.Email)
			if OtpAlreadyIn.Email == UserData.Email {
				fmt.Println("OTP Already Sent to your Email Enter the Received OTP or with for 5mintues")
			} else {
				var genOtp = e.RandomGenerater()
				otp.SendOtpToEmail(UserData.Email, genOtp)
				if err := cassession.Session.Query("insert into otp(uid,usermail,otp)values(?,?,?)USING TTL 300", UniqueId, UserData.Email, genOtp).Exec(); err != nil {
					fmt.Println("error while insert otp to OTP Table")
					fmt.Println(err)
				}
			}
			p := s.ErrorResult{Status: true, Message: "enter your OTP received by your email"}
			json.NewEncoder(w).Encode(p)
		} else {
			w.WriteHeader(http.StatusAlreadyReported)
			p := s.ErrorResult{Status: false, Message: "User already have a acoount"}
			json.NewEncoder(w).Encode(p)

		}
	} else if e.ValidateMobile(UserData.Mobile) && UserData.Mobile != "" {

		var mailcheck s.Otp
		result := cassession.Session.Query("select mobile from signup where mobile=? allow filtering", UserData.Mobile)
		result.Scan(&mailcheck.Mobile)
		if UserData.Mobile != mailcheck.Mobile {

			var OtpAlreadyIn s.Otp
			result1 := cassession.Session.Query("select mobile from otp where mobile=? allow filtering", UserData.Mobile)
			result1.Scan(&OtpAlreadyIn.Mobile)
			result1.Scan(&OtpAlreadyIn.Email)
			if OtpAlreadyIn.Mobile == UserData.Mobile {

				fmt.Println("OTP Already Sent to your Mobile Enter the Received OTP or with for 5mintues")
			} else {
				var genOtp = e.RandomGenerater()
				otp.SendOtpToMobile(UserData.Mobile, genOtp) //UserData.CountryCode+
				if err := cassession.Session.Query("insert into otp(uid,mobile,countrycode,otp)values(?,?,?,?)USING TTL 300", UniqueId, UserData.Mobile, UserData.CountryCode, genOtp).Exec(); err != nil {
					fmt.Println("error while insert otp to OTP Table")
					fmt.Println(err)
				}
			}
			p := s.ErrorResult{Status: true, Message: "enter your OTP received by your Mobile"}
			json.NewEncoder(w).Encode(p)
		} else {
			w.WriteHeader(http.StatusAlreadyReported)
			p := s.ErrorResult{Status: false, Message: "user already SignUp"}
			json.NewEncoder(w).Encode(p)
		}
	} else {
		w.WriteHeader(http.StatusNotAcceptable)
		p := s.ErrorResult{Status: false, Message: "invalid mobile or email"}
		json.NewEncoder(w).Encode(p)
	}

}
