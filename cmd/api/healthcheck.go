//Filename: cmd/api/healthcheck.go

package main

import (
	"fmt"
	"net/http"
)

// Declare a handler which writes a plain-text response with information about the
// application status, operating environment and version.
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "status: available")
	// fmt.Fprintf(w, "environment: %s\n", app.config.env)
	// fmt.Fprintf(w, "version: %s\n", version)
	js := `{"status": "available", "environment": %q, "version": %q}`
	js = fmt.Sprintf(js, app.config.env, version)
	//specifying serve of responses using json
	w.Header().Set("Content-Type", "application/json")
	//write the json as a http response body
	w.Write([]byte(js))
}
