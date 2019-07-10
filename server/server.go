package server

import (
	"../DB"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, fmt.Sprintf("./%s/Sign.html", "server/"))
	//fmt.Printf("\nuser : %s\n", r.FormValue("userIn"))
	//fmt.Printf("pass : %s\n", r.FormValue("pswIn"))
	db := DB.ConnectDB("manager", "123456")
	var a string
	a = r.FormValue("user")
	if strings.Compare(a, "") != 0 {
		DB.InsertNewUser(db, r.FormValue("user"), r.FormValue("psw"))
	}

}

func Serve() {

	http.HandleFunc("/", homeHandler)
	//
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	} else {
		log.Println("listen 8080")
	}
}
