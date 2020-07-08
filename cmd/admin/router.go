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
	myRouter.PathPrefix("/admin")
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/classes", returnAllClasses)
	myRouter.HandleFunc("/class", createNewClass).Methods("POST")
	myRouter.HandleFunc("/class/{id}", deleteClass).Methods("DELETE")
	myRouter.HandleFunc("/classes/{id}", deleteClasses).Methods("DELETE")
	myRouter.HandleFunc("/class/{id}", returnSingleClass)
	myRouter.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(version))
	})
	log.Fatal(http.ListenAndServe(port, myRouter))
}
