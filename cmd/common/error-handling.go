package common

import (
	"encoding/json"
	"net/http"
)

// There is not a lot of common things between two modules, but it cam be more in the future
// Also makes sense to create common structure for success responses, but I don't think it is necessary on this phase
type Error struct {
	Status  int
	Message string
}

func ErrorResponse(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(&Error{Status: 500, Message: message})
}
