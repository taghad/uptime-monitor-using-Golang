package server

import (
	"fmt"
	"log"
	"net/http"
)

func signHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "Sign.html")
	fmt.Printf("%s\n", r.FormValue("pswIn"))

}
func Serve() {

	http.HandleFunc("/", signHandler)
	//
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe", err)
	} else {
		log.Println("listen 8080")
	}
}
