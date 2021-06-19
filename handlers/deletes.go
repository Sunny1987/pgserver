package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"pgserver/middleware"
	"pgserver/models"
	"strconv"
)

func (n *NewLogger) DeleteAScan(rw http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	// convert the id in string to int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("Unable to convert the string into int.  %v", err)
	}

	// call the deleteUser, convert the int to int64
	deletedRows := middleware.DeleteScan(int64(id))

	// format the message string
	msg := fmt.Sprintf("User updated successfully. Total rows/record affected %v", deletedRows)

	// format the reponse message
	res := models.Response{
		ID:      int64(id),
		Message: msg,
	}

	// send the response
	json.NewEncoder(rw).Encode(res)
}
