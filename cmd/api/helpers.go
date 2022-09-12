// Filename: cmd/api/helpers

package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) readIDParam(r *http.Request) (int64, error) {
	//using ParamsFromContext func to get request as a slice
	params := httprouter.ParamsFromContext(r.Context())

	//getting value of id parameter
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64) //look for id in param var and return associated value
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}
	return id, nil
}

// interface meaning any valid go type
func (app *application) writeJSON(w http.ResponseWriter, status int, data interface{}, headers http.Header) error {
	//convert map to json object
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}
	//add new line to make viewing on terminal easier
	js = append(js, '\n')
	//add headers
	for key, value := range headers {
		w.Header()[key] = value //go will not crash if there are no headers
	}
	//specifying serve of responses using json
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	//writing out bite slice containing json response body
	w.Write(js)
	return nil
}
