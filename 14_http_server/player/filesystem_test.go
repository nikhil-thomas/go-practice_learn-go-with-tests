package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("/league from reader", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
            {"Name": "Cleo", "Wins": 10},
            {"Name": "Chris", "Wins": 33}]`)

		defer cleanDatabase()
		store, err := NewFileSystemStore(database)
		assertNoError(t, err)
		got := store.GetLeague()

		want := []Player{
			{"Chris", 33},
			{"Cleo", 10},
		}

		assertLeague(t, got, want)

		got = store.GetLeague()
		assertLeague(t, got, want)

	})

	t.Run("get player score", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
            {"Name": "Cleo", "Wins": 10},
            {"Name": "Chris", "Wins": 33}]`)

		defer cleanDatabase()

		store, err := NewFileSystemStore(database)
		assertNoError(t, err)
		got, _ := store.GetPlayerScore("Chris")

		want := 33

		assertScoreEquals(t, got, want)
	})

	t.Run("store wins for existing palyers", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
            {"Name": "Cleo", "Wins": 10},
            {"Name": "Chris", "Wins": 33}]`)

		defer cleanDatabase()

		store, err := NewFileSystemStore(database)
		assertNoError(t, err)

		store.RecordWin("Chris")
		got, _ := store.GetPlayerScore("Chris")
		want := 34

		assertScoreEquals(t, got, want)
	})

	t.Run("store wins for new players", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
            {"Name": "Cleo", "Wins": 10},
            {"Name": "Chris", "Wins": 33}]`)

		defer cleanDatabase()

		store, err := NewFileSystemStore(database)
		assertNoError(t, err)
		store.RecordWin("Pepper")
		got, _ := store.GetPlayerScore("Pepper")
		want := 1
		assertScoreEquals(t, got, want)
	})

	t.Run("works with an enpty file", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, "")
		defer cleanDatabase()

		_, err := NewFileSystemStore(database)
		assertNoError(t, err)
	})

}

func assertScoreEquals(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("didnt expect an error but got one, %v", err)
	}
}

func createTempFile(t *testing.T, initialData string) (*os.File, func()) {
	t.Helper()

	tmpfile, err := ioutil.TempFile("", "db")

	if err != nil {
		t.Fatalf("coundnot create tempfile %v", err)
	}

	tmpfile.Write([]byte(initialData))

	removefile := func() {
		os.Remove(tmpfile.Name())
	}

	return tmpfile, removefile

}
