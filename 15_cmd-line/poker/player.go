package poker

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Player struct defines Palyer type
type Player struct {
	Name string
	Wins int
}

// PlayerStore interface defines a PlayerStore
type PlayerStore interface {
	GetPlayerScore(string) (int, error)
	RecordWin(name string)
	GetLeague() League
}

// PlayerServer defines a PlayerServer
type PlayerServer struct {
	store PlayerStore
	http.Handler
}

// NewPlayerServer creates a new player server
func NewPlayerServer(store PlayerStore) *PlayerServer {

	p := new(PlayerServer)
	p.store = store

	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.playersHandler))
	p.Handler = router
	return p
}

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	leagueTable := p.store.GetLeague()
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(leagueTable)
}

func (p *PlayerServer) playersHandler(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]
	switch r.Method {
	case http.MethodGet:
		p.getScores(w, player)
	case http.MethodPost:
		p.processWins(w, player)
	}
}

func (p *PlayerServer) processWins(w http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}

func (p *PlayerServer) getScores(w http.ResponseWriter, player string) {

	//score := getPlayerStore(player)
	//w.Write([]byte(score))
	score, err := p.store.GetPlayerScore(player)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	fmt.Fprint(w, score)
	return
}
