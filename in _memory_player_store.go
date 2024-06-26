package main

type InMemoryPlayerStore struct {
	store    map[string]int
	winCalls []string
	league   []Player
}

// CONSTRUCTOR
func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{
		store: map[string]int{},
	}
}

func (i *InMemoryPlayerStore) RecordWin(name string) {
	i.store[name]++
	i.winCalls = append(i.winCalls, name)
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}

func (i *InMemoryPlayerStore) GetLeague() []Player {
	return i.league
}
