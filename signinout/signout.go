package signinout

import (
	"encoding/json"
	"net/http"
	s "sha/commonstruct"
)

func Signout(w http.ResponseWriter, r *http.Request) {
	p := s.ErrorResult{Status: true, Message: "Signout Successfully"}
	c := http.Cookie{
		Name: "token",
		// MaxAge: -1}
	}
	http.SetCookie(w, &c)
	json.NewEncoder(w).Encode(p)
	// w.Write([]byte("Old cookie deleted. Logged out!\n"))
}
