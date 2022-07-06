package countrystate

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sha/cassession"
)

type Country struct {
	CountryName string   `json:"countryname"`
	Statesname  []string `json:"statesname"`
}

type States struct {
	Statesname []string `json:"states"`
}

func CountryToFront(w http.ResponseWriter, r *http.Request) {
	var allCountry []Country
	iter := cassession.Session.Query("select * from countrystate").Iter()
	m := map[string]interface{}{}
	for iter.MapScan(m) {
		fmt.Println("called")
		allCountry = append(allCountry, Country{
			CountryName: m["countryname"].(string),
		})
		m = map[string]interface{}{}
	}
	Conv, _ := json.MarshalIndent(allCountry, "", "  ")
	fmt.Fprintf(w, "%s", string(Conv))

}
func StatesToFront(w http.ResponseWriter, r *http.Request) {
	var allStates []States
	iter := cassession.Session.Query("select * from countrystate").Iter()
	m := map[string]interface{}{}
	for iter.MapScan(m) {
		fmt.Println("called")
		allStates = append(allStates, States{
			Statesname: m["states"].([]string),
		})
		m = map[string]interface{}{}
	}
	Conv, _ := json.MarshalIndent(allStates, "", " ")
	fmt.Fprintf(w, "%s", string(Conv))
}

func Selectcountrystate(w http.ResponseWriter, r *http.Request) {
	var allStates []States
	req,_:=ioutil.ReadAll(r.Body)
	var country Country
	json.Unmarshal(req,&country)
	iter := cassession.Session.Query("select * from countrystate where countryname=? allow filtering",country.CountryName).Iter()
	m := map[string]interface{}{}
	for iter.MapScan(m) {
		fmt.Println("called")
		allStates = append(allStates, States{
			Statesname: m["states"].([]string),
		})
		m = map[string]interface{}{}
	}
	Conv, _ := json.MarshalIndent(allStates, "", " ")
	fmt.Fprintf(w, "%s", string(Conv))
}
