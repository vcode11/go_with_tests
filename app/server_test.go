package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type stubPlayerStore struct {
	scores map[string]int
}

func (p stubPlayerStore) GetPlayerScore(player string) int {
	return p.scores[player]
}

func TestGetPlayers(t *testing.T) {
	stubStore := stubPlayerStore{
		scores: map[string]int{
			"pepper": 20,
			"floyd":  10,
		},
	}

	playerServer := &PlayServer{stubStore}

	t.Run("return Pepper's score", func(t *testing.T) {
		request := getScoreRequest("pepper")
		response := httptest.NewRecorder()
		playerServer.ServeHTTP(response, request)
		got := response.Body.String()
		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, got, "20")
	})

	t.Run("return Floyd's score", func(t *testing.T) {
		request := getScoreRequest("floyd")
		response := httptest.NewRecorder()
		playerServer.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusOK)
		got := response.Body.String()
		assertResponseBody(t, got, "10")
	})

	t.Run("Handle missing player score", func(t *testing.T) {
		request := getScoreRequest("random_player_who_doesn't exist")
		response := httptest.NewRecorder()
		playerServer.ServeHTTP(response, request)
		got := response.Code
		want := http.StatusNotFound
		if got != want {
			t.Errorf("want %d, got %d", want, got)
		}
	})
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func getScoreRequest(player string) *http.Request {
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", player), nil)
	return request
}

func assertStatus(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}
