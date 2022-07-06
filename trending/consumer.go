package trending

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"sort"
	"syscall"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	// "github.com/gocql/gocql"
)

type schema struct {
	Userid string `json:"userid"`
	Views  int        `json:"views"`
	UpDate string     `json:"update"`
	Days   int        `json:"days"`
}

func MainCon() {
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":     "localhost:9092",
		"group.id":              "myGroup1",
		"broker.address.family": "v4",
		"auto.offset.reset":     "earliest",
		"enable.auto.commit":    false})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create consumer: %s\n", err)
		os.Exit(1)
	}
	//fmt.Println("Consumer started : ", c)

	c.SubscribeTopics([]string{"top2"}, nil)

	run := true
	i := 0
	map_1 := map[string]int{}

	for run {
		select {
		case sig := <-sigchan:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			run = false
		default:
			ev := c.Poll(100)
			if ev == nil && i > 5 {
				run = false
			} else if ev == nil {
				i++
				continue
			}
			switch e := ev.(type) {
			case *kafka.Message:
				//fmt.Printf("Consumed message on %s:%s\n", e.TopicPartition, string(e.Value))

				// exit consumer if all search words untill 3 min back were processed
				timestamp := e.Timestamp
				if timestamp.Unix() > (time.Now().Unix() - 180) {
					run = false
				}
				m := schema{}
				json.Unmarshal(e.Value, &m)
				// fmt.Printf(m.Id,m.Views,m.DelDate,m.Days)
				if m.Days < 15 {
					map_1[m.Userid] = m.Views
				} else {
					delete(map_1, m.Userid)
				}
				keys := make([]string, 0, len(map_1))
				for key := range map_1 {
					keys = append(keys, key)
				}
				sort.SliceStable(keys, func(i, j int) bool {
					return map_1[keys[i]] > map_1[keys[j]]
				})
				// for _, k := range keys{
				// 	fmt.Println(k, map_1[k])
				// }

				//fmt.Println(len(map_1))

				// fmt.Println("-------------------------------------------------------------------")
			case kafka.Error:
				// The client will automatically try to recover from all errors.
				fmt.Printf("Consumer error: %v (%v)\n", err, e)
			default:
				fmt.Printf("Ignore: %v\n", e)
			}
		}
	}
	fmt.Println(map_1)
	c.Close()

}
