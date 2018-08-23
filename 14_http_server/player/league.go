package main

import (
	"encoding/json"
	"io"
)

// NewLeague creates a new league
func NewLeague(rdr io.ReadSeeker) (League, error) {
	rdr.Seek(0, 0)
	var league League
	err := json.NewDecoder(rdr).Decode(&league)
	return league, err
}
