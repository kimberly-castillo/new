// Filename: cmd/api/entries

package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

// createEntryHandler for POST /v1/entries
func (app *application) createEntryHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create new entry")
}

// showEntryHandler for POST /v1/entries/:id
func (app *application) showEntryHandler(w http.ResponseWriter, r *http.Request) {
	//using ParamsFromContext func to get request as a slice
	params := httprouter.ParamsFromContext(r.Context())

	//getting value of id parameter
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64) //look for id in param var and return associated value
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	//display id
	fmt.Fprintf(w, "show the details for id %d\n", id)
}
