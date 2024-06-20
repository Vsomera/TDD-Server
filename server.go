package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

type PlayerServer struct {
	store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		p.processWin(w, r)
	case http.MethodGet:
		p.showScore(w, r)
	}
}

func (p *PlayerServer) showScore(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	p.store.RecordWin(player)

	w.WriteHeader(http.StatusAccepted)
}

type InMemoryPlayerStore struct {
	scores   map[string]int
	winCalls []string
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.scores[name]
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.winCalls = append(i.winCalls, name)
	i.scores[name] += 1
}

func main() {
	store := InMemoryPlayerStore{
		scores: map[string]int{
			"Bob":    10,
			"Robert": 20,
		},
		winCalls: nil,
	}

	server := &PlayerServer{&store}

	log.Fatal(http.ListenAndServe(":5000", server))
}
