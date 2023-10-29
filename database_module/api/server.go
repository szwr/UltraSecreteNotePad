package api

import (
	"fmt"
	"net/http"

    "example.com/db"
)

type Server struct {
	listenAddr string
    databaseClient *db.RedisClient
}

func NewServer(listenAddr string, databaseClient *db.RedisClient) *Server {
	return &Server{
		listenAddr: listenAddr,
        databaseClient: databaseClient,
	}
}

func (s *Server) Start() error {
	http.HandleFunc("/read-db", s.handleGetQuery)
	http.HandleFunc("/add-value", s.handleAddValue)
    return http.ListenAndServe(s.listenAddr, nil)
}

func (s *Server) handleGetQuery(w http.ResponseWriter, r *http.Request) {
    key := r.URL.Query().Get("link")

    val, err := s.databaseClient.GetKeyValue(key)
    if err != nil {
        fmt.Fprintf(w, "Couldn't find key value")
    } else {
        fmt.Fprintf(w, "%v", val)
    }
}

func (s *Server) handleAddValue(w http.ResponseWriter, r *http.Request) {
    key := r.URL.Query().Get("link")
    val := r.URL.Query().Get("message")

    err := s.databaseClient.AddKeyValue(key, val)
    if err != nil {
        fmt.Fprintf(w, "Couldn't add new key and value")
    } else {
        fmt.Fprintf(w, "key and value added", val)
    }
}
