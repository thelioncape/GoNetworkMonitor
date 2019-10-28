package httphandlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Needs to be modified to check for a valid token - if found it should launch straight into the dashboard
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

func Auth(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		username := r.Form.Get("gnm_login")
		password := r.Form.Get("gnm_password")
		fmt.Println(username)
		fmt.Println(password)
	} else {
		http.Redirect(w, r, "/login.html", http.StatusSeeOther)
	}

}