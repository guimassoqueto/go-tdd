package main

import (
	"fmt"
	"net/http"
	"time"
)

var customTimeout = 10 * time.Second

func Racer(urlA, urlB string) (winner string, err error) {
	return ConfigurableRacer(urlA, urlB, customTimeout)
}


func ConfigurableRacer(urlA, urlB string, timeout time.Duration) (winner string, err error) {
	select {
	case <- ping(urlA):
		return urlA, nil
	case <- ping(urlB):
		return urlB, nil
	case <- time.After(timeout):
		return "", fmt.Errorf("timeout error waiting %s and %s", urlA, urlB)
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