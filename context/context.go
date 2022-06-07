package main

import (
	"context"
	"fmt"
	"net/http"
)

type Store interface {
	Fetch(ctx context.Context) (string, error)
	// Cancel()
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, e := store.Fetch(r.Context())
		if e != nil {
			return
		} else {
			fmt.Fprint(w, data)
		}
	}
}
