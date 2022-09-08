package main

import (
	"fmt"
	"net/http"
	"time"
)

func Racer(a, b string) (winner string, err error) {

	// ch := make(chan bool)
	// close(ch)
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(10 * time.Second):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}

	// startA := time.Now()
	// http.Get(a)
	// aDuration := time.Since(startA)

	// startB := time.Now()
	// http.Get(b)
	// bDuration := time.Since(startB)

	// if aDuration < bDuration {
	// 	return a
	// }
	// return b
}

func ping(url string) chan bool {
	ch := make(chan bool)
	go func() {
		http.Get(url)
		// ch<-
		ch <- true
	}()
	return ch
}
