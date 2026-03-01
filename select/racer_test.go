package main

import "testing"

func TestRacer(t *testing.T){
	slowUrl := "https://jira.com"
	fastUrl := "https://netflix.com"
	want := fastUrl 
	got := Racer(slowUrl, fastUrl)
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}