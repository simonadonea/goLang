package main

import (
	"sync"
	"testing"
)

const (
	succed = "\u2713"
	failed = "\u2717"
)

func TestCounter(t *testing.T) {
	t.Run("Incrementing to 3", func(t *testing.T) {
		t.Logf("\tTest 0:\t%s", t.Name())
		counter := NewCounter()
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, counter, 3)
	})

	t.Run("Runs safetly concurrently", func(t *testing.T) {
		t.Logf("\tTest 1:\t%s", t.Name())
		wantedCount := 1000
		counter := NewCounter()

		var wg sync.WaitGroup
		wg.Add(wantedCount)

		for i := 0; i < wantedCount; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}
		wg.Wait()

		assertCounter(t, counter, wantedCount)
	})
}

func assertCounter(t testing.TB, got *Counter, want int) {
	t.Helper()
	if got.Value() == want {
		t.Logf("\t%s\tShould receive %d", succed, want)
	}
	if got.Value() != want {
		t.Errorf("\t%s\tShould receive %d, got %d", failed, got.Value(), want)
	}
}
