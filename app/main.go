package main

import (
	"log"
	"net/http"
)

func main() {
	playerStore := PlayerStoreImpl{}
	playerServer := &PlayServer{playerStore}
	handler := http.HandlerFunc(playerServer.ServeHTTP)
	log.Fatal(http.ListenAndServe(":5000", handler))
}
