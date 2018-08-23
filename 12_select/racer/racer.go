package racer

import (
	"errors"
	"net/http"
	"time"
)

const tenS = 10 * time.Second

// Racer takes two urls and returns a url
func Racer(a, b string) (string, error) {
	return ConfigurableRacer(a, b, tenS)
}

// ConfigurableRacer compares urls and returns fasters one
func ConfigurableRacer(a, b string, timeout time.Duration) (string, error) {
	select {
	case <-Ping(a):
		return a, nil
	case <-Ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", errors.New("time exceded 10s")
	}
}

// MeasureResponseTime returns response time of a request
func MeasureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}

// Ping returns url after successful ping
func Ping(url string) <-chan bool {
	ch := make(chan bool)
	//defer close(ch)
	go func() {
		http.Get(url)
		ch <- true
	}()
	return ch
}
