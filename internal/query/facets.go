package query

import "example.com/cookproposal/internal/domain"

type Facet struct {
	Statuses    []domain.Status
	Permissions []string
	Total       int
}

func BuildFacet(records []domain.Record) Facet {
	f := Facet{Statuses: DistinctStatuses(records), Permissions: []string{}, Total: len(records)}
	seen := map[string]bool{}
	for _, r := range records {
		if !seen[r.Permission] {
			seen[r.Permission] = true
			f.Permissions = append(f.Permissions, r.Permission)
		}
	}
	return f
}
func FacetContainsStatus(f Facet, s domain.Status) bool {
	for _, v := range f.Statuses {
		if v == s {
			return true
		}
	}
	return false
}
func FacetContainsPermission(f Facet, p string) bool {
	for _, v := range f.Permissions {
		if v == p {
			return true
		}
	}
	return false
}
func FilterWithFacet(records []domain.Record, f Facet) []domain.Record {
	out := []domain.Record{}
	for _, r := range records {
		if FacetContainsStatus(f, r.Status) && FacetContainsPermission(f, r.Permission) {
			out = append(out, r)
		}
	}
	return out
}
