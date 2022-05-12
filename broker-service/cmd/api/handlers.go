package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func (app *Application) Broker(w http.ResponseWriter, r *http.Request) {
	payload := jsonResponse{
		Error:   false,
		Message: "It's a BrokerApp",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	data, err := json.MarshalIndent(payload, "", "\t")
	if err != nil {
		log.Fatal(err)
	}

	_, err = w.Write(data)

}
