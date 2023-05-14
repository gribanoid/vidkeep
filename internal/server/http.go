package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/gribanoid/vidkeep/config"
)

func NewHTTPServer(c *config.Config) *http.Server {
	r := mux.NewRouter()
	staticHandler := http.StripPrefix("/static/", http.FileServer(http.Dir(c.Storage.Path)))
	if c.Server.NeedAuth {
		staticHandler = authMiddleware(staticHandler, c.Server.BearerToken)
	}
	r.PathPrefix("/static/").Handler(staticHandler)
	return &http.Server{
		Addr:              ":" + c.Server.Port,
		Handler:           r,
		ReadHeaderTimeout: 3 * time.Second,
	}
}

func authMiddleware(next http.Handler, token string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bearer := r.Header.Get("Authorization")
		if bearer == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		if fmt.Sprintf("Bearer %s", token) != bearer {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
