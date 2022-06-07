package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type StubStore struct {
	response   string
	isCanceled bool
}

func (s *StubStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}

func (s *StubStore) Cancel() {
	s.isCanceled = true
}

// func TestServer(t *testing.T) {
// 	data := "hello, world"
// 	svr := Server(&StubStore{data, false})

// 	request := httptest.NewRequest(http.MethodGet, "/", nil)
// 	response := httptest.NewRecorder()

// 	svr.ServeHTTP(response, request)

// 	if response.Body.String() != data {
// 		t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
// 	}
// }

// func TestServerCanceled(t *testing.T) {
// 	data := "hello, world"
// 	store := &StubStore{data, false}
// 	svr := Server(store)

// 	request := httptest.NewRequest(http.MethodGet, "/", nil)
// 	response := httptest.NewRecorder()

// 	cancellingCtx, cancel := context.WithCancel(request.Context())

// 	time.AfterFunc(5*time.Millisecond, cancel)
// 	request = request.WithContext(cancellingCtx)

// 	svr.ServeHTTP(response, request)

// 	assert.Equal(t, true, store.isCanceled)
// }

// func TestServerOk(t *testing.T) {
// 	data := "hello, world"
// 	store := &StubStore{data, false}
// 	svr := Server(store)

// 	request := httptest.NewRequest(http.MethodGet, "/", nil)
// 	response := httptest.NewRecorder()

// 	svr.ServeHTTP(response, request)

// 	assert.Equal(t, "hello, world", response.Body.String())
// 	assert.Equal(t, false, store.isCanceled)
// }

type SpyStore struct {
	response string
	t        *testing.T
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				log.Println("spy store got cancelled")
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

func TestHappy(t *testing.T) {
	data := "hello, world"
	store := &SpyStore{response: data, t: t}
	svr := Server(store)

	request := httptest.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	svr.ServeHTTP(response, request)

	if response.Body.String() != data {
		t.Errorf(`got "%s", want "%s"`, response.Body.String(), data)
	}
}

type SpyResponseWriter struct {
	written bool
}

func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}

func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}

func TestCancel(t *testing.T) {
	data := "hello, world"
	store := &SpyStore{response: data, t: t}
	svr := Server(store)

	request := httptest.NewRequest(http.MethodGet, "/", nil)

	cancellingCtx, cancel := context.WithCancel(request.Context())
	time.AfterFunc(5*time.Millisecond, cancel)
	request = request.WithContext(cancellingCtx)

	response := &SpyResponseWriter{}

	svr.ServeHTTP(response, request)

	if response.written {
		t.Error("a response should not have been written")
	}
}
