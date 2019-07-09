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

	fmt.Printf("resp time %f", elapsed.Seconds())
	DB.ConnectDB("manager", "123456")

	server.Serve()

}
