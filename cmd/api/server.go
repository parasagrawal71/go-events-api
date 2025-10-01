package main

import (
	"fmt"
	"go-events-api/cmd/api/config"
	"go-events-api/cmd/api/repository"
	"log"
	"net/http"
	"time"

	"github.com/joho/godotenv"
)

func (app *application) serve() error {
	err := godotenv.Load("../../.env") // loads values into os.Environ
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	// Connect to database
	config.ConnectDatabase()

	// Initialize all repos once
	repository.Init()

	log.Printf("Starting server on port %d\n", app.port)
	return server.ListenAndServe()
}
