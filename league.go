package main

import (
	"encoding/json"
	"fmt"
	"io"
)

type League []Player

// converts league to json with a pointer to the reader
func NewLeague(rdr io.Reader) ([]Player, error) {
	var league []Player
	err := json.NewDecoder(rdr).Decode(&league)
	if err != nil {
		err = fmt.Errorf("parsing error: %s", err)
		return nil, err
	}
	return league, nil
}

func (l League) Find(name string) *Player {
	for i, p := range l {
		if p.Name == name {
			return &l[i]
		}
	}
	return nil
}
