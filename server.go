package main

import (
	"fmt"
	"http-server/storage"
	"log"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

type PlayerServer struct {
	store  PlayerStore
	router *http.ServeMux
}

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
	w.WriteHeader(http.StatusOK)
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

// POST
func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}

func main() {
	store := storage.NewInMemoryPlayerStore()
	server := NewPlayerServer(store)

	log.Fatal(http.ListenAndServe(":5000", server))
}
