package countrystate

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sha/cassession"
)

type CountryState struct {
	CountryName string   `json:"countryname"`
	States      []string `json:"states"`
}

func CountryStateToFront(w http.ResponseWriter, r *http.Request) {
	var allCountry []CountryState
	iter := cassession.Session.Query("select * from countrystate").Iter()
	m := map[string]interface{}{}
	for iter.MapScan(m) {
		fmt.Println("called")
		allCountry = append(allCountry, CountryState{
			CountryName: m["countryname"].(string),
			States:      m["states"].([]string),
		})
		m = map[string]interface{}{}
	}
	Conv, _ := json.MarshalIndent(allCountry, "", "  ")
	fmt.Fprintf(w, "%s", string(Conv))
}
