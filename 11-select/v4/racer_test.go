package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("compare the speed of each server, retuning the faster one", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		got, err := Racer(slowURL, fastURL)
		want := fastURL

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}

		if err != nil {
			t.Errorf("did not expect an error but got one: %s", err.Error())
		}
	})

	t.Run("return an error if a server doesn't respond within the speciefied time", func(t *testing.T) {
		server := makeDelayedServer(30 * time.Millisecond)

		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, 10 * time.Millisecond)

		if err == nil {
			t.Errorf("expected an error but did not received one")
		}
	})
}


func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}