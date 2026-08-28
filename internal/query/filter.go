package query

import "example.com/cookproposal/internal/domain"

func ByStatus(records []domain.Record, status domain.Status) []domain.Record {
	return Search(records, Filter{Status: status})
}
func ByPermission(records []domain.Record, permission string) []domain.Record {
	return Search(records, Filter{Permission: permission})
}
func Titles(records []domain.Record) []string {
	out := make([]string, 0, len(records))
	for _, r := range records {
		out = append(out, r.Title)
	}
	return out
}
