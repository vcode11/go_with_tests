package main

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
}

type PlayerStoreImpl struct{}

type PlayServer struct {
	store PlayerStore
}

func (p *PlayServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	score := p.store.GetPlayerScore(player)
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "%d", score)
}

func (p PlayerStoreImpl) GetPlayerScore(player string) int {
	if player == "pepper" {
		return 20
	}
	if player == "floyd" {
		return 10
	}
	return 0
}
