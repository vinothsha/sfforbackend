package trending

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func CreateEvent(w http.ResponseWriter, r *http.Request) {
	var newEvent Proschema
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(reqBody, &newEvent)
	w.WriteHeader(http.StatusCreated)
	log.Printf(newEvent.Userid, newEvent.Views, newEvent.UpDate, newEvent.Days)
	json.NewEncoder(w).Encode(newEvent)

	//Push data in to producer
	mm, _ := json.Marshal(newEvent)
	mainPro(mm)
}
func Getevent(w http.ResponseWriter, r *http.Request) {
	go MainCon()
}
