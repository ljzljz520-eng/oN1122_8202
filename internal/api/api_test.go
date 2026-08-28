package api

import (
	"example.com/cookproposal/internal/flow027"
	"example.com/cookproposal/internal/store"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHTTPCreateAndGet(t *testing.T) {
	s, e := store.Open(t.TempDir() + "/api.db")
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	srv := NewServer(flow027.NewHandler(s))
	req := httptest.NewRequest("POST", "/records", strings.NewReader(`{"id":"r1","title":"One","summary":"S","permission":"team"}`))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	srv.HandlerFunc().ServeHTTP(rec, req)
	if rec.Code != 201 {
		t.Fatalf("%d", rec.Code)
	}
	get := httptest.NewRecorder()
	srv.HandlerFunc().ServeHTTP(get, httptest.NewRequest("GET", "/records/r1", nil))
	if get.Code != 200 {
		t.Fatalf("%d", get.Code)
	}
}
