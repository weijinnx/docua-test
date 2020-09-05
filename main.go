package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/weijinnx/docua-test/pkg/redis"

	"github.com/weijinnx/docua-test/pkg/errors"
	"github.com/weijinnx/docua-test/pkg/util"
)

var rdb *redis.Redis

func printMeAtHandler(jobs chan string, w http.ResponseWriter, r *http.Request) {
	// make sure that request method is GET
	if r.Method != "GET" {
		w.Header().Set("Allow", "GET")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// get time and message from query params
	t, msg, err := util.ParseQuery(r.URL.Query())
	if err != nil {
		return
	}

	printed := rdb.GetPrintedMessages()
	if !util.SliceContains(printed, msg) { // if message not in blacklist (printed already)
		// set received data to redis hash map
		rdb.AddMessage(msg, t.Format(time.RFC3339))

		// start async timer and push message after
		// it will be finished
		go util.AddJob(jobs, t, msg)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"status": true,
	})
}

func main() {
	rdb = redis.Connect()
	_, err := rdb.Ping().Result()
	if err != nil {
		log.Fatalf("%s: %v", errors.FailConnectRedisError, err)
	}

	jobs := make(chan string)
	go func() {
		for j := range jobs {
			// print message
			fmt.Println(j)

			// remove message from hash map
			rdb.RemoveMessage(j)
			// and add it to printed list
			rdb.AddToPrintedMessagesList(j)
		}
	}()

	// get all messages when server starts
	messages := rdb.GetMessages()
	for m, tstr := range messages {
		t, err := util.ParseTime(tstr)
		if err != nil {
			log.Println(err)
		}

		printed := rdb.GetPrintedMessages()
		if !util.SliceContains(printed, m) { // if message not in blacklist (printed already)
			util.AddJob(jobs, t, m)
		}
	}

	http.HandleFunc("/printMeAt", func(w http.ResponseWriter, r *http.Request) {
		printMeAtHandler(jobs, w, r)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
