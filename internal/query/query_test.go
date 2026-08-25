package query

import (
	"example.com/cookproposal/internal/domain"
	"testing"
)

func TestSearchFilters(t *testing.T) {
	rs := []domain.Record{domain.NewRecord("2", "Second", "", "public"), domain.NewRecord("1", "First", "", "team")}
	got := Search(rs, Filter{Text: "first"})
	if len(got) != 1 || got[0].ID != "1" {
		t.Fatalf("%v", got)
	}
}
