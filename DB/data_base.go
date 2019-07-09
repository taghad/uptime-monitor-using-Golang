package DB

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func createUserTable(db *sql.DB) {
	st, err0 := db.Prepare("CREATE TABLE IF NOT EXISTS users(id INTEGER,userName varchar(255),urlNum int)")
	if err0 != nil {
		panic(err0.Error())
	}
	_, err1 := st.Exec()
	if err1 != nil {
		panic(err1.Error())
	}

}

func createUrlTable(db *sql.DB) {

	st, err0 := db.Prepare("CREATE TABLE IF NOT EXISTS urls(id INTEGER,url varchar(255),userName varchar(255),HealthCheck int,respOkTime int, respWarTime int,respCritTime int)")
	if err0 != nil {
		panic(err0.Error())
	}
	_, err1 := st.Exec()
	if err1 != nil {
		panic(err1.Error())
	}

}

func createRequestsTable(db *sql.DB) {

	st, err0 := db.Prepare("CREATE TABLE IF NOT EXISTS reqs(timestamp date, url_id int,state int,status_code int,respTime int)")
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
