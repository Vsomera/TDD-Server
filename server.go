package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type Player struct {
	Name string
	Wins int
}

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	GetLeague() []Player
}

type PlayerServer struct {
	store  PlayerStore
	router *http.ServeMux
}

const jsonContentType = "application/json"

// constructor function for player server, does 1 time setup for creating the router
func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := &PlayerServer{
		store:  store,
		router: http.NewServeMux(),
	}

	p.router.Handle("/league", http.HandlerFunc(p.LeagueHandler))
	p.router.Handle("/players/", http.HandlerFunc(p.PlayerHandler))

	return p
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p.router.ServeHTTP(w, r)
}

func (p *PlayerServer) LeagueHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		// w.Header().Set("content-type", jsonContentType)
		json.NewEncoder(w).Encode(p.getLeagueTable())
		w.WriteHeader(http.StatusOK)
	}
}

// processes player requests and redirects to designated methods
func (p *PlayerServer) PlayerHandler(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}
}

// GET
func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func (p *PlayerServer) getLeagueTable() []Player {
	return p.store.GetLeague()
}

// POST
func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}

func main() {
	store := NewInMemoryPlayerStore()
	server := NewPlayerServer(store)

	log.Fatal(http.ListenAndServe(":5000", server))
}
