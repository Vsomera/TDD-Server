package storage

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type PostgresPlayerStore struct {
	store *sql.DB
}

func NewPostgresPlayerStore() *PostgresPlayerStore {

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	var (
		host     = os.Getenv("PSQL_HOST")
		user     = os.Getenv("PSQL_USER")
		password = os.Getenv("PSQL_PASSWORD")
		dbName   = os.Getenv("PSQL_DBNAME")
		port     = os.Getenv("PORT")
	)

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName,
	)

	// establish connection with database
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return &PostgresPlayerStore{store: db}
}

// find a player by name and return the player's score
func (s *PostgresPlayerStore) GetPlayerScore(name string) int {
	var score int
	db := s.store

	query := `SELECT player_score FROM players WHERE player_name = $1`
	db.QueryRow(query, name).Scan(&score)

	return score
}

// increment player score by 1 in database
func (s *PostgresPlayerStore) RecordWin(name string) error {
	exists := s.PlayerExists(name)
	var err error

	switch exists {
	case false:
		// if player does not exists, create a new player with a score of 1
		query := `
				INSERT INTO players (player_name, player_score)
					VALUES ($1, 1)
			`
		_, err = s.store.Exec(query, name)
	case true:
		// if player exists increment their score by 1
		query := `
			UPDATE players 
				SET player_score = player_score + 1 
				WHERE player_name = $1
			`
		_, err = s.store.Exec(query, name)
	}

	if err != nil {
		return err
	}

	return nil

}

// check if a player exists by name
func (s *PostgresPlayerStore) PlayerExists(name string) bool {
	var exists bool
	query := `SELECT EXISTS (SELECT 1 FROM players WHERE player_name = $1)`

	err := s.store.QueryRow(query, name).Scan(&exists)
	if err != nil {
		panic(err)
	}
	return exists
}
