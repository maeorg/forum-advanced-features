package main

import (
	"errors"
	"fmt"
	"forum/controller"
	"forum/repository"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	// initialize database
	var err error
	repository.Database, err = repository.InitializeDatabase()
	if err != nil {
		fmt.Println("Error initializing database. " + err.Error())
		return
	}
	defer repository.Database.Close()

	// create tables to database
	err = repository.CreateTables(repository.Database)
	if err != nil {
		fmt.Println("Error creating tables to database. " + err.Error())
		return
	}

	controller.Handler()

	// run the server
	fmt.Println("Running server at http://localhost:8080")
	fmt.Println("...to shut down server, press Ctrl+C")
	err = http.ListenAndServe(":8080", nil)

	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("Server closed\n")
	} else if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		os.Exit(1)
	}
}
