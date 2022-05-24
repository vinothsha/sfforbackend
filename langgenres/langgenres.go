package langgenres

import (
	"encoding/json"
	"net/http"
	"sha/cassession"
)

type Result struct {
	Genres   []string `json:"genres"`
	Language []string `json:"languages"`
}

func LanguagesGenresForDropDown(w http.ResponseWriter, r *http.Request) {
	//get Method to give language and genres to FrontEnd Dropdown

	var AllCat Result
	iter := cassession.Session.Query("SELECT genres FROM languagegenres where id=1")
	iter.Scan(&AllCat.Genres)
	iter1 := cassession.Session.Query("SELECT languages FROM languagegenres where id=1")
	iter1.Scan(&AllCat.Language)
	json.NewEncoder(w).Encode(AllCat)
	return
}
