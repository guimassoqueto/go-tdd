package main

import (
	"net/http"
	"time"
)

func Racer(urlA string, urlB string) (winner string) {
	aDuration := getResponseDuration(urlA)
	bDuration := getResponseDuration(urlB)

	if aDuration < bDuration {
		return urlA
	}

	return urlB
}

func getResponseDuration(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}