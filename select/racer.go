package main

import (
	"fmt"
	"net/http"
	"time"
)

var TimeOutError = fmt.Errorf("Url fetching took longer than 10 sec")
var tenSecondTimeOut = time.Second*10

func Racer(a, b string) (string, error) {
	return configurableRacer(a, b, tenSecondTimeOut)
}

func configurableRacer(a, b string, timeout time.Duration) (string, error) {
	select {
	case <- ping(a):
		return a, nil
	case <- ping(b):
		return b, nil
	case <- time.After(timeout):
	   return "", TimeOutError
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}