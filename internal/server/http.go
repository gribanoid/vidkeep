package server

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func NewHTTPServer(addr, storage string) *http.Server {
	r := mux.NewRouter()
	r.PathPrefix("/static/").
		Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(storage))))

	return &http.Server{Addr: addr, Handler: r, ReadHeaderTimeout: 3 * time.Second}
}
