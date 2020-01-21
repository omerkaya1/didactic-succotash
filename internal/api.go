package internal

import (
	"fmt"
	"net/http"
)

type Server struct {
	StorageService *Storage
}

func NewServer(s *Storage) *Server {
	return &Server{StorageService: s}
}

// BalanceHandler is the only REST API handler function we provide
func (s *Server) BalanceHandler(rw http.ResponseWriter, r *http.Request) {
	buf := make([]byte, r.ContentLength, r.ContentLength*2)

	if n, err := r.Body.Read(buf); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, "error: %s, read: %d", err, n)
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write(buf)
}

// Main server routine
func (s *Server) Run() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/update", s.BalanceHandler)
	return http.ListenAndServe(":8080", mux)
}
