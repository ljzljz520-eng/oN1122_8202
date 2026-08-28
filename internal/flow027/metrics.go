package flow027

import (
	"example.com/cookproposal/internal/domain"
)

type PermissionReport struct {
	Total   int
	Private int
	Team    int
	Public  int
	Empty   int
}

func PermissionMetrics(records []domain.Record) PermissionReport {
	m := PermissionReport{Total: len(records)}
	for _, r := range records {
		switch r.Permission {
		case "private":
			m.Private++
		case "team":
			m.Team++
		case "public":
			m.Public++
		default:
			m.Empty++
		}
	}
	return m
}
func IsBalanced(m PermissionReport) bool {
	if m.Total == 0 {
		return true
	}
	return m.Private > 0 && m.Team > 0 && m.Public > 0
}
func UniquePermissions(records []domain.Record) []string {
	seen := map[string]bool{}
	out := []string{}
	for _, r := range records {
		if !seen[r.Permission] {
			seen[r.Permission] = true
			out = append(out, r.Permission)
		}
	}
	return out
}
func PermissionForEach(records []domain.Record, fn func(domain.Record) bool) int {
	count := 0
	for _, r := range records {
		if fn(r) {
			count++
		}
	}
	return count
}
func FindPermission(records []domain.Record, permission string) []domain.Record {
	out := []domain.Record{}
	for _, r := range records {
		if r.Permission == permission {
			out = append(out, r)
		}
	}
	return out
}
