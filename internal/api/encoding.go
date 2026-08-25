package api

import (
	"encoding/json"
	"example.com/cookproposal/internal/domain"
	"net/http"
)

type RecordResponse struct {
	ID         string `json:"id"`
	Title      string `json:"title"`
	Summary    string `json:"summary"`
	Permission string `json:"permission"`
	Status     string `json:"status"`
	Version    int    `json:"version"`
}

func ToResponse(r domain.Record) RecordResponse {
	return RecordResponse{ID: r.ID, Title: r.Title, Summary: r.Summary, Permission: r.Permission, Status: string(r.Status), Version: r.Version}
}
func EncodeResponse(w http.ResponseWriter, r domain.Record) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(ToResponse(r))
}
func DecodeRecordRequest(req *http.Request) (map[string]string, error) {
	var body map[string]string
	if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
		return nil, err
	}
	return body, nil
}
func ErrorResponse(w http.ResponseWriter, status int, err error) {
	writeJSON(w, status, map[string]string{"error": err.Error()})
}
