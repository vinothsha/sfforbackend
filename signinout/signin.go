package signinout

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"sha/cassession"
	"sha/signup"

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

var jwtKey = []byte("secret_key")

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
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
	if signup.ValidateEmail(user.Email) && user.Email != "" {
		var checkpass User
		res := cassession.Session.Query("select password from signup where usermail=? allow filtering", user.Email)
		res.Scan(&checkpass.Password)
		if CheckPasswordHash(user.Password, checkpass.Password) {
			expirationTime := time.Now().Add(time.Minute * 30)
			claims := &Claims{
				Email: user.Email,
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: expirationTime.Unix(),
				},
			}
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			tokenString, err := token.SignedString(jwtKey)
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
			var p []signup.Result

			p =append(p, signup.Result{Status: true, Message: "Signin Successfully", UserId: uid, Usermail: user.Email})
			json.NewEncoder(w).Encode(p)

		} else {
			w.WriteHeader(http.StatusNotAcceptable)
			var p []signup.Result

			p = append(p,signup.Result{Status: false, Message: "invalid mobile/email or password"})
			json.NewEncoder(w).Encode(p)
		}
	} else if signup.ValidateMobile(user.Mobile) {
		var checkpass User
		res := cassession.Session.Query("select password from signup where mobile=? allow filtering", user.Mobile)
		res.Scan(&checkpass.Password)
		if CheckPasswordHash(user.Password, checkpass.Password) {
			expirationTime := time.Now().Add(time.Minute * 30)
			mob := user.Mobile
			claims := &Claims{
				Email: mob,
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: expirationTime.Unix(),
				},
			}
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			tokenString, err := token.SignedString(jwtKey)
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
			var p []signup.Result
			p = append(p,signup.Result{Status: true, Message: "Signin Successfully", UserId: uid, UserMobile: user.Mobile, CountryCode: ccode})
			json.NewEncoder(w).Encode(p)

		} else {
			w.WriteHeader(http.StatusNotAcceptable)
			var p []signup.Result
			p = append(p,signup.Result{Status: false, Message: "invalid mobile/email or password"})
			json.NewEncoder(w).Encode(p)

		}
	}
}
func IsAuthorized(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		tokenStr := cookie.Value
		claims := &Claims{}
		tkn, err := jwt.ParseWithClaims(tokenStr, claims,
			func(t *jwt.Token) (interface{}, error) {
				return jwtKey, nil
			})
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if !tkn.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		handler.ServeHTTP(w, r)
		return
	}

}

func Refresh(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenStr := cookie.Value

	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Second {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	expirationTime := time.Now().Add(time.Minute * 5)

	claims.ExpiresAt = expirationTime.Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w,
		&http.Cookie{
			Name:    "refresh_token",
			Value:   tokenString,
			Expires: expirationTime,
		})

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
