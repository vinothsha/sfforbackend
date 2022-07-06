package signinout

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"sha/cassession"
	e "sha/commonservices/commonfunctions"

	s "sha/commonstruct"
	auth "sha/middleware"

	"github.com/dgrijalva/jwt-go"
	"github.com/gocql/gocql"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	CountryCode string `json:"countrycode"`
	Mobile      string `json:"mobile"`
}

// type SigninEmailStruct struct {
// 	SendStatus signup.Result `json:"sendstatus"`
// 	UserId     gocql.UUID    `json:"userid"`
// 	UserEmail  string        `json:"useremail"`
// }
// type SigninMobileStruct struct {
// 	SendStatus  signup.Result `json:"sendstatus"`
// 	UserId      gocql.UUID    `json:"userid"`
// 	UserMobile  string        `json:"usermobile"`
// 	CountryCode string        `json:"countrycode"`
// }

func Signin(w http.ResponseWriter, r *http.Request) {
	var user User
	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("error in read data")
	}
	json.Unmarshal(req, &user)
	if e.ValidateEmail(user.Email) && user.Email != "" {
		var checkpass User
		res := cassession.Session.Query("select password from signup where usermail=? allow filtering", user.Email)
		res.Scan(&checkpass.Password)
		if CheckPasswordHash(user.Password, checkpass.Password) {
			expirationTime := time.Now().Add(time.Minute * 30)
			claims := &auth.Claims{
				Email: user.Email,
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: expirationTime.Unix(),
				},
			}
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			tokenString, err := token.SignedString(auth.JwtKey)
			if err != nil {
				fmt.Println("error in jwt toke produce")
				return
			}
			http.SetCookie(w, &http.Cookie{
				Name:    "token",
				Value:   tokenString,
				Expires: expirationTime,
			})
			var uid gocql.UUID
			res1 := cassession.Session.Query("select uid from signup where usermail=? allow filtering", user.Email)
			res1.Scan(&uid)
			var p []s.ResultEmail

			p = append(p, s.ResultEmail{Status: true, Message: "Signin Successfully", UserId: uid.String(),Email: user.Email})
			json.NewEncoder(w).Encode(p)

		} else {
			w.WriteHeader(http.StatusNotAcceptable)
			var p []s.ErrorResult

			p = append(p, s.ErrorResult{Status: false, Message: "invalid mobile/email or password"})
			json.NewEncoder(w).Encode(p)
		}
	} else if e.ValidateMobile(user.Mobile) {
		var checkpass User
		res := cassession.Session.Query("select password from signup where mobile=? allow filtering", user.Mobile)
		res.Scan(&checkpass.Password)
		if CheckPasswordHash(user.Password, checkpass.Password) {
			expirationTime := time.Now().Add(time.Minute * 30)
			mob := user.Mobile
			claims := &auth.Claims{
				Email: mob,
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: expirationTime.Unix(),
				},
			}
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			tokenString, err := token.SignedString(auth.JwtKey)
			if err != nil {
				fmt.Println("error in jwt toke produce")
				return
			}
			http.SetCookie(w, &http.Cookie{
				Name:    "token",
				Value:   tokenString,
				Expires: expirationTime,
			})
			var uid gocql.UUID
			res1 := cassession.Session.Query("select uid from signup where mobile=? allow filtering", user.Mobile)
			res1.Scan(&uid)
			var ccode string
			res1 = cassession.Session.Query("select countrycode from signup where mobile=? allow filtering", user.Mobile)
			res1.Scan(&ccode)
			var p []s.ResultMobile
			p = append(p, s.ResultMobile{Status: true, Message: "Signin Successfully", UserId: uid.String(),Mobile:user.Mobile})
			json.NewEncoder(w).Encode(p)

		} else {
			w.WriteHeader(http.StatusNotAcceptable)
			var p []s.ErrorResult
			p = append(p, s.ErrorResult{Status: false, Message: "invalid mobile/email or password"})
			json.NewEncoder(w).Encode(p)

		}
	}
}

func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello this is a Home Page"))
}

func CheckAPI(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this api is working...This is a get method for check... Other Method Or Post Method..."))
}
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
