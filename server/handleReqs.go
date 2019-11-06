package server

import (
	"../DB"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"
)
//find bug!!!
//we need to sleep for any url 5seconds but we sleep one by one
func HandleReqs(db *sql.DB, url_id int, url string, respOkTime int, respWarTime int, respCritTime int) {
	var state int
	start := time.Now()
	resp, err0 := http.Get(url)
	if err0 != nil {
		log.Print(url + " status changed to critical")
	}
	elapsed := time.Since(start)
	if resp == nil {
		DB.DeleteReqs(db, url_id)
		DB.InsertNewReq(db, url_id, 3, 404, int(elapsed.Seconds()))
		return

	} else if 200 > resp.StatusCode || resp.StatusCode > 299 {
		state = 3
		log.Println(url + " status changed to critical")
	} else if elapsed.Seconds()-float64(respOkTime) < 0 {
		state = 1
	} else if elapsed.Seconds()-float64(respWarTime) < 0 {
		state = 2
		log.Println(url + " status changed to warning")
	} else {
		state = 3
		log.Println(url + " status changed to critical")
	}
	DB.DeleteReqs(db, url_id)
	DB.InsertNewReq(db, url_id, state, resp.StatusCode, int(elapsed.Seconds()))
}

//just here use database dirty
//we hae bu here
func makeReqs(db *sql.DB) {
	//var t int
	var id, healthCheck, respOkTime, respWarTime, respCritTime int
	var url string
	for {
		//delete t as soon as posible
		/*if t == 20000 {
			t = 0
		}
		t += 5*/
		results, err0 := db.Query("SELECT id, url, HealthCheck, respOkTime, respWarTime, respCritTime FROM urls ")
		if err0 != nil {
			panic(err0.Error())
		}
		for results.Next() {
			results.Scan(&id, &url, &healthCheck, &respOkTime, &respWarTime, &respCritTime)
			//commit this part because of t variable
			/*if t%healthCheck == 0 {
				go HandleReqs(db, id, url, respOkTime, respWarTime, respCritTime)

			}*/
			time.Sleep(5 * time.Second)
		}

	}
}
func printReq(w http.ResponseWriter, state int, status_code int, respTime int, timest string) {
	if state == 1 {
		fmt.Fprintf(w, "state : OK    ")
	} else if state == 2 {
		fmt.Fprintf(w, "state : warning    ")
	} else {
		fmt.Fprintf(w, "state :critical    ")
	}
	fmt.Fprintf(w, "status code : %d    ", status_code)
	fmt.Fprintf(w, "response time(s) : %d    ", respTime)
	fmt.Fprintf(w, "timestamp : %s    \n", timest)
}

func showReqs(db *sql.DB, url string, user string, w http.ResponseWriter) {

	var id, healthCheck, respOkTime, respWarTime, respCritTime int
	results, err0 := db.Query("SELECT id, HealthCheck, respOkTime, respWarTime, respCritTime FROM urls where url = ? and userName = ?", url, user)
	if err0 != nil {
		panic(err0.Error())
	}
	hasNext := results.Next()
	if hasNext != false {
		results.Scan(&id, &healthCheck, &respOkTime, &respWarTime, &respCritTime)
	}
	results1, err1 := db.Query("SELECT state, status_code, respTime, timestamp FROM reqs where url_id = ?", id)
	if err1 != nil {
		panic(err1.Error())
	}
	//print reqs
	var state, status_code, respTime int
	var timest string
	for results1.Next() {
		results1.Scan(&state, &status_code, &respTime, &timest)
		printReq(w, state, status_code, respTime, timest)

	}

}
