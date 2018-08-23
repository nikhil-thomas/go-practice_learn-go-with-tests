package main

import (
	"fmt"
	"log"
	"os"

	"github.com/nikhil-thomas/go-practice_learn-go-with-tests/15_cmd-line/poker"
)

const dbFileName = "game.db.json"

func main() {
	fmt.Println("Let's play poker")
	fmt.Println("type {Name} wins to record a win")

	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}

	store, err := poker.NewFileSystemStore(db)
	if err != nil {
		log.Fatalf("problem creating file system player store, %v", err)
	}

	game := poker.NewPokerCLI(store, os.Stdin)
	game.PlayPoker()
}
