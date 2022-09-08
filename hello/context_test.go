package main

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

// const (
// 	succed = "\u2713"
// 	failed = "\u2717"
// )

type SpyStore struct {
	response string
	// cancelled bool
	t *testing.T
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				s.t.Log("spy store was cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
}

// func (s *SpyStore) Cancel() {
// 	s.cancelled = true
// }

type SpyResponseWriter struct {
	written bool
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("Not implemented")
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}

func TestStore(t *testing.T) {
	data := "hello, world"
	t.Run("returns data from store", func(t *testing.T) {

		t.Logf("\tTest 0:\t%s", t.Name())

		store := &SpyStore{data, t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		svr.ServeHTTP(response, request)

		if response.Body.String() == data {
			t.Logf("\t%s\tGot %s", succed, data)
		} else {
			t.Errorf("\t%s\tGot %s, want %s", failed, response.Body.String(), data)
		}

		// store.assertWasNotCanceled()
	})

	t.Run("tells store to cancel work if request is cancelled", func(t *testing.T) {
		t.Logf("\tTest1:\t%s", t.Name())

		store := &SpyStore{data, t}
		svr := Server(store)

		request := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(request.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		request = request.WithContext(cancellingCtx)

		response := &SpyResponseWriter{}

		svr.ServeHTTP(response, request)

		if !response.written {
			t.Logf("\t%s\tA response was written", succed)
		} else {
			t.Errorf("\t%s\tA response should not have been written", failed)
		}
		// store.assertWasCanceled()
	})
}

// func (s *SpyStore) assertWasCanceled() {
// 	s.t.Helper()
// 	if s.cancelled {
// 		s.t.Logf("\t%s\tThe store was cancelled", succed)
// 	} else {
// 		s.t.Errorf("\t%s\tStore was not told to cancel", failed)
// 	}
// }

// func (s *SpyStore) assertWasNotCanceled() {
// 	s.t.Helper()
// 	if !s.cancelled {
// 		s.t.Logf("\t%s\tStore was not told to cancel", succed)
// 	} else {
// 		s.t.Errorf("\t%s\tThe store was cancelled", failed)
// 	}
// }
