package main

// Contains functions called from router
import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

var sequenceId = 0

func (e *Error) Error() string {
	return e.Error()
}

func homePage(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint Hit: homePage")
	fmt.Fprintf(w, "Welcome to Bookinges User homepage!")
}

func returnAllBookings(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint Hit: returnAllBookings")
	vars := mux.Vars(r)
	userId, _ := vars["userId"]

	for _, booking := range Bookings {
		if booking.Name == userId {
			json.NewEncoder(w).Encode(booking)
		}
	}
}

func returnSingleBooking(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint Hit: returnSingleBooking")
	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["id"])
	found := false

	for _, booking := range Bookings {
		if booking.Id == key {
			found = true
			json.NewEncoder(w).Encode(booking)
		}
	}
	if !found {
		errorResponse(w, "Booking not found")
	}
}

func createNewBooking(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint Hit: createNewBooking")
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&Error{Status: 400, Message: err.Error()})
	}

	var booking Booking
	json.Unmarshal(reqBody, &booking)
	error := validateBooking(booking)
	if error != "" {
		errorResponse(w, error)
	} else {
		sequenceId++
		booking.Id = sequenceId
		json.NewEncoder(w).Encode(booking)
		Bookings = append(Bookings, booking)
	}
}

func deleteBooking(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint Hit: deleteBooking")
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	found := false

	for index, booking := range Bookings {
		if booking.Id == id {
			found = true
			Bookings = append(Bookings[:index], Bookings[index+1:]...)
		}
	}
	if !found {
		errorResponse(w, "Booking not found")
	}

}

func errorResponse(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(&Error{Status: 500, Message: message})
}

func validateBooking(booking Booking) string {
	_, errDate := time.Parse(timeLayout, booking.Date)
	if errDate != nil {
		return "Start date is mandatory Date format should be 'YYYY-MM-DD'"
	}
	if booking.Name == "" {
		return fmt.Sprint("Name should not be empty")
	}
	if booking.ClassId == 0 {
		return "Class Id is mandatory"
	}
	for _, b := range Bookings {
		if booking.Name == b.Name && booking.ClassId == b.ClassId {
			return "Booking already exists"
		}
	}
	return ""
}
