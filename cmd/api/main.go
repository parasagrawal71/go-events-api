package main

import (
	"go-events-api/internal/env"
	"log"

	/**
	By aliasing the import as _, you tell Go to import the package without
	directly using any of its exported functions, types, or variables in your
	code. For example, github.com/joho/godotenv/autoload loads environment
	variables from a .env file automatically when the program starts,
	 without you needing to explicitly call any function.
	*/
	_ "github.com/joho/godotenv/autoload"
)

type application struct {
	port int
}

func main() {
	app := &application{
		port: env.GetEnvInt("PORT", 8080),
	}

	if err := app.serve(); err != nil {
		log.Fatal(err)
	}
}
