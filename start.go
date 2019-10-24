package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/thelioncape/GoNetworkMonitor/config"
)

func main() {
	fmt.Println("Welcome to the Golang Network Monitor")
	storageType, err := config.GetStorageType()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(storageType)

	port := ":80"

	http.HandleFunc("/", login)
	http.HandleFunc("/login.html", loginhtml)
	http.HandleFunc("/login.css", logincss)
	err = http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Listening on " + port)
}

// Needs to be modified to check for a valid token - if none it should launch straight into the dashboard
func login(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/login.html", http.StatusSeeOther)
}

func loginhtml(w http.ResponseWriter, r *http.Request) {
	contents, err := ioutil.ReadFile("www/login/login.html")
	if err != nil {
		log.Fatal(err)
	}
	w.Write(contents)
}
func logincss(w http.ResponseWriter, r *http.Request) {
	contents, err := ioutil.ReadFile("www/login/login.css")
	if err != nil {
		log.Fatal(err)
	}
	w.Write(contents)
}
