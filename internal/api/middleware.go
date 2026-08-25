package api

import (
	"net/http"
	"strings"
)

func WithJSON(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
func WithMethod(method string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			ErrorResponse(w, http.StatusMethodNotAllowed, methodError(method))
			return
		}
		next.ServeHTTP(w, r)
	})
}
func methodError(method string) error { return &httpError{message: "method " + method + " required"} }

type httpError struct{ message string }

func (e *httpError) Error() string { return e.message }
func IsJSON(r *http.Request) bool {
	return strings.Contains(r.Header.Get("Content-Type"), "application/json")
}
func RequestID(r *http.Request) string {
	value := strings.TrimSpace(r.Header.Get("X-Request-ID"))
	if value == "" {
		return "deterministic-request"
	}
	return value
}
func RespondHealth(w http.ResponseWriter, healthy bool) {
	if healthy {
		writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	} else {
		writeJSON(w, http.StatusServiceUnavailable, map[string]string{"status": "unavailable"})
	}
}
