package viewspervideo

import (
	"fmt"
	"net/http"
	"sha/cassession"

	"github.com/gocql/gocql"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

func ViewsPerVideo(w http.ResponseWriter, r *http.Request) {
	Vid := mux.Vars(r)["vid"]
	if views := cassession.Session.Query("insert into views(viewuid,videouid)values(?,?);", gocql.UUID(uuid.New()), Vid).Exec(); views != nil {
		fmt.Println(views)
	}
	fmt.Println("view is success")
}
