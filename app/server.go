package main

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) string
}

type PlayerStoreImpl struct{}

type PlayServer struct {
	store PlayerStore
}

func (p *PlayServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	fmt.Fprintf(w, "%s", p.store.GetPlayerScore(player))
}

func (p PlayerStoreImpl) GetPlayerScore(player string) string {
	if player == "pepper" {
		return "20"
	}
	if player == "floyd" {
		return "10"
	}
	return "unknown"
}
