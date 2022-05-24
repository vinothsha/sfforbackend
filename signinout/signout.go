package signinout

import (
	"encoding/json"
	"net/http"
	"sha/signup"
)

func Signout(w http.ResponseWriter, r *http.Request) {
	p := signup.Result{Status: true, Message: "Signout Successfully"}
	c := http.Cookie{
		Name: "token",
		// MaxAge: -1}
	}
	http.SetCookie(w, &c)
	json.NewEncoder(w).Encode(p)
	// w.Write([]byte("Old cookie deleted. Logged out!\n"))
}
