package DB

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB(user string, password string) *sql.DB {
	db, err := sql.Open("mysql", user+":"+password+"@tcp(localhost)/monitorDB")
	if err != nil {
		panic(err.Error())
	}

	i, err1 := db.Prepare("CREATE TABLE IF NOT EXISTS urls(id INTEGER,url varchar(255),userName varchar(255),HealthCheck int,respOkTime int, respWarTime int,respCritTime int)")
	if err1 != nil {
		panic(err1.Error())
	}
	i.Exec()

	defer db.Close()
	return db

}

//insert to urls table
func insertNewURL(db *sql.DB, urlId int, url string, user string, healthCheckType int, respOkTime int, respWarTime int, respCritTime int) {

	st, error := db.Prepare("insert into urls (id, url, userName, HealthCheck, respOkTime, respWarTime, respCritTime) values (?,?,?,?,?,?,?)")
	if error != nil {
		panic(error.Error())
	}

	_, err := st.Exec(urlId, url, user, healthCheckType, respOkTime, respWarTime, respCritTime)
	if err != nil {
		panic(err.Error())
	}

}
