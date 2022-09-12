// Filename: cmd/api/entries

package main

import (
	"fmt"
	"net/http"
)

// createEntryHandler for POST /v1/entries
func (app *application) createEntryHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create new entry")
}

// showEntryHandler for POST /v1/entries/:id
func (app *application) showEntryHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	//display id
	fmt.Fprintf(w, "show the details for id %d\n", id)
}
