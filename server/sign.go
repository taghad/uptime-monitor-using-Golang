package server

import (
	"../DB"
	"database/sql"
	"log"
	"strings"
)

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
