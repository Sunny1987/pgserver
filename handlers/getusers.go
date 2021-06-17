package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"pgserver/middleware"
)

//NewLogger is the logger for the GET function
type NewLogger struct {
	l *log.Logger
}

//GetNewLogger returns the NewLogger handler
func GetNewLogger(l *log.Logger) *NewLogger {
	return &NewLogger{l: l}
}

func (n *NewLogger) GetUsers(rw http.ResponseWriter, r *http.Request) {
	n.l.Println("****Fetching the users********")
	users, err := middleware.GetUsers()
	if err != nil {
		n.l.Println(err)
	}
	json.NewEncoder(rw).Encode(users)
}
