package storage

type InMemoryPlayerStore struct {
	store map[string]int
}

// CONSTRUCTOR
func NewInMemoryPlayerStore() *InMemoryPlayerStore {
	return &InMemoryPlayerStore{
		store: map[string]int{},
	}
}

func (i *InMemoryPlayerStore) RecordWin(name string) error {
	i.store[name]++
	return nil
}

func (i *InMemoryPlayerStore) GetPlayerScore(name string) int {
	return i.store[name]
}
