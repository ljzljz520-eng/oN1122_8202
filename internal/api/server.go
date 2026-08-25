package api

import (
	"encoding/json"
	"example.com/cookproposal/internal/flow027"
	"example.com/cookproposal/internal/query"
	"net/http"
)

type Server struct {
	Handler *flow027.Handler
	Mux     *http.ServeMux
}

func NewServer(h *flow027.Handler) *Server {
	s := &Server{Handler: h, Mux: http.NewServeMux()}
	s.routes()
	return s
}
func (s *Server) HandlerFunc() http.Handler { return s.Mux }
func writeJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
func parseBody(r *http.Request) map[string]string {
	var body map[string]string
	_ = json.NewDecoder(r.Body).Decode(&body)
	return body
}
func (s *Server) records() ([]domainRecord, error) {
	rs, e := s.Handler.Store.ListRecords()
	out := make([]domainRecord, len(rs))
	for i, r := range rs {
		out[i] = domainRecord{ID: r.ID, Title: r.Title, Permission: r.Permission, Status: string(r.Status)}
	}
	return out, e
}

type domainRecord struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	Permission string `json:"permission"`
	Status     string `json:"status"`
}

var _ = query.Filter{}
