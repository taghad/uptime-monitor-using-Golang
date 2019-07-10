package server

import (
	"../DB"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func signUp(db *sql.DB, userName string, password string) {

	if (strings.Compare(userName, "") & strings.Compare(password, "")) != 0 {

		pass, _ := DB.SelectUser(db, userName)
		if strings.Compare(pass, "") == 0 {

			DB.InsertNewUser(db, userName, password)
			log.Println("user " + userName + " created")
		} else {
			log.Println("this user already exist")
		}

	}

}

func signIn(db *sql.DB, userName string, passIn string) {
	if (strings.Compare(userName, "") & strings.Compare(passIn, "")) != 0 {

		password, _ := DB.SelectUser(db, userName)

		if (strings.Compare(password, passIn)) == 0 {
			log.Println(userName + " Logged-in")
		} else {
			log.Println("your inf isn't true")
		}
	}

}
func homeHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, fmt.Sprintf("./%s/Sign.html", "server/"))
	db := DB.ConnectDB("manager", "123456")
	go signUp(db, r.FormValue("newUser"), r.FormValue("newPsw"))
	go signIn(db, r.FormValue("userIn"), r.FormValue("pswIn"))
}

func Serve() {

	http.HandleFunc("/", homeHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	} else {
		log.Println("listen 8080")
	}
}
