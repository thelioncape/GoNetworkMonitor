package httphandlers

import (
	"io/ioutil"
	"log"
	"net/http"
)

// Needs to be modified to check for a valid token - if none it should launch straight into the dashboard
func Login(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/login.html", http.StatusSeeOther)
}

func Loginhtml(w http.ResponseWriter, r *http.Request) {
	contents, err := ioutil.ReadFile("www/login/login.html")
	if err != nil {
		log.Fatal(err)
	}
	w.Write(contents)
}

func Logopng(w http.ResponseWriter, r *http.Request) {
	contents, err := ioutil.ReadFile("www/logo.png")
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Add("Content Type", "image/png")
	w.Write(contents)
}
