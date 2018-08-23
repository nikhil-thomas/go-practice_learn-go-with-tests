package main

import (
	"log"
	"net/http"
	"os"

	"github.com/nikhil-thomas/go-practice_learn-go-with-tests/15_cmd-line/poker"
)

const dbFileName = "fame.db.json"

func main() {

	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("problem openning %s, %v", dbFileName, err)
	}

	store, err := poker.NewFileSystemStore(db)
	if err != nil {
		log.Fatalf("problem creating file system player store, %v ", err)
	}

	//store := &InMemoryPlayerStore{}
	handler := poker.NewPlayerServer(store)
	if err := http.ListenAndServe(":5000", handler); err != nil {
		log.Fatalf("could not listen on port %v\n", 5000)
	}
}
