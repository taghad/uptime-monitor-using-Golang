package server

import (
	"../DB"
	"database/sql"
	"log"
	"net/http"
	"time"
)

func HandleReqs(db *sql.DB, url_id int, url string, respOkTime int, respWarTime int, respCritTime int) {
	var state int
	start := time.Now()
	resp, err0 := http.Get(url)
	if err0 != nil {
		panic(err0.Error())
	}
	elapsed := time.Since(start)
	if elapsed.Seconds()-float64(respOkTime) < 0 {
		state = 1
	} else if elapsed.Seconds()-float64(respWarTime) < 0 {
		state = 2
		log.Fatal(url + "status changed to warning")
	} else {
		state = 3
		log.Fatal(url + "status changed to critical")
	}
	DB.InsertNewReq(db, url_id, state, resp.StatusCode, int(elapsed.Nanoseconds()))

}

func makeReqs(db *sql.DB) {

}
