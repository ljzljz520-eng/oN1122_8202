package query

import (
	"example.com/cookproposal/internal/domain"
	"sort"
	"strings"
)

type Filter struct {
	Text       string
	Status     domain.Status
	Permission string
	Limit      int
}

func Search(records []domain.Record, f Filter) []domain.Record {
	out := make([]domain.Record, 0)
	for _, r := range records {
		if !matches(r, f) {
			continue
		}
		out = append(out, r)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].ID < out[j].ID })
	if f.Limit > 0 && len(out) > f.Limit {
		return out[:f.Limit]
	}
	return out
}
func matches(r domain.Record, f Filter) bool {
	if f.Status != "" && r.Status != f.Status {
		return false
	}
	if f.Permission != "" && r.Permission != f.Permission {
		return false
	}
	if f.Text != "" && !strings.Contains(strings.ToLower(r.Title), strings.ToLower(f.Text)) {
		return false
	}
	return true
}
