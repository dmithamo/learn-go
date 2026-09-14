package main

import (
	"encoding/json"
	"net/http"
)

// healthHandler reports on the health of the api
func (a *application) healthHandler(w http.ResponseWriter, r *http.Request) {
	h, err := json.Marshal(map[string]string{"env": a.config.env,  "version": version})
	if err != nil {
		http.Error(w, "unable to marshal health to json", http.StatusInternalServerError)
	}

	h = append(h, '\n')
	w.Header().Set("Content-Type", "application/json")
	w.Write(h)
}
