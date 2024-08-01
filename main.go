package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

const dbFileName = "game.db.json"

func main() {
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err.Error())
	}

	store := &FileSystemPlayerStore{db: db}
	server := NewPlayerServer(store)

	fmt.Printf("Server running on localhost:5000")

	log.Fatal(http.ListenAndServe(":5000", server))
}
