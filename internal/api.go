package internal

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	uuid "github.com/satori/go.uuid"
)

type Server struct {
	StorageService *Storage
}

type Payload struct {
	State         string  `json:"state"`
	Amount        float32 `json:"amount"`
	TransactionID string  `json:"transaction_id"`
}

func NewServer(s *Storage) *Server {
	return &Server{StorageService: s}
}

func (s *Server) BalanceHandler(rw http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		rw.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintf(rw, "error: wrong request method")
		return
	}
	source := r.Header.Get("Source-Type")
	if !validateHeaderVal(source) {
		rw.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(rw, "error: missing or Source-Type header value")
		return
	}
	// Don't use in production!
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, "error: %s", err)
		return
	}
	defer r.Body.Close()

	// Unmarshal the payload
	p := Payload{}
	if err := json.Unmarshal(buf, &p); err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, "error: %s", err)
		return
	}

	t := Transaction{
		ID:            uuid.NewV4(),
		Time:          time.Now(),
		State:         p.State,
		OpAmount:      p.Amount,
		TransactionID: p.TransactionID,
	}
	// Main weight lifting
	id, err := s.StorageService.UpdateBalance(r.Context(), t)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(rw, "error: %s", err)
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Write(id.Bytes())
}

func validateHeaderVal(val string) bool {
	switch val {
	case "game", "payment", "server":
		return true
	}
	return false
}

func (s *Server) Run() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/update", s.BalanceHandler)
	return http.ListenAndServe(":8080", mux)
}
