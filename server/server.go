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
		http.ServeFile(w, r, "server/getUrls.html")
		url := r.FormValue("url")
		healthCheck, _ := strconv.Atoi(r.FormValue("healthCheck"))
		respOkTime, _ := strconv.Atoi(r.FormValue("respOkTime"))
		respWarTime, _ := strconv.Atoi(r.FormValue("respWarTime"))
		respCritTime, _ := strconv.Atoi(r.FormValue("respCritTime"))
		if strings.Compare(url, "") != 0 {
			fmt.Println("jeeeediiiiii")
			urlId, _ := DB.SelectUrl(db, url, onlineUser)
			_, urlNum := DB.SelectUser(db, onlineUser)
			if (urlId == 0) && (urlNum < 5) {
				DB.IncrementUrlNum(db, onlineUser)
				addUrl(db, url, onlineUser, healthCheck, respOkTime, respWarTime, respCritTime)

				http.Redirect(w, r, "http://localhost:10000", 301)
			} else if urlNum == 5 {
				fmt.Fprintf(w, "you have now 5 urls\n you can't add anymore")

			}
		}
	default:
		fmt.Fprintf(w, "Sorry, only GET method is supported.")
	}

}

func logHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		http.ServeFile(w, r, "server/Sign.html")

	case "POST":
		userName := r.FormValue("newUser")
		password := r.FormValue("newPsw")
		userIn := r.FormValue("userIn")
		passIn := r.FormValue("pswIn")
		url := r.FormValue("url")
		healthCheck, _ := strconv.Atoi(r.FormValue("healthCheck"))
		respOkTime, _ := strconv.Atoi(r.FormValue("respOkTime"))
		respWarTime, _ := strconv.Atoi(r.FormValue("respWarTime"))
		respCritTime, _ := strconv.Atoi(r.FormValue("respCritTime"))

		signup := signUp(db, userName, password)
		signin := signIn(db, userIn, passIn)
		if signup {
			onlineUser = strings.Replace(onlineUser, onlineUser, userName, -1)
			http.ServeFile(w, r, "server/getUrls.html")
		} else if signin {
			onlineUser = strings.Replace(onlineUser, onlineUser, userIn, -1)
			http.ServeFile(w, r, "server/getUrls.html")
		} else {
			http.ServeFile(w, r, "server/Sign.html")
		}
		if strings.Compare(url, "") != 0 {
			println(onlineUser)
			urlId, _ := DB.SelectUrl(db, url, onlineUser)
			_, urlNum := DB.SelectUser(db, onlineUser)
			if (urlId == 0) && (urlNum < 5) {
				DB.IncrementUrlNum(db, onlineUser)
				addUrl(db, url, onlineUser, healthCheck, respOkTime, respWarTime, respCritTime)
			} else if urlNum == 5 {
				log.Fatal(onlineUser + " have 5 urls yet")
			}
		}

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}

}

func handleRequset() {
	myRouter.HandleFunc("/", logHandler)

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func Serve() {
	myRouter = mux.NewRouter().StrictSlash(true)
	db = DB.ConnectDB("manager", "123456")
	handleRequset()
}
