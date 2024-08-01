package main

import (
	"encoding/json"
	"io"
)

type FileSystemPlayerStore struct {
	db io.ReadWriteSeeker
}

func (f *FileSystemPlayerStore) GetLeague() League {
	f.db.Seek(0, io.SeekStart)

	league, _ := NewLeague(f.db)
	return league
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {

	player := f.GetLeague().Find(name)
	if player != nil {
		return player.Wins
	}
	return 0
}

func (f *FileSystemPlayerStore) RecordWin(name string) {

	league := f.GetLeague()
	player := league.Find(name)

	if player != nil {
		player.Wins++
	} else {
		league = append(league, Player{name, 1})
	}

	f.db.Seek(0, io.SeekStart)
	json.NewEncoder(f.db).Encode(league)
}
