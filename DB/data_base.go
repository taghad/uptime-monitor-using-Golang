package DB

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

func createUserTable(db *sql.DB) {
	st, err0 := db.Prepare("CREATE TABLE IF NOT EXISTS users(id INTEGER NOT NULL AUTO_INCREMENT,userName varchar(255),password varchar(20),urlNum int,PRIMARY KEY (id))")
	if err0 != nil {
		panic(err0.Error())
	}
	_, err1 := st.Exec()
	if err1 != nil {
		panic(err1.Error())
	}

}

func createUrlTable(db *sql.DB) {

	st, err0 := db.Prepare("CREATE TABLE IF NOT EXISTS urls(id INTEGER NOT NULL AUTO_INCREMENT,url varchar(255),userName varchar(255),HealthCheck int,respOkTime int,respWarTime int,respCritTime int,PRIMARY KEY (id))")
	if err0 != nil {
		panic(err0.Error())
	}
	_, err1 := st.Exec()
	if err1 != nil {
		panic(err1.Error())
	}

}

func createRequestsTable(db *sql.DB) {

	st, err0 := db.Prepare("CREATE TABLE IF NOT EXISTS reqs(id int NOT NULL AUTO_INCREMENT,url_id int ,state int,status_code int,respTime int, timestamp date,PRIMARY KEY (id))")
	if err0 != nil {
		panic(err0.Error())
	}
	_, err1 := st.Exec()
	if err1 != nil {
		panic(err1.Error())
	}

}

func ConnectDB(user string, password string) *sql.DB {
	db, err0 := sql.Open("mysql", user+":"+password+"@tcp(localhost)/monitorDB")
	if err0 != nil {
		panic(err0.Error())
	}

	//create tables
	createUserTable(db)
	createUrlTable(db)
	createRequestsTable(db)

	defer db.Close()
	return db
}

//insert to urls table
func insertNewURL(db *sql.DB, urlId int, url string, user string, healthCheckType int, respOkTime int, respWarTime int, respCritTime int) {

	st, err0 := db.Prepare("insert into urls (id, url, userName, HealthCheck, respOkTime, respWarTime, respCritTime) values (?,?,?,?,?,?,?)")
	if err0 != nil {
		panic(err0.Error())
	}

	_, err1 := st.Exec(urlId, url, user, healthCheckType, respOkTime, respWarTime, respCritTime)
	if err1 != nil {
		panic(err1.Error())
	}

}

func insertNewUser(db *sql.DB, userName string, password string) {

	st, err0 := db.Prepare("insert into users (userName, password, urlNum) values (?,?,?)")
	if err0 != nil {
		panic(err0.Error())
	}

	_, err1 := st.Exec(userName, password, 0)
	if err1 != nil {
		panic(err1.Error())
	}

}

func insertNewReq(db *sql.DB, url_id int, state int, status_code int, respTime int) {

	st, err0 := db.Prepare("insert into reqs (url_id, state, status_code, respTime, timestamp ) values (?,?,?,?,?)")
	if err0 != nil {
		panic(err0.Error())
	}

	_, err1 := st.Exec(url_id, state, status_code, respTime, time.Now().Format("2006-01-02 15:04:05"))
	if err1 != nil {
		panic(err1.Error())
	}

}

//select funcs :

func selectUser(db *sql.DB, username string) (password string, urlNum int) {

	results, err0 := db.Query("SELECT password, urlNum FROM users where userName = ?", username)
	if err0 != nil {
		panic(err0.Error()) // proper error handling instead of panic in your app
	}

	hasNext := results.Next()
	if hasNext == true {
		results.Scan(&password, &urlNum)
		return password, urlNum
	}
	//user not found
	{
		log.Print("user not found")
		return "", 0
	}

}

func selectUrl(db *sql.DB, url string, userName string) (id int, HealthCheck int) {

	results, err0 := db.Query("SELECT id, HealthCheck FROM urls where url = ? and userName = ?", url, userName)
	if err0 != nil {
		panic(err0.Error()) // proper error handling instead of panic in your app
	}

	hasNext := results.Next()
	if hasNext == true {
		results.Scan(&id, &HealthCheck)
		return id, HealthCheck
	}
	//url not found
	{
		log.Print("url not found")
		return 0, 0
	}

}

//not complete
func selectreq(db *sql.DB) (url_id int, state int, status_code int, respTime int, timest string) {
	//nothing
	return 0, 0, 0, 0, "0"

}
