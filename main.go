package main

import (
	"github.com/aemengo/investorbook-api/controller"
	"github.com/aemengo/investorbook-api/db"
	"github.com/aemengo/investorbook-api/server"
	"log"
	"net/http"
)

func main() {
	db, err := db.New("postgres://postgres:postgres@localhost:5432/postgres")
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