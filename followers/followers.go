package followers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sha/cassession"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/gocql/gocql"
)

type Follow struct {
	Userid    gocql.UUID `json:"userid"`
	Followers gocql.UUID `json:"followers"`
	Datetime  string     `json:"datetime"`
}

type Getfollow struct {
	Userid gocql.UUID `json:"userid"`
}

func Followers(w http.ResponseWriter, r *http.Request) {
	var Event Follow
	Event.Datetime = time.Now().Format("2006-01-02 15:04:05")
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Enter Data")
	}
	json.Unmarshal(reqBody, &Event)
	var Userid gocql.UUID
	var Followers map[gocql.UUID]string
	n := 0
	v := 0
	iter := cassession.Session.Query("SELECT * from followers WHERE userid = ?", Event.Userid).Iter()
	for iter.Scan(&Userid, &Followers) {
		n += 1
	}
	if n == 0 {
		if err := cassession.Session.Query("UPDATE followers SET followers[?] = ? WHERE userid = ?", Event.Followers, Event.Datetime, Event.Userid).Exec(); err != nil {
			fmt.Println("Error while updating")
			fmt.Println(err)
		}
	}
	if n != 0 {
		iter := cassession.Session.Query("SELECT * from followers WHERE userid = ?", Event.Userid).Iter()
		for iter.Scan(&Userid, &Followers) {
			for k, z := range Followers {
				if k == Event.Followers {
					v += 1
					strings.Split(z, "")
				} else {
					v += 0
				}
			}
		}
		// fmt.Println(v)
		if v == 0 {
			if err := cassession.Session.Query("UPDATE followers SET followers[?] = ? WHERE userid = ?", Event.Followers, Event.Datetime, Event.Userid).Exec(); err != nil {
				fmt.Println("Error while updating")
				fmt.Println(err)
			}
		}
		if v != 0 {
			if err := cassession.Session.Query("DELETE followers[?] FROM followers WHERE userid = ?", Event.Followers, Event.Userid).Exec(); err != nil {
				fmt.Println("Error while updating")
				fmt.Println(err)
			}
		}
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Println("Data processed sucessfully")
}

type Result struct {
	Followcount string `json:"followerscount"`
}

func GetFollowers(w http.ResponseWriter, r *http.Request) {
	var Event Getfollow
	var Res Result
	var Followers map[gocql.UUID]string
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Enter Data")
	}
	json.Unmarshal(reqBody, &Event)

	if err := cassession.Session.Query("SELECT followers FROM followers WHERE userid = ?", Event.Userid).Scan(&Followers); err != nil {
		fmt.Println("Error in getting followers")
		fmt.Println(err)
	}
	x := ""
	t := strconv.Itoa(len(Followers))
	if len(t) < 4 {
		x = t
	}
	if len(t) == 4 {
		str := strings.Split(t, "")
		x = (str[0] + "." + str[1] + "k")
	}
	if len(t) == 5 {
		str := strings.Split(t, "")
		x = (str[0] + str[1] + "." + str[2] + "k")
	}
	if len(t) == 6 {
		str := strings.Split(t, "")
		x = (str[0] + str[1] + str[2] + "." + str[3] + "k")
	}
	Res.Followcount = x
	json.NewEncoder(w).Encode(Res)
}
func GetTop(w http.ResponseWriter, r *http.Request) {
	var User_id string
	var Followers []string
	m := make(map[string]int)
	keys := make([]string, 0, len(m))
	iter := cassession.Session.Query("SELECT * FROM followers").Iter()
	for iter.Scan(&User_id, &Followers) {
		m[User_id] = len(Followers)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
	}
	for key := range m {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return m[keys[i]] > m[keys[j]]
	})
	z := 0
	for _, j := range keys {
		if z < 10 {
			z++
			fmt.Println(j)
		}

	}

}

// func main() {
// 	router := mux.NewRouter().StrictSlash(true)
// 	router.HandleFunc("/followers", Followers).Methods("POST")
// 	router.HandleFunc("/getfollowers", GetFollowers).Methods("GET")
// 	router.HandleFunc("gettopfollow", GetTop).Methods("GET")
// 	log.Fatal(http.ListenAndServe(":8080", router))
// }
