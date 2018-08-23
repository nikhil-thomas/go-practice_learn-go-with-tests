package main

import (
	"log"
	"net/http"
	"os"
)

const dbFileName = "fame.db.json"

type tape struct {
	file *os.File
}

func (t *tape) Write(p []byte) (n int, err error) {
	t.file.Truncate(0)
	t.file.Seek(0, 0)
	return t.file.Write(p)
}

func main() {

	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("problem openning %s, %v", dbFileName, err)
	}

	store, err := NewFileSystemStore(db)
	if err != nil {
		log.Fatalf("problem creating file system player store, %v ", err)
	}

	//store := &InMemoryPlayerStore{}
	handler := NewPlayerServer(store)
	if err := http.ListenAndServe(":5000", handler); err != nil {
		log.Fatalf("could not listen on port %v\n", 5000)
	}
}
