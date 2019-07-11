package server

import (
	"../DB"
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var myRouter *mux.Router
var db *sql.DB
var onlineUser string

func signUp(db *sql.DB, userName string, password string) (ok bool) {

	if (strings.Compare(userName, "") & strings.Compare(password, "")) != 0 {

		pass, _ := DB.SelectUser(db, userName)
		if strings.Compare(pass, "") == 0 {

			DB.InsertNewUser(db, userName, password)
			log.Println("user " + userName + " created")
			return true
		} else {
			log.Println("this user already exist")
		}

	}
	return false

}

func signIn(db *sql.DB, userName string, passIn string) (ok bool) {
	if (strings.Compare(userName, "") & strings.Compare(passIn, "")) != 0 {

		password, _ := DB.SelectUser(db, userName)

		if (strings.Compare(password, passIn)) == 0 {
			log.Println(userName + " Logged-in")
			return true
		} else {
			log.Println("your inf isn't true")
		}
	}
	return false

}

func addUrl(db *sql.DB, url string, userName string, healthCheck int, respOkTime int, respWarTime int, respCritTime int) {
	DB.InsertNewURL(db, url, userName, healthCheck, respOkTime, respWarTime, respCritTime)

}
func getUrlHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "server/getUrls.html")
	case "POST":
		url := r.FormValue("url")
		healthCheck, _ := strconv.Atoi(r.FormValue("healthCheck"))
		respOkTime, _ := strconv.Atoi(r.FormValue("respOkTime"))
		respWarTime, _ := strconv.Atoi(r.FormValue("respWarTime"))
		respCritTime, _ := strconv.Atoi(r.FormValue("respCritTime"))
		urlId, _ := DB.SelectUrl(db, url, onlineUser)
		if urlId == 0 {
			addUrl(db, url, onlineUser, healthCheck, respOkTime, respWarTime, respCritTime)
		} else {

		}

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}

}

func logHandler(w http.ResponseWriter, r *http.Request) {

	db = DB.ConnectDB("manager", "123456")

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "server/Sign.html")

	case "POST":
		userName := r.FormValue("newUser")
		password := r.FormValue("newPsw")
		signup := signUp(db, userName, password)
		if signup {
			onlineUser = userName
			http.Redirect(w, r, "http://localhost:10000/getUrl", 301)
			myRouter.HandleFunc("/getUrl", getUrlHandler)
		}

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}

}

func handleRequset() {
	myRouter = mux.NewRouter().StrictSlash(true)
	DB.ConnectDB("manager", "123456")

	myRouter.HandleFunc("/", logHandler)

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}
func Serve() {
	handleRequset()

}
