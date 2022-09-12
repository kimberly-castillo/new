//Filename: cmd/api/healthcheck.go

package main

import (
	"encoding/json"
	"net/http"
)

// Declare a handler which writes a plain-text response with information about the
// application status, operating environment and version.
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "status: available")
	// fmt.Fprintf(w, "environment: %s\n", app.config.env)
	// fmt.Fprintf(w, "version: %s\n", version)

	//write the json as a http response body
	//w.Write([]byte(js))

	// one line
	// js := `{"status": "available", "environment": %q, "version": %q}`
	// js = fmt.Sprintf(js, app.config.env, version)

	//creating a map to hold the healtcheck data
	data := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}

	//convert map to json object
	js, err := json.Marshal(data)

	if err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered a problemand could not process your request", http.StatusInternalServerError)
		return
	}
	//add new line to make viewing on terminal easier
	js = append(js, '\n')
	//specifying serve of responses using json
	w.Header().Set("Content-Type", "application/json")
	//writing out bite slice containing json response body
	w.Write(js)

}
