package main

import (
	"fmt"
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

	http.HandleFunc("/", homepage)
	err = http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Listening on " + port)
}

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}
