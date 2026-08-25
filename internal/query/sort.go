package query

import (
	"example.com/cookproposal/internal/domain"
	"sort"
)

type SortField string

const (
	SortByID      SortField = "id"
	SortByTitle   SortField = "title"
	SortByVersion SortField = "version"
)

func Sort(records []domain.Record, field SortField, descending bool) []domain.Record {
	out := make([]domain.Record, len(records))
	copy(out, records)
	sort.SliceStable(out, func(i, j int) bool {
		less := false
		switch field {
		case SortByTitle:
			less = out[i].Title < out[j].Title
		case SortByVersion:
			less = out[i].Version < out[j].Version
		default:
			less = out[i].ID < out[j].ID
		}
		if descending {
			return !less
		}
		return less
	})
	return out
}
func DistinctStatuses(records []domain.Record) []domain.Status {
	seen := map[domain.Status]bool{}
	out := []domain.Status{}
	for _, r := range records {
		if !seen[r.Status] {
			seen[r.Status] = true
			out = append(out, r.Status)
		}
	}
	return out
}
