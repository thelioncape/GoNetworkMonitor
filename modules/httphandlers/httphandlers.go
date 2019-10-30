package httphandlers

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/thelioncape/GoNetworkMonitor/modules/authentication"
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
		var creds authentication.CredentialsUnhashed
		creds.Username = r.Form.Get("gnm_login")
		creds.Password = r.Form.Get("gnm_password")
		authentication.CheckCreds(creds)
	} else {
		http.Redirect(w, r, "/login.html", http.StatusSeeOther)
	}

}
