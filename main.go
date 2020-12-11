package main

import (
	"github.com/aemengo/investorbook-api/controller"
	"github.com/aemengo/investorbook-api/db"
	"github.com/aemengo/investorbook-api/server"
	"log"
	"net/http"
	"os"
)

func main() {
	databaseURI := os.Getenv("DATABASE_URI")
	if databaseURI == "" {
		log.Fatalln("DATABASE_URI environment variable must be specified")
		os.Exit(1)
	}

	db, err := db.New(databaseURI)
	expectNoError(err)

	con := controller.New(db)
	s := server.New(con)

	log.Println("Listening on :8000...")
	err = http.ListenAndServe(":8000", s.Router())
	expectNoError(err)
}

func expectNoError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}