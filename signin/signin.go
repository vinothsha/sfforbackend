package signin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"

	"sha/cassession"

	"github.com/dgrijalva/jwt-go"
	"github.com/dongri/phonenumber"
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

func Signin(w http.ResponseWriter, r *http.Request) {
	var user User
	req, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("error in read data")
	}
	json.Unmarshal(req, &user)
	if ValidateEmail(user.Email) && user.Email != "" {
		var checkpass User
		res := cassession.Session.Query("select password from signup where usermail=? allow filtering", user.Email)
		res.Scan(&checkpass.Password)
		fmt.Println(checkpass.Password)
		if CheckPasswordHash(user.Password, checkpass.Password) {
			expirationTime := time.Now().Add(time.Second * 30)
			claims := &Claims{
				Email: user.Email,
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: expirationTime.Unix(),
				},
			}
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			tokenString, err := token.SignedString(jwtKey)
			fmt.Println(tokenString)
			if err != nil {
				fmt.Println("error in jwt toke produce")
				return
			}
			http.SetCookie(w, &http.Cookie{
				Name:    "token",
				Value:   tokenString,
				Expires: expirationTime,
			})

			fmt.Println("loged in successfully")
			fmt.Fprintln(w, "homepage")
		} else {
			fmt.Println("Invalid email/ password")
		}
	} else if ValidateMobile(user.Mobile) {
		fmt.Println("this is a mobile number")
		var checkpass User
		res := cassession.Session.Query("select password from signup where mobile=? allow filtering", user.Mobile)
		res.Scan(&checkpass.Password)
		fmt.Println(checkpass.Password)
		if CheckPasswordHash(user.Password, checkpass.Password) {
			expirationTime := time.Now().Add(time.Minute * 1)
			mob := user.CountryCode + user.Mobile
			claims := &Claims{
				Email: mob,
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: expirationTime.Unix(),
				},
			}
			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
			tokenString, err := token.SignedString(jwtKey)
			fmt.Println(tokenString)
			if err != nil {
				fmt.Println("error in jwt toke produce")
				return
			}
			http.SetCookie(w, &http.Cookie{
				Name:    "token",
				Value:   tokenString,
				Expires: expirationTime,
			})

			fmt.Println("loged in successfully")
			fmt.Fprintln(w, "homepage")
		} else {
			fmt.Println("Invalid email/ password")
		}
	}
}
func IsAuthorized(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("isauth is called")
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
func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello this is a Home Page"))
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
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	fmt.Println(err == nil)
	fmt.Println(password, hash)
	return err == nil
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
func CheckAPI(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this api is working...This is a get method for check... Other Method Or Post Method..."))
}
