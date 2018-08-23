package racer

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("compares speed of urls and returns the fastest one", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		defer slowServer.Close()

		fastServer := makeDelayedServer(0 * time.Millisecond)
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, _ := Racer(slowURL, fastURL)

		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	})

	t.Run("returns an error if server doesn't respond within 10s", func(t *testing.T) {
		delayServer := makeDelayedServer(10 * time.Millisecond)

		defer delayServer.Close()

		_, err := ConfigurableRacer(delayServer.URL, delayServer.URL, 5*time.Millisecond)

		if err == nil {
			t.Error("expected error but did't get one")
		}
	})
}

func makeDelayedServer(t time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(t)
		w.WriteHeader(http.StatusOK)
	}))
}
