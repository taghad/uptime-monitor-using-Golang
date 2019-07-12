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

func addUrl(db *sql.DB, url string, userName string, healthCheck int, respOkTime int, respWarTime int, respCritTime int) {
	DB.InsertNewURL(db, url, userName, healthCheck, respOkTime, respWarTime, respCritTime)
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
		showUrl := r.FormValue("showUrl")
		healthCheck, _ := strconv.Atoi(r.FormValue("healthCheck"))
		respOkTime, _ := strconv.Atoi(r.FormValue("respOkTime"))
		respWarTime, _ := strconv.Atoi(r.FormValue("respWarTime"))
		respCritTime, _ := strconv.Atoi(r.FormValue("respCritTime"))

		signup := signUp(db, userName, password)
		signin := signIn(db, userIn, passIn)
		if strings.Compare(showUrl, "") != 0 {
			showReqs(db, showUrl, onlineUser, w)
		} else if signup {
			onlineUser = strings.Replace(onlineUser, onlineUser, userName, -1)
			http.ServeFile(w, r, "server/getUrls.html")
		} else if signin {
			onlineUser = strings.Replace(onlineUser, onlineUser, userIn, -1)
			http.ServeFile(w, r, "server/getUrls.html")
		} else {
			http.ServeFile(w, r, "server/Sign.html")
		}
		if strings.Compare(url, "") != 0 {
			urlId, _ := DB.SelectUrl(db, url, onlineUser)
			_, urlNum := DB.SelectUser(db, onlineUser)
			if (urlId == 0) && (urlNum < 5) {
				DB.IncrementUrlNum(db, onlineUser)
				addUrl(db, url, onlineUser, healthCheck, respOkTime, respWarTime, respCritTime)
			} else if urlNum == 5 {
				log.Println(onlineUser + " have 5 urls yet")
			}
		}

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}

}

func Serve() {

	db = DB.ConnectDB("manager", "123456")
	go makeReqs(db)
	myRouter = mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", logHandler)
	log.Fatal(http.ListenAndServe(":10000", myRouter))

}
