package cassession

import (
	"fmt"

	"github.com/gocql/gocql"
)

func init() {
	var err error
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "sf"
	cluster.Consistency = gocql.LocalOne
	Session, err = cluster.CreateSession()
	if err != nil {
		fmt.Println("error while connect to cassandra")
	} else {
		fmt.Println("cassandra is connected")
	}
}
