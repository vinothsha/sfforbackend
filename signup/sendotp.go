package signup

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"

	"github.com/dongri/phonenumber"
	"github.com/google/uuid"

	"sha/cassession"

	gomail "gopkg.in/mail.v2"
)

var UserData Otp
var UniqueId string = uuid.New().String()

func OtpSenderInputverifier(w http.ResponseWriter, r *http.Request) {
	reqData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("error while read data from user when Email_mobile_Verifier")
	}
	json.Unmarshal(reqData, &UserData)
	if ValidateEmail(UserData.Email) && UserData.Email != "" {
		p := Result{Status: true, Message: "enter your OTP received by your email"}
		var mailcheck Otp
		result := cassession.Session.Query("select usermail from signup where usermail=? allow filtering", UserData.Email)
		result.Scan(&mailcheck.Email)
		// fmt.Println(UserData.Email)
		if UserData.Email != mailcheck.Email {
			fmt.Fprintln(w, "user not present")
			var genOtp = RandomGenerater()
			var OtpAlreadyIn Otp
			result1 := cassession.Session.Query("select usermail from otp where usermail=? allow filtering", UserData.Email)
			result1.Scan(&OtpAlreadyIn.Email)
			if OtpAlreadyIn.Email == UserData.Email {
				if err := cassession.Session.Query("update otp set otp=? where uid=?", genOtp, UniqueId).Exec(); err != nil {
					fmt.Println("error while update the otp")
				}
			} else {
				if err := cassession.Session.Query("insert into otp(uid,usermail,otp)values(?,?,?)USING TTL 20", UniqueId, UserData.Email, genOtp).Exec(); err != nil {
					fmt.Println("error while insert otp to OTP Table")
					fmt.Println(err)
				}
			}
			SendOtpToEmail(UserData.Email, genOtp)
			json.NewEncoder(w).Encode(p)
		} else {
			p := Result{Status: false, Message: "User already have a acoount"}
			json.NewEncoder(w).Encode(p)

		}
	} else if ValidateMobile(UserData.Mobile) && UserData.Mobile != "" {
		p := Result{Status: true, Message: "enter your OTP received by your Mobile"}
		var mailcheck Otp
		result := cassession.Session.Query("select mobile from signup where mobile=? allow filtering", UserData.Mobile)
		result.Scan(&mailcheck.Mobile)
		if UserData.Mobile != mailcheck.Mobile {
			var genOtp = RandomGenerater()
			var OtpAlreadyIn Otp
			result1 := cassession.Session.Query("select mobile from otp where mobile=? allow filtering", UserData.Mobile)
			result1.Scan(&OtpAlreadyIn.Mobile)
			result1.Scan(&OtpAlreadyIn.Email)
			if OtpAlreadyIn.Email == UserData.Email {
				if err := cassession.Session.Query("update otp set otp=? where uid=?", genOtp, UniqueId).Exec(); err != nil {
					fmt.Println("error while update the otp")
				}
			} else {
				if err := cassession.Session.Query("insert into otp(uid,mobile,countrycode,otp)values(?,?,?,?)USING TTL 20", UniqueId, UserData.Mobile, UserData.CountryCode, genOtp).Exec(); err != nil {
					fmt.Println("error while insert otp to OTP Table")
					fmt.Println(err)
				}
			}
			SendOtpToMobile(UserData.CountryCode+UserData.Mobile, genOtp)
			json.NewEncoder(w).Encode(p)
		} else {
			p := Result{Status: false, Message: "user already SignUp"}
			json.NewEncoder(w).Encode(p)
		}
	} else {
		p := Result{Status: false, Message: "invalid mobile or email"}
		json.NewEncoder(w).Encode(p)
	}
	// time.Sleep(1 * time.Minute)
	// if err := cassession.Session.Query("delete from otp where uid=?", UniqueId).Exec(); err != nil {
	// 	fmt.Println("error while delete mobile otp")
	// 	fmt.Println(err)
	// }
}

func ValidateEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(email)
}
func ValidateMobile(email string) bool {
	number := phonenumber.Parse(email, "")
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
	abc := gomail.NewMessage()
	abc.SetHeader("From", "kvinothsha@gmail.com")
	abc.SetHeader("To", email)
	abc.SetHeader("Subject", "Test Email subject abc")
	abc.SetBody("text", num)
	a := gomail.NewDialer("smtp.gmail.com", 587, "kvinothsha@gmail.com", "Vinoth97$$@@")

	if err := a.DialAndSend(abc); err != nil {
		fmt.Println("error at email send for invalid email")
		fmt.Println(err)
		panic(err)
	}
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
	// time.Sleep(1 * time.Minute)
	// if err := cassession.Session.Query("delete from otp where uid=?", UniqueId).Exec(); err != nil {
	// 	fmt.Println("error while delete mobile otp")
	// 	fmt.Println(err)
	// }
}
