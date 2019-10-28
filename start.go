package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/thelioncape/GoNetworkMonitor/config"
	"github.com/thelioncape/GoNetworkMonitor/modules/httphandlers"
)

func main() {
	fmt.Println("Welcome to the Golang Network Monitor")
	storageType, err := config.GetStorageType()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(storageType)

	port := ":80"

	http.HandleFunc("/", httphandlers.Login)
	http.HandleFunc("/login.html", httphandlers.Loginhtml)
	http.HandleFunc("/logo.png", httphandlers.Logopng)
	http.HandleFunc("/auth", httphandlers.Auth)
	err = http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatalln(err)
	}
}
