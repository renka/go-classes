package main

// Contains routing of application

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func handleRequests() {
	log.Printf("Started %s", port)
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/bookings/{userId}", returnAllBookings) // all bookings for user
	myRouter.HandleFunc("/booking", createNewBooking).Methods("POST")
	myRouter.HandleFunc("/bookings/{id}", deleteBooking).Methods("DELETE")
	myRouter.HandleFunc("/booking/{id}", returnSingleBooking)
	myRouter.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(version))
	})
	log.Fatal(http.ListenAndServe(port, myRouter))
}
