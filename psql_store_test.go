package main

import (
	"http-server/storage"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPlayersPsql(t *testing.T) {

	store := storage.NewPostgresPlayerStore()
	server := PlayerServer{store: store}

	t.Run("get a player that does not exist", func(t *testing.T) {

		player := "William"
		request := newGetScoreRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertCorrectStatus(t, response.Code, http.StatusNotFound)
	})

}

func TestPOSTPlayersPsql(t *testing.T) {

	store := storage.NewPostgresPlayerStore()
	server := PlayerServer{store: store}

	t.Run("records a win on post", func(t *testing.T) {
		player := "Pepper"

		request := newPostWinRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertCorrectStatus(t, response.Code, http.StatusAccepted)
	})
}
