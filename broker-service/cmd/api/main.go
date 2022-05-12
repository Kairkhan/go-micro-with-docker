package main

import (
	"fmt"
	"log"
	"net/http"
)

const appPort = "80"

type Application struct{}

func main() {
	app := Application{}

	log.Printf("Starting server at port :%s", appPort)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", appPort),
		Handler: app.routes(),
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Panic(err)
	}

}
