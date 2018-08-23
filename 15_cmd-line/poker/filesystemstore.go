package poker

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"
)

// FileSystemStore defines a store based on a file
type FileSystemStore struct {
	database *json.Encoder
	league   League
}

// League is a list of Player(s)
type League []Player

// Find finds a player with matching name
func (l League) Find(name string) *Player {
	for i, player := range l {
		if player.Name == name {
			return &l[i]
		}
	}
	return nil
}

// NewFileSystemStore creates a new FileSystemStore
func NewFileSystemStore(file *os.File) (*FileSystemStore, error) {
	err := initializePlayerDBFile(file)
	if err != nil {
		return nil, fmt.Errorf("initializing db from file %s, %v", file.Name(), err)
	}

	league, err := NewLeague(file)
	if err != nil {
		return nil, fmt.Errorf("problem loading player store from file %s, %v", file.Name(), err)
	}
	return &FileSystemStore{
		database: json.NewEncoder(&tape{file}),
		league:   league,
	}, nil
}

func initializePlayerDBFile(file *os.File) error {
	file.Seek(0, 0)

	info, err := file.Stat()

	if err != nil {
		return fmt.Errorf("problem getting file info from file %s, %v", file.Name(), err)
	}

	if info.Size() == 0 {
		file.Write([]byte("[]"))
		file.Seek(0, 0)
	}
	return nil
}

// GetLeague returns league
func (f *FileSystemStore) GetLeague() League {
	//league, _ := NewLeague(f.database)
	sort.Slice(f.league, func(i, j int) bool {
		return f.league[i].Wins < f.league[j].Wins
	})
	return f.league
}

// GetPlayerScore returns palyer score
func (f *FileSystemStore) GetPlayerScore(name string) (int, error) {
	//league := f.GetLeague()
	player := f.league.Find(name)
	if player != nil {
		return player.Wins, nil
	}

	return 0, nil
}

// RecordWin records wins
func (f *FileSystemStore) RecordWin(name string) {
	//league := f.GetLeague()
	player := f.league.Find(name)
	if player != nil {
		player.Wins++
	} else {
		f.league = append(f.league, Player{name, 1})
	}
	//f.database.Seek(0, 0)
	//json.NewEncoder(f.database).Encode(f.league)
	f.database.Encode(f.league)
}
