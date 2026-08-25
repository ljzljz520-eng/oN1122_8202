package api

import "net/http"

func (s *Server) routes() {
	s.Mux.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) { writeJSON(w, 200, map[string]string{"status": "ok"}) })
	s.Mux.HandleFunc("/records", s.handleRecords)
	s.Mux.HandleFunc("/records/", s.handleRecord)
}
func (s *Server) handleRecords(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		b := parseBody(r)
		rec, e := s.Handler.Register(b["id"], b["title"], b["summary"], b["permission"])
		if e != nil {
			writeJSON(w, 400, map[string]string{"error": e.Error()})
			return
		}
		writeJSON(w, 201, rec)
		return
	}
	if r.Method == http.MethodGet {
		rs, e := s.records()
		if e != nil {
			writeJSON(w, 500, map[string]string{"error": e.Error()})
			return
		}
		writeJSON(w, 200, rs)
		return
	}
	writeJSON(w, 405, nil)
}
func (s *Server) handleRecord(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("/records/"):]
	if id == "" {
		writeJSON(w, 404, nil)
		return
	}
	if r.Method == http.MethodGet {
		rec, e := s.Handler.Get(id)
		if e != nil {
			writeJSON(w, 404, map[string]string{"error": e.Error()})
			return
		}
		writeJSON(w, 200, rec)
		return
	}
	if r.Method == http.MethodPatch {
		b := parseBody(r)
		rec, e := s.Handler.UpdatePermission(id, b["permission"])
		if e != nil {
			writeJSON(w, 400, map[string]string{"error": e.Error()})
			return
		}
		writeJSON(w, 200, rec)
		return
	}
	writeJSON(w, 405, nil)
}
