package main

import (
	"encoding/json"
	"net/http"
	"person"
)

func Create(w http.ResponseWriter, r *http.Request) {

	var req Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		HTTPError(w, err, http.StatusBadRequest)
		return
	}

	if err := person.Create(&req.Obj); err != nil {
		HTTPError(w, err, http.StatusBadRequest)
		return
	}

	output, err := json.Marshal(Response{
		Error:  err,
		Object: req.Obj,
	})

	if err != nil {
		HTTPError(w, err, http.StatusBadRequest)
		return
	}

	w.Write(output)
}
