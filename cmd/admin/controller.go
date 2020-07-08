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
	fmt.Fprintf(w, "Welcome to Classes Admin!")
}

func returnAllClasses(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint Hit: returnAllClasses")
	json.NewEncoder(w).Encode(Classes)
}

func returnSingleClass(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint Hit: returnSingleClass")
	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["id"])
	found := false

	for _, class := range Classes {
		if class.Id == key {
			found = true
			json.NewEncoder(w).Encode(class)
		}
	}
	if !found {
		errorResponse(w, "Class not found")
	}
}

func createNewClass(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint Hit: createNewClass")
	reqBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&Error{Status: 400, Message: err.Error()})
	}

	var entireClass EntireClass
	json.Unmarshal(reqBody, &entireClass)
	error := validateClass(entireClass)
	if error != "" {
		errorResponse(w, error)
	} else {
		writeSingleClassesResponse(entireClass, w)
	}
}

func deleteClass(w http.ResponseWriter, r *http.Request) {
	log.Println("Endpoint Hit: deleteClass")
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	found := false

	for index, class := range Classes {
		if class.Id == id {
			found = true
			Classes = append(Classes[:index], Classes[index+1:]...)
		}
	}
	if !found {
		errorResponse(w, "Class not found")
	}

}

func deleteClasses(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := vars["id"]
	found := false

	for index, class := range Classes {
		if class.ClassId == id {
			found = true
			Classes = append(Classes[:index], Classes[index+1:]...)
		}
	}
	if !found {
		errorResponse(w, "Classes not found")
	}
}

func errorResponse(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(&Error{Status: 500, Message: message})
}

// Generate separate classes from class description and write response
func writeSingleClassesResponse(entireClass EntireClass, w http.ResponseWriter) {
	timeStart, _ := time.Parse(timeLayout, entireClass.StartDate)
	timeEnd, _ := time.Parse(timeLayout, entireClass.EndDate)

	days := timeEnd.Day() - timeStart.Day()
	for i := 0; i <= days; i++ {
		sequenceId++
		singleClass := SingleClass{
			Class: Class{
				Id:       sequenceId,
				Name:     entireClass.Name,
				Capacity: entireClass.Capacity,
			},
			ClassId: fmt.Sprintf("%d%d%s", timeStart.Month(), timeStart.Day(), entireClass.Name[:minNameLength]),
			Date:    timeStart.AddDate(0, 0, i).Format(timeLayout),
		}
		json.NewEncoder(w).Encode(singleClass)
		Classes = append(Classes, singleClass)
	}
}

func validateClass(class EntireClass) string {
	start, errStartDate := time.Parse(timeLayout, class.StartDate)
	end, errEndDate := time.Parse(timeLayout, class.EndDate)
	if errStartDate != nil {
		return "Start date is mandatory Date format should be 'YYYY-MM-DD'"
	}
	if errEndDate != nil {
		return "End date is mandatory Date format should be 'YYYY-MM-DD'"
	}
	if end.Day()-start.Day() < 0 {
		return "End date can't be before start date'"
	}
	if len(class.Name) < minNameLength {
		return fmt.Sprintf("Name should be longer %d", minNameLength)
	}
	if class.Capacity == 0 {
		return "Capacity is mandatory"
	}
	for _, c := range Classes {
		date, _ := time.Parse(timeLayout, c.Date)
		if c.Name == class.Name && (start.Before(date) && end.After(date)) {
			return "Class with the same name is already registered on the same dates"
		}
	}
	return ""
}
