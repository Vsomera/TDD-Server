package main

import (
	"io"
)

type FileSystemPlayerStore struct {
	db io.ReadSeeker
}

func (f *FileSystemPlayerStore) GetLeague() []Player {
	f.db.Seek(0, io.SeekStart)

	league, _ := NewLeague(f.db)
	return league
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	league, _ := NewLeague(f.db)
	for _, p := range league {
		if p.Name == name {
			return p.Wins
		}
	}

	return 0
}

func (f *FileSystemPlayerStore) RecordWin(name string) {

}
