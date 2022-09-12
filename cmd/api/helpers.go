// Filename: cmd/api/helpers

package main

import (
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
