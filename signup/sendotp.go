package signup

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/smtp"
	"net/url"
	"regexp"
	"strings"

	"github.com/dongri/phonenumber"
	"github.com/google/uuid"

	"sha/cassession"
)

func OtpSenderInputverifier(w http.ResponseWriter, r *http.Request) {
	var UserData Otp
	var UniqueId string = uuid.New().String()
	reqData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("error while read data from user when Email_mobile_Verifier")
	}
	json.Unmarshal(reqData, &UserData)
	if ValidateEmail(UserData.Email) && UserData.Email != "" {
		var mailcheck Otp
		result := cassession.Session.Query("select usermail from signup where usermail=? allow filtering", UserData.Email)
		result.Scan(&mailcheck.Email)
		if UserData.Email != mailcheck.Email {
			var OtpAlreadyIn Otp
			result1 := cassession.Session.Query("select usermail from otp where usermail=? allow filtering", UserData.Email)
			result1.Scan(&OtpAlreadyIn.Email)
			if OtpAlreadyIn.Email == UserData.Email {
				fmt.Println("OTP Already Sent to your Email Enter the Received OTP or with for 5mintues")
			} else {
				var genOtp = RandomGenerater()
				SendOtpToEmail(UserData.Email, genOtp)
				if err := cassession.Session.Query("insert into otp(uid,usermail,otp)values(?,?,?)USING TTL 300", UniqueId, UserData.Email, genOtp).Exec(); err != nil {
					fmt.Println("error while insert otp to OTP Table")
					fmt.Println(err)
				}
			}
			p := Result{Status: true, Message: "enter your OTP received by your email"}
			json.NewEncoder(w).Encode(p)
		} else {
			w.WriteHeader(http.StatusAlreadyReported)
			p := Result{Status: false, Message: "User already have a acoount"}
			json.NewEncoder(w).Encode(p)

		}
	} else if ValidateMobile(UserData.Mobile) && UserData.Mobile != "" {

		var mailcheck Otp
		result := cassession.Session.Query("select mobile from signup where mobile=? allow filtering", UserData.Mobile)
		result.Scan(&mailcheck.Mobile)
		if UserData.Mobile != mailcheck.Mobile {

			var OtpAlreadyIn Otp
			result1 := cassession.Session.Query("select mobile from otp where mobile=? allow filtering", UserData.Mobile)
			result1.Scan(&OtpAlreadyIn.Mobile)
			result1.Scan(&OtpAlreadyIn.Email)
			if OtpAlreadyIn.Mobile == UserData.Mobile {

				fmt.Println("OTP Already Sent to your Mobile Enter the Received OTP or with for 5mintues")
			} else {
				var genOtp = RandomGenerater()
				SendOtpToMobile(UserData.Mobile, genOtp) //UserData.CountryCode+
				if err := cassession.Session.Query("insert into otp(uid,mobile,countrycode,otp)values(?,?,?,?)USING TTL 300", UniqueId, UserData.Mobile, UserData.CountryCode, genOtp).Exec(); err != nil {
					fmt.Println("error while insert otp to OTP Table")
					fmt.Println(err)
				}
			}
			p := Result{Status: true, Message: "enter your OTP received by your Mobile"}
			json.NewEncoder(w).Encode(p)
		} else {
			w.WriteHeader(http.StatusAlreadyReported)
			p := Result{Status: false, Message: "user already SignUp"}
			json.NewEncoder(w).Encode(p)
		}
	} else {
		w.WriteHeader(http.StatusNotAcceptable)
		p := Result{Status: false, Message: "invalid mobile or email"}
		json.NewEncoder(w).Encode(p)
	}

}

func ValidateEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(email)
}
func ValidateMobile(mobile string) bool {
	number := phonenumber.Parse(mobile, "")
	re := regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)
	?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)
	return re.MatchString(number)
}

func RandomGenerater() string {
	p, _ := rand.Prime(rand.Reader, 18)
	fmt.Println(p)
	return p.String()

}
func SendOtpToEmail(email string, num string) {
	// abc := gomail.NewMessage()
	// abc.SetHeader("From", "vinothkkvs@gmail.com")
	// abc.SetHeader("To", email)
	// abc.SetHeader("Subject", "Test Email subject abc")
	// abc.SetBody("text", num)
	// a := gomail.NewDialer("smtp.gmail.com", 587, "vinothkkvs@gmail.com", "Vinothsha97$$")

	// if err := a.DialAndSend(abc); err != nil {
	// 	fmt.Println("error at email send for invalid email")
	// 	fmt.Println(err)
	// 	panic(err)
	// }
	from := "kvinothsha@gmail.com"
	password := "SfBeta1.0"

	// Receiver email address.
	to := []string{
		email,
	}

	// smtp server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.
	message := []byte("This is a test email message " + num)

	// Authentication.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Sending email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, message)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Email Sent Successfully!")
}
func SendOtpToMobile(mob string, num string) {
	accountSid := "ACac40d86f1e4383335d6e208ffe96c130"
	authToken := "0713eb9d2f5077cb7efadbd9e7fcc052"
	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + accountSid + "/Messages.json"

	// rand.Seed(time.Now().Unix())
	msgData := url.Values{}
	msgData.Set("To", mob) //vicky--9629381169 hussain--9094501317
	msgData.Set("From", "+16592013522")
	msgData.Set("Body", num)
	fmt.Println(mob)
	msgDataReader := *strings.NewReader(msgData.Encode())
	client := &http.Client{}
	req, _ := http.NewRequest("POST", urlStr, &msgDataReader)
	req.SetBasicAuth(accountSid, authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, _ := client.Do(req)
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		var data map[string]interface{}
		decoder := json.NewDecoder(resp.Body)
		err := decoder.Decode(&data)
		if err == nil {
			fmt.Println(data["sid"])
		}
	} else {
		fmt.Println(resp.Status)
	}
}
