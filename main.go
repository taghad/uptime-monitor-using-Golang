package main

import (
	"./DB"
	"./server"
	"fmt"
	"net/http"
	"time"
)

func main() {
	start := time.Now()

	resp, err := http.Get("http://google.com/")

	elapsed := time.Since(start)
	//dt = time.Now()
	// a = dt.Nanosecond() -a
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(resp.StatusCode)

	fmt.Printf("resp time %d", elapsed.Nanoseconds())
	DB.ConnectDB("manager", "123456")
	/*
		results, err0 := db.Query("SELECT url_id,url,HealthCheck FROM urls")
		if err0 != nil {
			panic(err0.Error())
		}
		//print reqs
		var url_id, HealthCheck int
		var url string
		for results.Next() {
			results.Scan(&url_id, &url, &HealthCheck)
			resp, err0 := http.Get(url)
			if err0 != nil {
				panic(err0.Error())
			}
			fmt.Println(resp.StatusCode)

		}*/

	server.Serve()

}
