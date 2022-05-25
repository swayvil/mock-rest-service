package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func postMessageAPI(w http.ResponseWriter, r *http.Request) {
	if r.ContentLength > 10 {
		log.Printf("Message KO\n")
		w.WriteHeader(http.StatusBadRequest) // 400
	} else {
		log.Printf("Message OK\n")
		w.WriteHeader(http.StatusOK) // 200
	}
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/mock", postMessageAPI).Methods("POST")
	log.Printf("Listening to 9999\n")
	log.Fatal(http.ListenAndServe(":9999", myRouter))
}

func main() {
	handleRequests() // Start web server
}
