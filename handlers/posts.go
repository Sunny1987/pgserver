package handlers

import (
	"encoding/json"
	"net/http"
	"pgserver/middleware"
	"pgserver/models"
)

func (n *NewLogger) AddScan(rw http.ResponseWriter, r *http.Request) {
	n.l.Println("****inserting the scans********")
	var result models.Result

	err := json.NewDecoder(r.Body).Decode(&result)
	if err != nil {
		n.l.Println(err)
	}

	d, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		n.l.Println(err)
	}
	var scan models.Scan
	scan.Url = result.Request.URL
	scan.Person = result.Person
	scan.Result = string(d)

	insertId := middleware.InsertScan(scan)

	res := models.Response{
		ID:      insertId,
		Message: "Inserted Successfully",
	}

	n.l.Println(res)
}
